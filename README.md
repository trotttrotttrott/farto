# Farto

Static site generator for browsing fartos on S3.

## Commands

### farto site generate

```
farto site generate
```

Generate static site locally in `site` directory.

### farto site publish

```
farto site publish
```

Push your static site to S3 from `site` directory. All files in `site` are
uploaded so you can other files if you want (css, js, or whatever).

### farto fartos normalize

```
farto fartos normalize <path>
```

Create normalized versions of your fartos. As in, create new image files from
the originals that are a consistent size and format (jpg).

`path` must be local. The images will be created in sibling directories:

* `<path>.farto.n`
* `<path>.farto.n.t`

`n` = "normalized", `t` = "thumbnail".

### farto fartos upload

```
farto fartos upload <path>
```

Uploads fartos in `path` and also uploads normalized sibling directories if
present.

## Config

Expects `farto.yaml` file.

```yaml
s3Region:
s3Bucket:
s3Prefix:
siteTitle: Farto # HTML title field
siteHeadline: A Farto Site # Content for h1 tag at top of page
siteCopy: |- # Content for p tag right below h1 tag
  Welcome to this Farto site!
```

## Access Control

Basic auth with Lambda & Cloudfront. Got the idea from [this blog
post](https://medium.com/hackernoon/serverless-password-protecting-a-static-website-in-an-aws-s3-bucket-bfaaa01b8666).

## Terraform

The [terraform](./terraform) directory contains modules for creating all
necessary AWS resources.

[terraform/main.example.tf](./terraform/main.example.tf) is an example of how
you can use them.
