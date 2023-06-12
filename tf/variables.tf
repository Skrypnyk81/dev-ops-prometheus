variable "GOOGLE_REGION" {
  type        = string
  default     = "europe-west12-a"
  description = "GCP region to use"
}

variable "GOOGLE_PROJECT" {
  type        = string
  description = "Bitcoin Bot"
}

variable "GKE_NUM_NODES" {
  type           = number
  default        = 1
  description = "GKE nodes"

}
