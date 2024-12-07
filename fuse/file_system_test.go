package fuse

import (
	"testing"

	"go.uber.org/goleak"
)

func TestHello(t *testing.T) {
	t.Fail()
}

func TestFS(m *testing.M) {
	goleak.VerifyTestMain(m)
}
