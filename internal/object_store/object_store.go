package objectstore

import (
	"context"
	"fmt"
	"io"

	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"
	// "gocloud.dev/blob/s3blob"
)

var storeClient *s3blob.URLOpener

func FormatBucket(imageName, imageSize, blockSize string) {
	// Placeholder function to format the bucket
	// Implement the actual logic to format the bucket here
	fmt.Printf("Bucket formatted with image %s, size %s, and block size %s\n", imageName, imageSize, blockSize)
}

func UploadObject(ctx context.Context, bucket *blob.Bucket, key string, data []byte) error {
	w, err := bucket.NewWriter(ctx, key, nil)
	if err != nil {
		return fmt.Errorf("failed to create writer: %v", err)
	}
	defer w.Close()

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data: %v", err)
	}

	return nil
}

func DownloadObject(ctx context.Context, bucket *blob.Bucket, key string) ([]byte, error) {
	r, err := bucket.NewReader(ctx, key, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create reader: %v", err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %v", err)
	}

	return data, nil
}

func DeleteObject(ctx context.Context, bucket *blob.Bucket, key string) error {
	err := bucket.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to delete object: %v", err)
	}

	return nil
}
