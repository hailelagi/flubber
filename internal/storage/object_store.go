package storage

import (
	"context"
	"io"

	"fmt"

	"github.com/hailelagi/flubber/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"gocloud.dev/blob"
)

type StoreClient struct {
	*minio.Client
}

func (*StoreClient) Get(ctx context.Context, offset uint64) ([]byte, error) {
	return []byte{}, nil
}

func (*StoreClient) Append(ctx context.Context, offset uint64) error {
	return nil
}

func (*StoreClient) Scan(ctx context.Context, offset uint64) ([][]byte, error) {
	return [][]byte{}, nil
}

func InitObjectStoreClient(config *config.Storage) *StoreClient {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})

	if err != nil {
		zap.S().Fatal("could not connect to block store:", zap.Error(err))
	}

	return &StoreClient{Client: client}
}

func FormatBucket(imageName string, blockSize, pageSize int) {
	config := config.GetStorageConfig()
	client := InitObjectStoreClient(config)

	fmt.Println(client.BucketExists(context.TODO(), "test"))
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
