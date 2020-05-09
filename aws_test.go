package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type mockS3Client struct {
	s3iface.S3API
}

func (m *mockS3Client) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	out := s3.ListObjectsV2Output{}
	return &out, nil
}

func TestWalkBucket(t *testing.T) {
	mockSvc := &mockS3Client{}
	keys, err := walkBucket(mockSvc, "farto.cloud", "test")
	if err != nil {
		t.Errorf("Unexpected error walking bucket: %s", err)
	}
	if len(keys) != 0 {
		t.Errorf("Unexpected number of keys: %d", len(keys))
	}
}
