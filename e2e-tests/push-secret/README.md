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
| <a name="input_expire_after_days"></a> [expire\_after\_days](#input\_expire\_after\_days) | Number of days after which the push expires | `number` | `2` | no |
| <a name="input_expire_after_views"></a> [expire\_after\_views](#input\_expire\_after\_views) | Number of views after which the push expires | `number` | `3` | no |
| <a name="input_note"></a> [note](#input\_note) | Optional note associated with the push | `string` | `""` | no |
| <a name="input_payload"></a> [payload](#input\_payload) | The secret payload to be pushed | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_push_url"></a> [push\_url](#output\_push\_url) | n/a |
<!-- END_TF_DOCS -->