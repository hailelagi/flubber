package storage

import "context"

type Storage interface {
	Get(ctx context.Context, offset uint64) ([]byte, error)
	Append(ctx context.Context, offset uint64) error
	Scan(ctx context.Context, offset uint64) ([][]byte, error)
}

type Transaction interface {
	NewTxn(ctx context.Context, wal *FSWal) *Transaction
	Begin() error
	Commit() error
}
