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

data "aws_iam_policy_document" "farto_cloud" {

  policy_id = "PolicyForCloudFrontPrivateContent"

  dynamic "statement" {

    for_each = {
      mom  = module.cloudfront_mom.origin_access_identity_iam_arn,
      test = module.cloudfront_test.origin_access_identity_iam_arn,
    }

    content {
      actions = [
        "s3:GetObject",
      ]
      resources = [
        "arn:aws:s3:::farto.cloud/${statement.key}/*",
      ]
      principals {
        type = "AWS"
        identifiers = [
          statement.value,
        ]
      }
    }
  }
}

resource "aws_s3_bucket" "farto_cloud" {
  bucket = "farto.cloud"
  acl    = "private"
}

resource "aws_s3_bucket_policy" "farto_cloud" {
  bucket = aws_s3_bucket.farto_cloud.id
  policy = data.aws_iam_policy_document.farto_cloud.json
}

module "cloudfront_mom" {

  source = "./modules/cloudfront"

  aliases            = ["mom.farto.cloud"]
  bucket_domain_name = aws_s3_bucket.farto_cloud.bucket_domain_name
  origin_id          = "S3-farto.cloud/mom"
  origin_path        = "/mom"
  subdomain          = "mom"
}

module "cloudfront_test" {

  source = "./modules/cloudfront"

  aliases            = ["test.farto.cloud"]
  bucket_domain_name = aws_s3_bucket.farto_cloud.bucket_domain_name
  origin_id          = "S3-farto.cloud/test"
  origin_path        = "/test"
  subdomain          = "test"
}
