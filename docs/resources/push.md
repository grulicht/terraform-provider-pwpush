# 🚀 Resource Documentation: `pwpush_push`

## Overview

The `pwpush_push` resource allows you to create and expire secure, temporary **password pushes** using [Password Pusher](https://pwpush.com) – a self-hosted or cloud service for one-time secret sharing. You can push secrets with optional expiration, note, and files.

You can use the resource to:

- Create temporary secrets (`pwpush_push`)
- Push files, links or QR codes
- Automatically expire secrets based on views or days
- Protect with passphrases
- Delete/expire pushes using `terraform destroy`

---

## Example Usage

### Create a Simple Secret Push

```hcl
resource "pwpush_push" "example" {
  payload            = "my_secret_password"
  note               = "Login to internal VPN"
  expire_after_days  = 2
  expire_after_views = 3
}

output "push_url" {
  value = pwpush_push.example.html_url
}
```

### Create a URL Push

```hcl
resource "pwpush_push" "link_push" {
  payload = "https://vpn.company.com"
  kind    = "url"
}

output "push_url" {
  value = pwpush_push.link_push.html_url
}

```

### Create QR Push
```hcl
resource "pwpush_push" "qr_push" {
  payload = "WiFi: MyNetwork;Password: s3cr3t"
  kind    = "qr"
}

output "push_url" {
  value = pwpush_push.qr_push.html_url
}

```

### Create a Push With Files (requires authentication)

```hcl
provider "pwpush" {
  email  = "you@example.com"
  token  = "your_api_token"
}

resource "pwpush_push" "with_files" {
  payload  = "confidential"
  note     = "Employment contract"

  files = [
    "./contracts/employment.pdf",
    "./contracts/nda.pdf"
  ]

  expire_after_views = 1
  passphrase         = "strong-passphrase"
}

output "push_url" {
  value = pwpush_push.with_files.html_url
}
```

---

## Lifecycle & Behavior

- Secrets are stored and retrieved using Password Pusher API.
- Once created, pushes can be **expired** (deleted) by Terraform via `terraform destroy`.
- **Updates are not supported** – changing any field will recreate the push (ForceNew).

---

## Arguments Reference

| Name                  | Type     | Required | Description                                                                 |
|-----------------------|----------|----------|-----------------------------------------------------------------------------|
| `payload`             | string   | ✅ yes   | The secret value or URL to share.                                           |
| `kind`                | string   | ❌ no    | Type of push: `text` (default), `url`, `qr`, `file`.                        |
| `note`                | string   | ❌ no    | Internal note (visible only to creator).                                    |
| `name`                | string   | ❌ no    | Display name shown in dashboards or emails.                                 |
| `passphrase`          | string   | ❌ no    | Optional passphrase required to open the push.                              |
| `expire_after_days`   | integer  | ❌ no    | Delete the push after this many days.                                       |
| `expire_after_views`  | integer  | ❌ no    | Delete the push after this many views.                                      |
| `deletable_by_viewer` | bool     | ❌ no    | Allow the recipient to delete the push after viewing.                       |
| `retrieval_step`      | bool     | ❌ no    | Avoid URL scanners consuming the push link (adds human step).               |
| `account_id`          | integer  | ❌ no    | Account ID (for multi-account setups).                                      |
| `files`               | list     | ❌ no    | List of file paths to attach to the push (authentication required).         |

---

## 🧮 Computed Outputs

| Name       | Description                                             |
|------------|---------------------------------------------------------|
| `id`       | The push `url_token`, used as a unique identifier       |
| `html_url` | The full URL to share with the recipient (copy & paste) |
