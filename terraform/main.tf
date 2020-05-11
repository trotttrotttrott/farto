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

provider "aws" {
  region = "us-east-1"
  alias  = "us_east_1"
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

# Use this when you can figure out how to set dynamic basic auth creds.
#
# resource "aws_lambda_function" "farto_cloud_auth" {
#   provider         = aws.us_east_1
#   function_name    = "fartoCloudAuth"
#   filename         = "fartoCloudAuth.zip"
#   role             = aws_iam_role.farto_cloud_lambda.arn
#   handler          = "index.handler"
#   source_code_hash = filebase64sha256("fartoCloudAuth.zip")
#   runtime          = "nodejs12.x"
# }

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
