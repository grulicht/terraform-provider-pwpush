variable "qr_payload" {
  description = "The content to be encoded in the QR push"
  type        = string
}

variable "qr_note" {
  description = "Optional note for the QR push"
  type        = string
  default     = "QR-based push example"
}
