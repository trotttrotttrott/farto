data "aws_iam_policy_document" "farto_cloud_s3" {

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

data "aws_iam_policy_document" "farto_cloud_lambda" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type = "Service"
      identifiers = [
        "edgelambda.amazonaws.com",
        "lambda.amazonaws.com",
      ]
    }
  }
}
