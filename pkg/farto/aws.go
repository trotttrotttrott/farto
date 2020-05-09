package farto

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
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

func upload(svc s3iface.S3API, bucket string, prefix string, localDir string) error {

	f, err := os.Open(fmt.Sprintf("%s/index.html", localDir))
	if err != nil {
		return err
	}

	uploader := s3manager.NewUploaderWithClient(svc)
	objects := []s3manager.BatchUploadObject{
		{
			Object: &s3manager.UploadInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(fmt.Sprintf("%s/%s", prefix, "index.html")),
				Body:   f,
			},
		},
	}

	iter := &s3manager.UploadObjectsIterator{Objects: objects}
	err = uploader.UploadWithIterator(aws.BackgroundContext(), iter)

	return err
}
