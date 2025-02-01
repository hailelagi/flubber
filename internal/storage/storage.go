package storage

import (
	"context"
)

type Storage interface {
	Get(ctx context.Context, offset uint64) ([]byte, error)
	Append(ctx context.Context, data []byte) (uint64, error)
	Scan(ctx context.Context, offset uint64) ([][]byte, error)
}

type Transaction interface {
	Begin() error
	Commit() error
	Rollback() error
}

type WalOperation interface {
	Undo() error
	Redo() error
}
