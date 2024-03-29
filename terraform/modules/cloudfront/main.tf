terraform {
  required_providers {
    dnsimple = {
      source = "dnsimple/dnsimple"
    }
  }
}

locals {

  lambda_auth_enabled = var.lambda_auth != null

  dns_name = var.domain_root ? "" : var.name
  dns_type = var.domain_root ? "ALIAS" : "CNAME"
}

resource "aws_cloudfront_origin_access_identity" "farto_cloud" {
  comment = "access-identity-farto.cloud${var.origin_path}"
}

resource "dnsimple_record" "farto_cloud" {
  domain = var.domain
  name   = local.dns_name
  value  = aws_cloudfront_distribution.farto_cloud.domain_name
  type   = local.dns_type
}

resource "aws_cloudfront_distribution" "farto_cloud" {

  aliases             = var.aliases
  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "site/index.html"
  price_class         = "PriceClass_100"

  origin {
    domain_name = var.bucket_domain_name

    origin_id   = var.origin_id
    origin_path = var.origin_path

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.farto_cloud.cloudfront_access_identity_path
    }
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn            = var.acm_certificate_arn
    cloudfront_default_certificate = false
    minimum_protocol_version       = "TLSv1.2_2018"
    ssl_support_method             = "sni-only"
  }

  default_cache_behavior {
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = var.origin_id
    viewer_protocol_policy = "redirect-to-https"
    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }
  }

  ordered_cache_behavior {
    path_pattern           = "/site/*"
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = var.origin_id
    default_ttl            = 0
    viewer_protocol_policy = "redirect-to-https"
    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }

    dynamic "lambda_function_association" {
      for_each = toset(local.lambda_auth_enabled ? [0] : [])
      content {
        event_type = "viewer-request"
        lambda_arn = aws_lambda_function.farto_auth[0].qualified_arn
      }
    }
  }
}

# Lambda@Edge functions must be in us-east-1 for some reason.
provider "aws" {
  region = "us-east-1"
  alias  = "us_east_1"
}

resource "aws_lambda_function" "farto_auth" {

  count = local.lambda_auth_enabled ? 1 : 0

  provider      = aws.us_east_1
  function_name = "farto-auth-${var.name}"
  filename      = "lambda-auth-${var.name}.zip"
  role          = var.lambda_role_arn
  handler       = "lambda-auth-${var.name}.handler"
  runtime       = "nodejs18.x"
  publish       = true

  depends_on = [null_resource.lambda_zip]
}

resource "null_resource" "lambda_zip" {

  count = local.lambda_auth_enabled ? 1 : 0

  triggers = {
    template = data.template_file.lambda_auth[0].rendered
  }
  provisioner "local-exec" {
    command = <<EOC
echo "${data.template_file.lambda_auth[0].rendered}" > lambda-auth-${var.name}.js \
  && zip lambda-auth-${var.name}.zip lambda-auth-${var.name}.js;
EOC
  }
}

data "template_file" "lambda_auth" {

  count = local.lambda_auth_enabled ? 1 : 0

  template = file("${path.module}/lambda-auth.tpl.js")
  vars = {
    user     = var.lambda_auth.user
    password = var.lambda_auth.password
  }
}
