package storage

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hailelagi/flubber/internal/config"
)

func TestBeginTxn(t *testing.T) {
	config.SetupTestConfig()
	ctx := context.Background()
	// config := config.GetStorageConfig()
	// tore := storage.InitObjectStoreClient(config)

	// todo create wal segment
	// store.ComposeObject()

	wal := NewFSWal(&s3.Client{}, "test-bucket", "test")

	t.Run("begin and commit", func(t *testing.T) {
		t1 := NewWalTxn(ctx, wal)
		err := t1.Begin()
		if err != nil {
			t.Errorf("failed to start txn %v", err)
		}

		err = t1.Commit()
		if err != nil {
			t.Fatalf("did not commit")
		}
	})

	t.Run("begin duplicate", func(t *testing.T) {
		txn := NewWalTxn(ctx, wal)
		err := txn.Begin()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		err = txn.Begin()
		if err != ErrTxnAlreadyExists {
			t.Errorf("Expected ErrTxnAlreadyExists, got %v", err)
		}
	})
}
