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
| <a name="input_file_expire_after_days"></a> [file\_expire\_after\_days](#input\_file\_expire\_after\_days) | Days until expiration | `number` | `3` | no |
| <a name="input_file_expire_after_views"></a> [file\_expire\_after\_views](#input\_file\_expire\_after\_views) | Views until expiration | `number` | `2` | no |
| <a name="input_file_note"></a> [file\_note](#input\_file\_note) | Optional note for the file push | `string` | `"File-based push example"` | no |
| <a name="input_file_passphrase"></a> [file\_passphrase](#input\_file\_passphrase) | Optional passphrase required to retrieve the push | `string` | `""` | no |
| <a name="input_file_paths"></a> [file\_paths](#input\_file\_paths) | List of file paths to attach | `list(string)` | n/a | yes |
| <a name="input_file_payload"></a> [file\_payload](#input\_file\_payload) | The payload string for the file-based push | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_push_url"></a> [push\_url](#output\_push\_url) | n/a |
<!-- END_TF_DOCS -->