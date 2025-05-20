variable "file_payload" {
  description = "The payload string for the file-based push"
  type        = string
}

variable "file_note" {
  description = "Optional note for the file push"
  type        = string
  default     = "File-based push example"
}

variable "file_paths" {
  description = "List of file paths to attach"
  type        = list(string)
}

variable "file_expire_after_days" {
  description = "Days until expiration"
  type        = number
  default     = 3
}

variable "file_expire_after_views" {
  description = "Views until expiration"
  type        = number
  default     = 2
}

variable "file_passphrase" {
  description = "Optional passphrase required to retrieve the push"
  type        = string
  default     = ""
}
