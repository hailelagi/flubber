package internal

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Entry struct {
	offset uint64
	buffer []byte
	txn    *WalTxn
}

// todo populate with object store methods/fields
type FSWal struct{}

func NewFSWal(client *s3.Client, bucketName, prefix string) *FSWal {
	return &FSWal{}
}

func (w *FSWal) Append(ctx context.Context, data []byte) (uint64, error) {
	return 0, nil
}

func (w *FSWal) Read(ctx context.Context, offset uint64) (Entry, error) {
	return Entry{}, nil
}
