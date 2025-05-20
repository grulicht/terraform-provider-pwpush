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
| <a name="input_qr_note"></a> [qr\_note](#input\_qr\_note) | Optional note for the QR push | `string` | `"QR-based push example"` | no |
| <a name="input_qr_payload"></a> [qr\_payload](#input\_qr\_payload) | The content to be encoded in the QR push | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_push_url"></a> [push\_url](#output\_push\_url) | n/a |
<!-- END_TF_DOCS -->