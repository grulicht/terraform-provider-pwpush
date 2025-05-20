<!-- BEGIN_TF_DOCS -->


## Providers

| Name | Version |
|------|---------|
| <a name="provider_pwpush"></a> [pwpush](#provider\_pwpush) | n/a |

## Resources

| Name | Type |
|------|------|
| [pwpush_push.example](https://registry.terraform.io/providers/grulicht/pwpush/latest/docs/resources/push) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_url_note"></a> [url\_note](#input\_url\_note) | Optional note for the URL push | `string` | `"URL-based push example"` | no |
| <a name="input_url_payload"></a> [url\_payload](#input\_url\_payload) | The URL to be pushed | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_push_url"></a> [push\_url](#output\_push\_url) | n/a |
<!-- END_TF_DOCS -->