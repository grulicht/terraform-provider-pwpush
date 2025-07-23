package internal

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider defines the Terraform provider for Password Pusher (pwpush).
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PWPUSH_ENDPOINT", "https://pwpush.com"),
				Description: "Base URL of the Password Pusher service.",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PWPUSH_EMAIL", nil),
				Description: "User email for authenticated pushes.",
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PWPUSH_TOKEN", nil),
				Description: "User token for authenticated pushes.",
			},
			"skip_ssl_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Skip TLS certificate verification.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"pwpush_push": resourcePwpushPush(),
		},
		ConfigureContextFunc: configureProvider,
	}
}

// APIClient holds client config for pwpush.
type APIClient struct {
	Endpoint   string
	Email      string
	Token      string
	HTTPClient *http.Client
}

// DoJSONRequest sends a JSON-encoded request.
func (c *APIClient) DoJSONRequest(method, path string, headers map[string]string, body interface{}) (*http.Response, error) {
	var buf io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, c.Endpoint+path, buf)
	if err != nil {
		return nil, err
	}

	if _, ok := headers["Content-Type"]; !ok {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.Email != "" {
		req.Header.Set("X-User-Email", c.Email)
	}
	if c.Token != "" {
		req.Header.Set("X-User-Token", c.Token)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return c.HTTPClient.Do(req)
}

// DoMultipartRequest sends a multipart/form-data request.
func (c *APIClient) DoMultipartRequest(req *http.Request) (*http.Response, error) {
	if c.Email != "" {
		req.Header.Set("X-User-Email", c.Email)
	}
	if c.Token != "" {
		req.Header.Set("X-User-Token", c.Token)
	}
	return c.HTTPClient.Do(req)
}

// configureProvider initializes the API client.
func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	endpoint := d.Get("endpoint").(string)
	email := d.Get("email").(string)
	token := d.Get("token").(string)
	skipSSL := d.Get("skip_ssl_verify").(bool)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipSSL,
		},
	}
	httpClient := &http.Client{Transport: transport}

	client := &APIClient{
		Endpoint:   endpoint,
		Email:      email,
		Token:      token,
		HTTPClient: httpClient,
	}

	var diags diag.Diagnostics
	return client, diags
}
