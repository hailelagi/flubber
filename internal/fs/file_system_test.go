//go:build !darwin

package fs

import (
	"os"
	"testing"

	"github.com/hailelagi/flubber/internal/config"
	"github.com/hanwen/go-fuse/v2/fs"
)

func initBlockFS(t *testing.T) (mountPoint string, cleanup func()) {
	config.SetupTestConfig()
	root, err := NewBlockFileSystem("/mnt")
	if err != nil {
		t.Fatalf("new block fs creation failed: %v", err)
	}

	mountPoint = t.TempDir()
	opts := &fs.Options{}
	server, _ := fs.Mount(mountPoint, root, opts)

	return mountPoint, func() {
		_ = server.Unmount()
	}
}

func TestNewBlockFS(t *testing.T) {
	mountPoint, clean := initBlockFS(t)
	defer clean()

	_, err := os.ReadDir(mountPoint)

	if err != nil {
		t.Fatalf("ReadDir failed: %v", err)
	}
}
