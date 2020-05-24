provider "aws" {
  region = "us-west-2"
}

module "s3" {

  source = "./modules/s3"

  bucket = "mybucket.com"
  cloudfront_arns = {
    test  = module.cloudfront_test.origin_access_identity_iam_arn,
    test2 = module.cloudfront_test2.origin_access_identity_iam_arn,
  }
}

module "lambda_role" {
  source = "./modules/lambda-role"
}

module "cloudfront_test" {

  source = "./modules/cloudfront"

  aliases            = ["test.mybucket.com"]
  bucket_domain_name = module.s3.bucket_domain_name
  origin_id          = "S3-mybucket.com/test"
  origin_path        = "/test"
  subdomain          = "test"

  # Leave out lambda_* if you don't want authentication.
  lambda_role_arn = module.lambda_role.arn
  lambda_auth = {
    user     = "test"
    password = "hey-test-away"
  }
}

module "cloudfront_test2" {
  # Implement this module for each Farto!
}
