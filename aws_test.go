package main

import (
	"testing"
)

func TestWalkBucket(t *testing.T) {
	b, err := walkBucket("farto.cloud")
	if err != nil {
		t.Errorf("Unexpected error walking bucket: %s", err)
	}
	if b != "" {
		t.Errorf("Shiiiit: %s", b)
	}
}
