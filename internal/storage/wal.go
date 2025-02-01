package storage

import (
	"context"
	"sync/atomic"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Tuple struct {
	Offset uint64
	Buffer []byte
}

type Event struct {
	Type      EventType
	Tombstone bool
	Original  *Tuple
}

type EventType int

const (
	OperationPut EventType = iota + 1
	OperationDelete
)

type FSWal struct {
	Storage

	// todo: order
	Txns              map[uint64]*Transaction
	baseDir           string
	LastCommitedTxnId atomic.Uint64
}

// todo build high level put/del over append

func NewFSWal(client *s3.Client, bucketName, prefix string) *FSWal {
	// todo: init store interface
	return &FSWal{baseDir: bucketName, LastCommitedTxnId: atomic.Uint64{}}
}

func (w *FSWal) Append(ctx context.Context, data []byte) (uint64, error) {
	return 0, nil
}

func (w *FSWal) Read(ctx context.Context, offset uint64) (Tuple, error) {
	return Tuple{}, nil
}
