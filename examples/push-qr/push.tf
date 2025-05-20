resource "pwpush_push" "example" {
  payload = var.qr_payload
  kind    = "qr"
  note    = var.qr_note
}