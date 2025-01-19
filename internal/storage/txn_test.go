package storage

import (
	"context"
	"testing"

	"github.com/hailelagi/flubber/internal/config"
	"github.com/spf13/viper"
)

func setupConfig() {
	viper.Set("bucket.url", "localhost:9000")
	viper.Set("bucket.name", "test")
	viper.Set("credentials.access_key_id", "minioadmin")
	viper.Set("credentials.secret_access_key", "minioadmin")

}

func TestBeginTxn(t *testing.T) {
	setupConfig()
	ctx := context.Background()
	config := config.GetStorageConfig()
	store := InitObjectStoreClient(config)

	// todo create wal segment
	// store.ComposeObject()

	wal := &FSWal{
		store: store,
		txns:  make(map[uint64]*WalTxn),
	}

	t.Run("begin and commit", func(t *testing.T) {
		t1 := NewTxn(ctx, wal)
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
		txn := NewTxn(ctx, wal)
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
