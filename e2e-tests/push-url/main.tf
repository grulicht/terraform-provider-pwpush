terraform {
  required_providers {
    pwpush = {
      source = "grulicht/pwpush"
    }
  }
}

provider "pwpush" {
  # endpoint          = "some-pwpush-url"
  # email             = "your-e-mail@some-domain.com"
  # token             = "your-token"
  # skip_ssl_verify   = "true"
}