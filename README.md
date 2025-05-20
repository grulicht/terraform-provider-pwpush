<p align="center">
  <a href="https://registry.terraform.io/providers/grulicht/pwpush/latest/docs">
    <img src="https://www.datocms-assets.com/2885/1731373310-terraform_white.svg" alt="Terraform Logo" width="200">
  </a>
  &nbsp;&nbsp;&nbsp;
  <a href="https://github.com/grulicht/terraform-provider-pwpush">
    <img src="https://docs.pwpush.com/assets/images/logo-transparent-sm-bare.png" alt="pwpush-provider-terraform" width="200">
  </a>
  &nbsp;&nbsp;&nbsp;
  <a href="https://search.opentofu.org/provider/grulicht/pwpush/latest">
    <img src="https://raw.githubusercontent.com/opentofu/brand-artifacts/main/full/transparent/SVG/on-dark.svg#gh-dark-mode-only" alt="pwpush-provider-opentofu" width="200">
  </a>
  <h3 align="center" style="font-weight: bold">Terraform Provider for Password Pusher</h3>
  <p align="center">
    <a href="https://github.com/grulicht/terraform-provider-pwpush/graphs/contributors">
      <img alt="Contributors" src="https://img.shields.io/github/contributors/grulicht/terraform-provider-pwpush">
    </a>
    <a href="https://golang.org/doc/devel/release.html">
      <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/grulicht/terraform-provider-pwpush">
    </a>
    <a href="https://github.com/grulicht/terraform-provider-pwpush/actions?query=workflow%3Arelease">
      <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/grulicht/terraform-provider-pwpush/release.yml?label=release">
    </a>
    <a href="https://github.com/grulicht/terraform-provider-pwpush/releases">
      <img alt="GitHub release (latest by date including pre-releases)" src="https://img.shields.io/github/v/release/grulicht/terraform-provider-pwpush?include_prereleases">
    </a>
  </p>
  <p align="center">
    <a href="https://github.com/grulicht/terraform-provider-pwpush/tree/main/docs"><strong>Explore the docs »</strong></a>
  </p>
</p>

# Pwpush Terraform Provider

A [Terraform](https://www.terraform.io) provider to manage [Password Pusher](https://pwpush.com/) resources via its REST API using Terraform.

It supports provisioning and configuration of secure one-time pushes – including secrets, files, URLs, and QR codes – with expiration and passphrase protection.

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

## Requirements

- Terraform v0.13+
- Password Pusher server or hosted instance
- Go 1.21+ (if building from source)

---

## Building and Installing

```bash
make build
```

---

## Provider Support

| Provider                                                                 | Status                                                              |
|--------------------------------------------------------------------------|----------------------------------------------------------------------|
| [Terraform](https://registry.terraform.io/providers/grulicht/pwpush/latest)  | ![Done](https://img.shields.io/badge/status-done-brightgreen)       |
| [OpenTofu](https://search.opentofu.org/provider/grulicht/pwpush/latest)      | ![Done](https://img.shields.io/badge/status-done-brightgreen)       |

---

## Example Provider Configuration

```hcl
provider "pwpush" {
  endpoint = "https://pwpush.example.com"       # optional
  email    = "your@example.com"                 # optional
  token    = "your-api-token"                   # optional
  skip_ssl_verify = true                        # optional
}
```

---

## Authentication

> 🔐 **Authentication** is required for file uploads and notes. The provider uses:
- `X-User-Email`
- `X-User-Token`

Example:

```hcl
provider "pwpush" {
  email = "user@example.com"
  token = "super_secret_token"
}
```

### Environment Variables

```bash
export PWPUSH_ENDPOINT="https://pwpush.example.com"
export PWPUSH_EMAIL="your@example.com"
export PWPUSH_TOKEN="your-api-token"
```

---

| Name             | Type    | Required | Description                                                                        |
|------------------|---------|----------|------------------------------------------------------------------------------------|
| `endpoint`       | string  | ❌ no    | The URL of the Password Pusher server - Default: `https://pwpush.com`              |
| `email`          | string  | ❌ no    | Email for authenticated pushes (if needed)                                         |
| `token`          | string  | ❌ no    | API token for authenticated pushes                                                 |
| `skip_ssl_verify`| bool    | ❌ no    | Skip TLS certificate verification (default: false)                                 |

---

## 🧩 Supported Resources

| Resource       | Docs                                  | Example                       | Status                                                        | E2E Tests             |
|----------------|---------------------------------------|-------------------------------|---------------------------------------------------------------|-----------------------|
| `pwpush_push`  | [push](docs/resources/push.md) | [example (push*)](examples/)     | ![Done](https://img.shields.io/badge/status-done-brightgreen) |  ✅ Daily verified    |

---

### 💡 Missing a resource?

Open a [feature request](https://github.com/grulicht/terraform-provider-pwpush/issues) or [contribute](https://github.com/grulicht/terraform-provider-pwpush/pulls) directly.

See [CONTRIBUTING.md](./.github/CONTRIBUTING.md) for details.

---

## 💬 Community & Feedback

We welcome all questions, issues and ideas — open a GitHub [Issue](https://github.com/grulicht/terraform-provider-pwpush/issues) or [PR](https://github.com/grulicht/terraform-provider-pwpush/pulls).

Have questions, suggestions or want to contribute ideas?  
[GitHub Discussions](https://github.com/grulicht/terraform-provider-pwpush/discussions)

---

## ✅ Daily End-to-End Testing

This provider is validated daily using GitHub Actions on real Password Pusher instances using real `terraform plan` and `terraform apply`.

> 💡 This helps catch regressions and API changes early.

---

### 🧪 Localy Testing
To test the provider locally, start the Pwpush Web UI using Docker Compose:
```sh
make up
```
Then open `http://localhost:5100` in your browser by:
```sh
make launch
```

---

### Testing a new version of the Pwpush provider
After making changes to the provider source code, follow these steps:
Build the provider binary:
```sh
make build
```
Install the binary into the local Terraform plugin directory:
```sh
make install-plugin
```
Update your main.tf to use the local provider source
Add the following to your Terraform configuration:
```sh
terraform {
  required_providers {
    pwpush = {
      source  = "localdomain/local/pwpush"
    }
  }
}
```
Now you're ready to test your provider against the local or cloud Pwpush instance.

---

## 📍 Roadmap

See [Issues](https://github.com/grulicht/terraform-provider-pwpush/issues) and [CONTRIBUTING.md](./.github/CONTRIBUTING.md) for development plans.

---

## 👤 Author

Maintained by [Tomáš Grulich](https://github.com/grulicht).

---

## 📜 License

MIT License. See [LICENSE](./LICENSE).

---

## 🙌 Acknowledgements

- [Password Pusher](https://pwpush.com)
- [HashiCorp Terraform](https://www.terraform.io)
- [OpenTofu](https://opentofu.org)
