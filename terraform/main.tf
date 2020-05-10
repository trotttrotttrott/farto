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

resource "aws_s3_bucket" "farto_cloud" {
  bucket = "farto.cloud"
  acl    = "private"
}

module "cloudfront_test" {

  source = "./modules/cloudfront"

  aliases            = ["test.farto.cloud"]
  bucket_domain_name = aws_s3_bucket.farto_cloud.bucket_domain_name
  origin_id          = "S3-farto.cloud/test"
  origin_path        = "/test"
  subdomain          = "test"
}

module "cloudfront_mom" {

  source = "./modules/cloudfront"

  aliases            = ["mom.farto.cloud"]
  bucket_domain_name = aws_s3_bucket.farto_cloud.bucket_domain_name
  origin_id          = "S3-farto.cloud/mom"
  origin_path        = "/mom"
  subdomain          = "mom"
}
