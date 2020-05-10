resource "aws_cloudfront_origin_access_identity" "farto_cloud" {
  comment = "access-identity-farto.cloud${var.origin_path}"
}

resource "dnsimple_record" "farto_cloud" {
  domain = "farto.cloud"
  name   = var.subdomain
  value  = aws_cloudfront_distribution.farto_cloud.domain_name
  type   = "CNAME"
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
    acm_certificate_arn            = "arn:aws:acm:us-east-1:081549132651:certificate/cdb6e100-6cd2-41f8-9bec-56ac0fd03293"
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
    lambda_function_association {
      event_type = "viewer-request"
      lambda_arn = "arn:aws:lambda:us-east-1:081549132651:function:fartoCloudAuth:2"
    }
  }
}
