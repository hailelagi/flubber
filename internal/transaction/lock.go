package transaction

import (
	"sync"
	"sync/atomic"
)

type lockType uint8

const (
	SHARED = 1 << iota
	EXCLUSIVE
	UNLOCKED
)

type TwoPhaseLock struct {
	phase
	statistics atomic.Uint64
}

type phase struct {
	lock    sync.Locker
	t       lockType
	growing bool
}

func (l *TwoPhaseLock) Lock()   {}
func (l *TwoPhaseLock) Unlock() {}
