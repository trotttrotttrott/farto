package farto

import (
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func walkBucket(svc s3iface.S3API, bucket string, prefix string) (keys []string, err error) {
	var continuationToken *string
	for {
		resp, err := svc.ListObjectsV2(
			&s3.ListObjectsV2Input{
				Bucket:            aws.String(bucket),
				Prefix:            aws.String(prefix),
				MaxKeys:           aws.Int64(1000),
				ContinuationToken: continuationToken,
			},
		)
		if err != nil {
			return nil, err
		}
		for _, item := range resp.Contents {

			key := strings.TrimPrefix(*item.Key, prefix)
			d, f := path.Split(key)

			if !strings.HasPrefix(d, "/site/") &&
				path.Ext(f) != "" &&
				!strings.Contains(d, ".farto.") {
				keys = append(keys, key)
			}
		}
		if !aws.BoolValue(resp.IsTruncated) {
			break
		}
		continuationToken = resp.NextContinuationToken
	}
	return keys, nil
}

func upload(svc s3iface.S3API, bucket string, prefix string, localDir string, preservePath bool) error {

	uploader := s3manager.NewUploaderWithClient(svc)
	var objects []s3manager.BatchUploadObject

	err := filepath.Walk(localDir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			f, err := os.Open(p)
			if err != nil {
				return err
			}
			contentType := mime.TypeByExtension(filepath.Ext(f.Name()))
			if contentType == "" {
				buffer := make([]byte, 512)
				_, err = f.Read(buffer)
				if err != nil {
					return err
				}
				contentType = http.DetectContentType(buffer)
			}
			_, err = f.Seek(0, 0)
			if err != nil {
				return err
			}
			var key string
			if preservePath {
				key = path.Join(prefix, p)
			} else {
				key = path.Join(prefix, filepath.Base(filepath.Dir(p)), info.Name())
			}
			objects = append(
				objects,
				s3manager.BatchUploadObject{
					Object: &s3manager.UploadInput{
						Bucket:      aws.String(bucket),
						Key:         aws.String(key),
						Body:        f,
						ContentType: &contentType,
					},
				},
			)
		}
		return nil
	})

	iter := &s3manager.UploadObjectsIterator{Objects: objects}
	err = uploader.UploadWithIterator(aws.BackgroundContext(), iter)

	return err
}
