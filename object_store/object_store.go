package objectstore

import (
	"context"
	"io"

	"fmt"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/s3blob"
)

/*
provides a mapping between the filesystem object, the inode, and the
dir object with the blob?
*/

/*
func NewBucket() {
	ctx = context.Background()
	bucket, err := s3blob.OpenBucket(ctx, sess, "mybucket")

	if err != nil {
		return bucket
	}
}
*/

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
