package storage

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type ObjectStorage interface {
	Get(ctx context.Context, offset uint64) ([]byte, error)
	Append(ctx context.Context, offset uint64) error
	Scan(ctx context.Context, offset uint64) ([][]byte, error)
}

type WalTxn struct {
	wal       *FSWal
	id        uint64
	committed atomic.Bool
	timestamp time.Duration
	mu        sync.Mutex
	events    []Event
	ctx       context.Context
}

var (
	ErrTxnNotFound      = errors.New("txn not found")
	ErrTxnAlreadyExists = errors.New("txn duplicate")
	ErrTxnCommitted     = errors.New("txn committed")
	ErrKeyNotFound      = errors.New("key not found")
)

func NewTxn(ctx context.Context, wal *FSWal) *WalTxn {
	return &WalTxn{
		wal:       wal,
		id:        wal.prevTxnId.Add(1),
		timestamp: time.Duration(time.Now().UnixNano()),
		events:    make([]Event, 0),
		ctx:       ctx,
	}
}

func (t *WalTxn) Begin() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.wal.txnsMu.Lock()
	defer t.wal.txnsMu.Unlock()

	if _, exists := t.wal.txns[t.id]; exists {
		return ErrTxnAlreadyExists
	}

	t.wal.txns[t.id] = t
	return nil
}

func (t *WalTxn) Get(id uint64) (Entry, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// todo: hit local cache

	b, err := t.wal.store.Get(context.Background(), id)

	if err != nil {
		return Entry{buffer: b}, err
	} else {
		return Entry{}, err
	}
}

func (t *WalTxn) Put(id uint64, value []byte) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.committed.Load() {
		return ErrTxnCommitted
	}

	originalEntry, err := t.Get(id)
	if err != nil && !errors.Is(err, ErrKeyNotFound) {
		return err
	}

	t.events = append(t.events, Event{
		Type:     OperationPut,
		Original: &originalEntry,
	})

	return nil
}

func (t *WalTxn) Delete(id uint64) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.committed.Load() {
		return ErrTxnCommitted
	}

	originalEntry, err := t.Get(id)
	if err != nil {
		return err
	}

	t.events = append(t.events, Event{
		Type:     OperationDelete,
		Original: &originalEntry,
	})

	return nil
}

func (t *WalTxn) Commit() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.committed.Load() {
		return ErrTxnCommitted
	}

	// Phase 1: Prepare
	// todo: retry logic
	// Phase 2: Commit
	// Mark transaction as committed
	// Cleanup prepare marker

	t.committed.Store(true)
	// delete(t.wal.txns, t.id)
	return nil
}

// Rollback undoes all operations in the transaction
func (t *WalTxn) Rollback() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.committed.Load() {
		return ErrTxnCommitted
	}

	return nil
}
