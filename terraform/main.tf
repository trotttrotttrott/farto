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
