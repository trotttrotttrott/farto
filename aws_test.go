package main

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type mockS3Client struct {
	s3iface.S3API
}

func (m *mockS3Client) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return nil, errors.New("barf")
}

func TestWalkBucket(t *testing.T) {

	mockSvc := &mockS3Client{}

	b, err := walkBucket(mockSvc, "farto.cloud")
	if err != nil {
		t.Errorf("Unexpected error walking bucket: %s", err)
	}
	if b != "" {
		t.Errorf("Shiiiit: %s", b)
	}
}
