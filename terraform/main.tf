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

resource "aws_s3_bucket_policy" "farto_cloud" {
  bucket = aws_s3_bucket.farto_cloud.id
  policy = data.aws_iam_policy_document.farto_cloud_s3.json
}

resource "aws_iam_role" "farto_cloud_lambda" {
  name               = "fartoCloudAuth-role-6drno30z"
  assume_role_policy = data.aws_iam_policy_document.farto_cloud_lambda.json
  path               = "/service-role/"
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
