resource "pwpush_push" "example" {
  payload            = var.payload
  note               = var.note
  expire_after_days  = var.expire_after_days
  expire_after_views = var.expire_after_views
}
