variable "name" {}

variable "domain" {
  description = "DNS domain for the DNSimple record"
}

variable "domain_root" {
  default     = false
  description = <<EOD
Whether this should resolve at the domain root.
true: an ALIAS record will be used.
false: a CNAME will be used with the name variable as its subdomain.
EOD
}

variable "aliases" {
  type        = list(string)
  description = <<EOD
List of alternate domains.
Note you can only use one TLS certificate per distribution.
EOD
}

variable "acm_certificate_arn" {}
variable "bucket_domain_name" {}

variable "origin_id" {}
variable "origin_path" {}

variable "lambda_auth" {
  default     = null
  type        = map(any)
  description = <<EOD
Map with "user" and "password" keys.
If unset, the cloudfront distribution will not be password protected.
EOD
}
variable "lambda_role_arn" { default = null }
