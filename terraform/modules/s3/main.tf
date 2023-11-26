resource "aws_s3_bucket" "farto_cloud" {
  bucket = var.bucket
}

resource "aws_s3_bucket_server_side_encryption_configuration" "farto_cloud" {
  bucket = aws_s3_bucket.farto_cloud.id

  rule {
    bucket_key_enabled = false

    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_public_access_block" "farto_cloud" {
  bucket = aws_s3_bucket.farto_cloud.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_acl" "farto_cloud" {
  bucket = aws_s3_bucket.farto_cloud.id
  acl    = "private"
}

resource "aws_s3_bucket_policy" "farto_cloud" {
  bucket = aws_s3_bucket.farto_cloud.id
  policy = data.aws_iam_policy_document.farto_cloud_s3.json
}

data "aws_iam_policy_document" "farto_cloud_s3" {

  policy_id = "PolicyForCloudFrontPrivateContent"

  dynamic "statement" {

    for_each = var.cloudfront_arns

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
