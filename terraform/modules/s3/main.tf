resource "aws_s3_bucket" "farto_cloud" {
  bucket = var.bucket
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
