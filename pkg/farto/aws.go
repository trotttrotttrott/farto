package farto

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

func walkBucket(svc s3iface.S3API, bucket string, prefix string) (keys []string, err error) {
	resp, err := svc.ListObjectsV2(
		&s3.ListObjectsV2Input{
			Bucket: aws.String(bucket),
			Prefix: aws.String(prefix),
		},
	)
	if err != nil {
		return
	}
	for _, item := range resp.Contents {
		keys = append(keys, *item.Key)
	}
	return
}
