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
the originals that are a similar size and same encoding.

`path` must be local. The images will be created in sibling directories:

* `<path>-n`
* `<path>-n-sm`

## Config

TODO

## Access Control

Lambda & Cloudfront?

## Cost Considerations

[S3 is pretty cheap](https://aws.amazon.com/s3/pricing/). At the time of writing
this, at most it's $0.023 per GB a month. A dollar and change for 100 GB a month.
