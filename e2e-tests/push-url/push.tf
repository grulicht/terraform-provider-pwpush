resource "pwpush_push" "example" {
  payload = var.url_payload
  kind    = "url"
  note    = var.url_note
}
