package fs

import (
	"context"

	"github.com/hailelagi/flubber/internal/config"
	"github.com/hailelagi/flubber/internal/metrics"
	"github.com/hailelagi/flubber/internal/storage"
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
	"go.uber.org/zap"
)

type flubberRoot struct {
	client *storage.StoreClient
	fs.Inode
}

func (r *flubberRoot) Close(ctx *context.Context) {
	// todo cleanup
}

func NewBlockFileSystem(name string) (fs.InodeEmbedder, error) {
	config := config.GetStorageConfig()
	client := storage.InitObjectStoreClient(config)

	return &flubberRoot{client: client}, nil
}

func InitMount(mountpoint string, config *config.Mount) error {
	opts := &fs.Options{
		AttrTimeout:  config.Ttl,
		EntryTimeout: config.Ttl,
		MountOptions: fuse.MountOptions{
			Debug: config.Debug,
		},
	}

	root, err := NewBlockFileSystem(opts.Name)

	if err != nil {
		zap.L().Fatal("root block creation failed:", zap.Error(err))
	}

	server, err := fs.Mount(mountpoint, root, opts)

	if err != nil {
		zap.L().Fatal("mount failure", zap.Error(err))
	}

	metrics.StartMetricsServer()
	metrics.StartMetricsPprof(config)

	server.Wait()

	return nil
}
