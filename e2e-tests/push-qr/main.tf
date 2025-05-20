terraform {
  required_providers {
    pwpush = {
      source  = "grulicht/pwpush"
    }
  }
}

provider "pwpush" {
  endpoint            = "http://localhost:5100"
  # email             = "your-e-mail@some-domain.com"
  # token             = "your-token"
  # skip_ssl_verify   = "true"
}