variable "aliases" { type = list }
variable "bucket_domain_name" {}
variable "origin_id" {}
variable "origin_path" {}
variable "subdomain" {}

variable "lambda_auth" {
  default     = null
  type        = map
  description = <<EOD
Map with "user" and "password" keys.
If unset, the cloudfront distribution will not be password protected.
EOD
}
variable "lambda_role_arn" { default = null }
