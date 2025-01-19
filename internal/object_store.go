package internal

import (
	"context"
	"fmt"
	"io"

	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"
	// "gocloud.dev/blob/s3blob"
)

var storeClient *s3blob.URLOpener

func FormatBucket(imageName string, blockSize, pageSize int) {
	// Placeholder function to format the bucket

	/*
		bucketURL := viper.GetString("bucket.url")
		bucketName := viper.GetString("bucket.name")
		accessKeyId = viper.GetString("credentials.access_key_id")
		secretAccessKey := viper.GetString("credentials.secret_access_key")
	*/
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
