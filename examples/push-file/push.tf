resource "pwpush_push" "example" {
  payload            = var.file_payload
  note               = var.file_note
  files              = var.file_paths
  expire_after_days  = var.file_expire_after_days
  expire_after_views = var.file_expire_after_views
  passphrase         = var.file_passphrase
}