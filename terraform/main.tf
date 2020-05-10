terraform {
  backend "s3" {
    bucket = "ttttfstate"
    key    = "farto"
    region = "us-west-2"
  }
}

provider "aws" {
  region = "us-west-2"
}

locals {
  origin_id   = "S3-farto.cloud/test"
  origin_path = "/test"
  aliases     = ["test.farto.cloud"]
}

resource "aws_s3_bucket" "farto_cloud" {
  bucket = "farto.cloud"
  acl    = "private"
}

resource "aws_cloudfront_origin_access_identity" "farto_cloud" {
  comment = "access-identity-farto.cloud/test"
}

resource "aws_cloudfront_distribution" "farto_cloud" {

  origin {
    domain_name = aws_s3_bucket.farto_cloud.bucket_domain_name

    origin_id   = local.origin_id
    origin_path = local.origin_path

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.farto_cloud.cloudfront_access_identity_path
    }
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  aliases             = local.aliases
  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "site/index.html"
  price_class         = "PriceClass_All"

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = local.origin_id
    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }
    viewer_protocol_policy = "redirect-to-https"
  }

  viewer_certificate {
    acm_certificate_arn            = "arn:aws:acm:us-east-1:081549132651:certificate/cdb6e100-6cd2-41f8-9bec-56ac0fd03293"
    cloudfront_default_certificate = false
    minimum_protocol_version       = "TLSv1.2_2018"
    ssl_support_method             = "sni-only"
  }
}
