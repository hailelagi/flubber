package storage

import (
	"context"
	"errors"
	"sync/atomic"
	"time"
)

var (
	ErrTxnNotFound      = errors.New("txn not found")
	ErrTxnAlreadyExists = errors.New("txn duplicate")
	ErrTxnCommitted     = errors.New("txn committed")
	ErrKeyNotFound      = errors.New("key not found")
)

type TxLevel uint8
type TxStatus uint8

const (
	READ_UNCOMMITTED TxLevel = 1
	SERIALIZABLE
)

const (
	RUNNING TxStatus = 1
	ABORTED
	COMMITTED
)

type WalTx struct {
	id        uint64
	beginAt   time.Duration
	endAt     time.Duration
	status    TxStatus
	committed atomic.Bool

	tuples []Tuple
	ctx    context.Context
}

// var _ storage.Transaction = (*WalTx)(nil)

func NewWalTxn(ctx context.Context, wal *FSWal) *WalTx {
	return &WalTx{
		id:     wal.LastCommitedTxnId.Add(1),
		tuples: make([]Tuple, 0),
		ctx:    ctx,
	}
}

func (t *WalTx) Begin() error {
	/*
		if _, exists := s.Txns[t.id]; exists {
			return ErrTxnAlreadyExists
		}

	*/

	t.beginAt = time.Duration(time.Now().UnixNano())
	// s.Txns[t.id] = t
	return nil
}

func (t *WalTx) Get(id uint64, s Storage) (Tuple, error) {
	// todo: hit local cache

	b, err := s.Get(context.Background(), id)

	if err != nil {
		return Tuple{Buffer: b}, err
	} else {
		return Tuple{}, err
	}
}

func (t *WalTx) Put(id uint64, value []byte) error {
	/*
		if t.committed.Load() {
			return ErrTxnCommitted
		}

		originalTuple, err := t.Get(id)
		if err != nil && !errors.Is(err, ErrKeyNotFound) {
			return err
		}

		t.events = append(t.events, Event{
			Type:     OperationPut,
			Original: &originalTuple,
		})
	*/

	return nil
}

func (t *WalTx) Delete(id uint64) error {
	/*
		if t.committed.Load() {
			return ErrTxnCommitted
		}

		originalTuple, err := t.Get(id)
		if err != nil {
			return err
		}

		t.events = append(t.events, Event{
			Type:     OperationDelete,
			Original: &originalTuple,
		})
	*/

	return nil
}

func (t *WalTx) Commit() error {
	if t.committed.Load() {
		return ErrTxnCommitted
	}

	// Phase 1: Prepare
	// todo: retry logic
	// Phase 2: Commit
	// Mark transaction as committed
	// Cleanup prepare marker
	t.endAt = time.Duration(time.Now().UnixNano())
	t.committed.Store(true)
	// delete(t.wal.txns, t.id)
	return nil
}

// Rollback undoes all operations in the transaction
func (t *WalTx) Rollback() error {
	if t.committed.Load() {
		return ErrTxnCommitted
	}

	return nil
}
