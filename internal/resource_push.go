package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePwpushPush() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePwpushPushCreate,
		DeleteContext: resourcePwpushPushDelete,
		ReadContext:   resourcePwpushPushReadNoop, // No API for reading the content
		Schema: map[string]*schema.Schema{
			"payload": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kind": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "text",
			},
			"note": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"passphrase": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"expire_after_days": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"expire_after_views": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"deletable_by_viewer": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"retrieval_step": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"account_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"files": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"html_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourcePwpushPushCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*APIClient)

	// Multipart buffer
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Required payload
	_ = writer.WriteField("password[payload]", d.Get("payload").(string))

	// Optional fields
	for key, tfKey := range map[string]string{
		"password[kind]":                "kind",
		"password[note]":                "note",
		"password[name]":                "name",
		"password[passphrase]":          "passphrase",
		"password[expire_after_days]":   "expire_after_days",
		"password[expire_after_views]":  "expire_after_views",
		"password[deletable_by_viewer]": "deletable_by_viewer",
		"password[retrieval_step]":      "retrieval_step",
		"password[account_id]":          "account_id",
	} {
		if v, ok := d.GetOk(tfKey); ok {
			_ = writer.WriteField(key, fmt.Sprintf("%v", v))
		}
	}

	// Files
	if v, ok := d.GetOk("files"); ok {
		files := v.([]interface{})
		for _, f := range files {
			path := f.(string)
			file, err := os.Open(path)
			if err != nil {
				return diag.FromErr(fmt.Errorf("failed to open file %s: %w", path, err))
			}
			defer file.Close()

			part, err := writer.CreateFormFile("password[files][]", filepath.Base(path))
			if err != nil {
				return diag.FromErr(fmt.Errorf("failed to create multipart for file %s: %w", path, err))
			}
			if _, err := io.Copy(part, file); err != nil {
				return diag.FromErr(fmt.Errorf("failed to copy file %s into request: %w", path, err))
			}
		}
	}

	writer.Close()

	// Prepare request
	req, err := http.NewRequest("POST", client.Endpoint+"/p.json", &requestBody)
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Auth headers if set
	if client.Email != "" {
		req.Header.Set("X-User-Email", client.Email)
	}
	if client.Token != "" {
		req.Header.Set("X-User-Token", client.Token)
	}

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return diag.Errorf("unexpected response from pwpush API: %d - %s", resp.StatusCode, body)
	}

	var result struct {
		UrlToken string `json:"url_token"`
		HtmlURL  string `json:"html_url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return diag.FromErr(err)
	}

	if result.HtmlURL == "" && result.UrlToken != "" {
		endpoint := strings.TrimRight(client.Endpoint, "/")
		result.HtmlURL = fmt.Sprintf("%s/p/%s", endpoint, result.UrlToken)
	}

	// Set ID and html_url
	d.SetId(result.UrlToken)
	_ = d.Set("html_url", result.HtmlURL)

	return nil
}

func resourcePwpushPushDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*APIClient)
	id := d.Id()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/p/%s.json", client.Endpoint, id), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	if client.Email != "" {
		req.Header.Set("X-User-Email", client.Email)
	}
	if client.Token != "" {
		req.Header.Set("X-User-Token", client.Token)
	}

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 || resp.StatusCode == 422 {
		// Read body just in case
		body, _ := io.ReadAll(resp.Body)
		if strings.Contains(string(body), "already expired") {
			d.SetId("")
			return nil
		}
		return diag.Errorf("failed to delete push: %d - %s", resp.StatusCode, body)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		return diag.Errorf("failed to delete push: %d - %s", resp.StatusCode, body)
	}

	d.SetId("")
	return nil
}

// Read is a no-op, since pwpush does not offer GET /p/<id>.json without consuming the view
func resourcePwpushPushReadNoop(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
