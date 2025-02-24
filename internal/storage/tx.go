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
	ACTIVE    TxStatus = 1 // tx begin
	FAILED                 // tx abort
	ABORTED                // tx cancel
	EXECUTED               // tx processed but not replicated
	COMMITTED              // tx 2pc replicated
)

type WalTx struct {
	id        uint64
	beginAt   time.Duration
	endAt     time.Duration
	status    TxStatus
	committed atomic.Bool

	tuples       []Tuple
	ctx          context.Context
	cancellation context.CancelFunc
}

var timeout time.Duration = 5 * time.Second

// var _ Transaction = (*WalTx)(nil)

func NewWalTxn(ctx context.Context, wal *FSWal) *WalTx {
	ctx, cancel := context.WithTimeout(ctx, timeout)

	return &WalTx{
		id:           wal.LastCommitedTxnId.Add(1),
		status:       ACTIVE,
		tuples:       make([]Tuple, 0),
		ctx:          ctx,
		cancellation: cancel,
	}
}

// Retrieve a transaction yet to be applied
func (t *WalTx) Get(id uint64, s Storage) (Tuple, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	b, err := s.Get(ctx, id)

	if err != nil {
		return Tuple{Buffer: b}, err
	} else {
		return Tuple{}, err
	}
}

// Insert a running/active transaction to the list of to be applied txs
func (t *WalTx) Put(id uint64, s Storage, value []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if t.committed.Load() {
		return ErrTxnCommitted
	}

	_, err := s.Get(ctx, id)

	if err != nil && !errors.Is(err, ErrKeyNotFound) {
		return err
	}

	/*
		todo: apply over wal
				t.events = append(t.events, Event{
				Type:     OperationPut,
				Original: &originalTuple,
			})
	*/

	return nil
}

// Kill a running/ACTIVE tx
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

// func (t *WalTx) Begin() error {
// 	if _, exists := t.Txns[t.id]; exists {
// 		return ErrTxnAlreadyExists
// 	}

// 	t.beginAt = time.Duration(time.Now().UnixNano())
// 	// s.Txns[t.id] = t
// 	return nil
// }

// func (t *WalTx) Commit() error {
// 	if t.committed.Load() {
// 		return ErrTxnCommitted
// 	}

// 	// Phase 1: Prepare
// 	// todo: retry logic
// 	// Phase 2: Commit
// 	// Mark transaction as committed
// 	// Cleanup prepare marker
// 	t.endAt = time.Duration(time.Now().UnixNano())
// 	t.committed.Store(true)
// 	// delete(t.wal.txns, t.id)
// 	return nil
// }

// // Rollback undoes all operations in the transaction
// func (t *WalTx) Rollback() error {
// 	if t.committed.Load() {
// 		return ErrTxnCommitted
// 	}

// 	return nil
// }
