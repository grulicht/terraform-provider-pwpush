variable "payload" {
  description = "The secret payload to be pushed"
  type        = string
  default     = "some-secret"
}

variable "note" {
  description = "Optional note associated with the push"
  type        = string
  default     = ""
}

variable "expire_after_days" {
  description = "Number of days after which the push expires"
  type        = number
  default     = 2
}

variable "expire_after_views" {
  description = "Number of views after which the push expires"
  type        = number
  default     = 3
}
