data "aws_iam_policy_document" "farto_lambda" {
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

resource "aws_iam_role" "farto_lambda" {
  assume_role_policy = data.aws_iam_policy_document.farto_lambda.json
  path               = "/service-role/"
}
