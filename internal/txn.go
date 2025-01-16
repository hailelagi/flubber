package internal

import (
	"sync"
	"sync/atomic"
	"time"
)

type WalTxn struct {
	wal       *FSWal
	committed atomic.Bool
	timestamp time.Duration
	mu        sync.Mutex
}

func (*WalTxn) Get() (Entry, error) {
	return Entry{}, nil
}

func (*WalTxn) Put() (Entry, error) {
	return Entry{}, nil
}

func (*WalTxn) Scan() (Entry, error) {
	return Entry{}, nil
}

func (*WalTxn) Del() (Entry, error) {
	return Entry{}, nil
}

func (*WalTxn) Commit() (Entry, error) {
	return Entry{}, nil
}
