# Farto

Static site generator for browsing fartos on S3.

## Commands

```
farto site generate
```

Generate static site locally.

```
farto site publish
```

Push your static site to S3.

```
farto fartos normalize <path>
```

Create normalized versions of your fartos. As in, create new image files from
the originals that are a consistent size and format.

`path` must be local. The images will be created in sibling directories:

* `<path>.farto.n`
* `<path>.farto.n.t`

```
farto fartos upload <path>
```

Uploads fartos in `path` and also uploads normalized sibling directories if
present.

## Config

TODO

## Access Control

Basic auth with Lambda & Cloudfront. Links:

* [Blog post](https://medium.com/hackernoon/serverless-password-protecting-a-static-website-in-an-aws-s3-bucket-bfaaa01b8666)
* [IAM permissions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-edge-permissions.html)

## Cost Considerations

[S3 is pretty cheap](https://aws.amazon.com/s3/pricing/). At the time of writing
this, at most it's $0.023 per GB a month. A dollar and change for 100 GB a month.
