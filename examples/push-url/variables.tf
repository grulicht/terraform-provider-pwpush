variable "url_payload" {
  description = "The URL to be pushed"
  type        = string
}

variable "url_note" {
  description = "Optional note for the URL push"
  type        = string
  default     = "URL-based push example"
}
