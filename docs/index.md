# 🔐 Pwpush Terraform/OpenTofu Provider

A Terraform/OpenTofu provider to manage secure one-time secret pushes using [Password Pusher](https://github.com/pglombardo/PasswordPusher) via its REST API.

It supports provisioning of pushes with text, URLs, QR codes, or attached files, including optional expirations, passphrases, and notes.

---

## 💼 Use Cases

Here are practical scenarios where using the `pwpush` Terraform provider is ideal compared to manual usage:

- ✅ **GitOps & automation**: Push secrets in CI/CD pipelines as part of automated infrastructure provisioning.
- 🔐 **Temporary credential sharing**: Safely share credentials with newly created users, contractors, or services.
- 🛠️ **Immutable environments**: Recreate secrets on every Terraform run as part of a secure, reproducible workflow.
- 📤 **Credential delivery**: Deliver generated passwords (e.g., via Terraform random provider) to human recipients.
- 🚫 **Avoid storing secrets in state files**: Send secrets via expiring link instead of embedding in outputs or storage.
- 🧪 **Secret testing environments**: Create disposable, time-limited secrets for test or staging use.
- 🧾 **Auditable provisioning**: Keep track of when secrets are pushed and for what purpose via version control.
- 📦 **Multi-platform provisioning**: Coordinate secret delivery alongside cloud or on-prem resource provisioning.

This provider makes it easy to integrate secure secret sharing into infrastructure-as-code workflows, especially where automation and reproducibility are key.

---

## 🏷️ Provider Support
| Provider       | Provider Support Status              |
|----------------|--------------------------------------|
| [Terraform](https://registry.terraform.io/providers/grulicht/pwpush/latest) | ![Done](https://img.shields.io/badge/status-done-brightgreen) |
| [OpenTofu](https://search.opentofu.org/provider/grulicht/pwpush/latest)     | ![Done](https://img.shields.io/badge/status-done-brightgreen) |

---

## ⚙️ Example Provider Configuration

```hcl
provider "pwpush" {
  endpoint = "..." # optional, The URL of the Password Pusher server - Default: `https://pwpush.com`
  email    = "..."  # optional, required for file uploads
  token    = "..."  # optional, required for file uploads
}
```

---

## 🔐 Authentication

This provider supports authentication via **email and token** using custom HTTP headers:

- `X-User-Email`
- `X-User-Token`

### Authentication is only required for:

- File uploads (`files[]`)
- Notes visible only to the push creator
- Account-specific features

---

### Environment variables

You can provide configuration via environment variables:

```bash
export PWPUSH_ENDPOINT="https://pwpush.com"
export PWPUSH_EMAIL="you@example.com"
export PWPUSH_TOKEN="your_api_token"
```

---

### Arguments Reference

| Name             | Type    | Required | Description                                                                        |
|------------------|---------|----------|------------------------------------------------------------------------------------|
| `endpoint`       | string  | ❌ no    | The URL of the Password Pusher server - Default: `https://pwpush.com`              |
| `email`          | string  | ❌ no    | Email for authenticated pushes (if needed)                                         |
| `token`          | string  | ❌ no    | API token for authenticated pushes                                                 |
| `skip_ssl_verify`| bool    | ❌ no    | Skip TLS certificate verification (default: false)                                 |

---

## 🧩 Supported Resources

| Resource       | Status                              |
|----------------|-------------------------------------|
| `pwpush_push`  | ![Done](https://img.shields.io/badge/status-done-brightgreen) |

### 💡 Missing a resource?
Is there a Pwpush resource you'd like to see supported?

👉 [Open an issue](https://github.com/grulicht/terraform-provider-pwpush/issues) and we’ll consider it for implementation — or even better, submit a [Pull Request](https://github.com/grulicht/terraform-provider-pwpush/pulls) to contribute directly!

📘 See [CONTRIBUTING.md](https://github.com/grulicht/terraform-provider-pwpush/blob/main/.github/CONTRIBUTING.md) for guidelines.

---

## 💬 Community & Feedback
Have questions, suggestions or want to contribute ideas?  
[GitHub Discussions](https://github.com/grulicht/terraform-provider-pwpush/discussions)

Want to report issues, submit pull requests or browse the source code?  
Check out the [GitHub Repository](https://github.com/grulicht/terraform-provider-pwpush) for this provider.

## 🧪 Development & Contribution

We welcome contributions!  
Open an issue or submit a PR on [GitHub](https://github.com/grulicht/terraform-provider-pwpush).

---

## ✅ Daily End-to-End Testing
To ensure maximum reliability and functionality of this provider, **automated end-to-end tests are executed every day** via GitHub Actions.

These tests run against a real Pwpush instance (started using docker compose) and validate the majority of supported resources using real Terraform plans and applies.

> 💡 This helps catch regressions early and ensures the provider remains fully operational and compatible with the Pwpush API.

---

## 👤 Author

Maintained by [Tomáš Grulich](https://github.com/grulicht).

---

## 📜 License

This provider is licensed under the [MIT License](https://github.com/grulicht/terraform-provider-pwpush/blob/main/LICENSE).

---

## 🌐 Acknowledgements

- [Password Pusher](https://pwpush.com)
- HashiCorp Terraform
- [OpenTofu](https://opentofu.org/)
