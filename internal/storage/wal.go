package storage

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Entry struct {
	offset    uint64
	buffer    []byte
	txn       *WalTxn
	tombstone bool
}

type Event struct {
	Type      EventType
	tombstone bool
	Original  *Entry
}

type EventType int

const (
	OperationPut EventType = iota + 1
	OperationDelete
)

type FSWal struct {
	store     Storage
	txnsMu    sync.RWMutex
	txns      map[uint64]*WalTxn
	baseDir   string
	prevTxnId atomic.Uint64
}

// todo build high level put/del over append

func NewFSWal(client *s3.Client, bucketName, prefix string) *FSWal {
	// todo: init store interface
	return &FSWal{baseDir: bucketName, prevTxnId: atomic.Uint64{}}
}

func (w *FSWal) Append(ctx context.Context, data []byte) (uint64, error) {
	return 0, nil
}

func (w *FSWal) Read(ctx context.Context, offset uint64) (Entry, error) {
	return Entry{}, nil
}
