package fuse

import (
	"context"
	"flag"
	"syscall"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

type flubberRoot struct {
	fs.Inode
}

var (
	// syscall access method interfaces
	_ fs.NodeGetattrer = (*flubberRoot)(nil)
	_ fs.NodeOnAdder   = (*flubberRoot)(nil)
	_ fs.NodeOpener    = (*flubberRoot)(nil)
	_ fs.NodeReader    = (*flubberRoot)(nil)

	// todo on a new file node creation
	// _ = (fs.FileReader)((*flubberRoot)(nil))
	// _ = (fs.FileWriter)((*flubberRoot)(nil))
)

func Mount(mountpoint string) error {
	debug := flag.Bool("debug", true, "print debug data")

	opts := &fs.Options{
		MountOptions: fuse.MountOptions{
			Debug: *debug,
		},
	}

	server, err := fs.Mount(mountpoint, &flubberRoot{}, opts)
	if err != nil {
		return err
	}

	server.Wait()
	return nil
}

/*
todo move to during new file/inode allocation/creation
func (r *flubberRoot) OnAdd(ctx context.Context) {
	ch := r.NewPersistentInode(
		ctx, &fs.MemRegularFile{
			Data: []byte("Hello, world!\n"),
			Attr: fuse.Attr{
				Mode: 0o644,
			},
		}, fs.StableAttr{Ino: 2})
	r.AddChild("file.txt", ch, false)
}
*/

// on mount, create path traversal and hello.txt
func (r *flubberRoot) OnAdd(ctx context.Context) {
	r.AddChild(".", r.EmbeddedInode(), true)

	if _, p := r.Parent(); p != nil {
		r.AddChild("..", p, true)

		ch := r.NewPersistentInode(
			ctx, &fs.MemRegularFile{
				Data: []byte("Hello, world!\n"),
				Attr: fuse.Attr{
					Mode: 0o644,
				},
			}, fs.StableAttr{Ino: 2})
		r.AddChild("hello.txt", ch, false)
	}
}

func (r *flubberRoot) Getattr(ctx context.Context, fh fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
	out.Mode = 0o755
	return 0
}

func (r *flubberRoot) Open(ctx context.Context, flags uint32) (fs.FileHandle, uint32, syscall.Errno) {
	return nil, fuse.FOPEN_KEEP_CACHE, 0
}

func (r *flubberRoot) Read(ctx context.Context, fd fs.FileHandle, dest []byte, off int64) (fuse.ReadResult, syscall.Errno) {
	data := []byte("Hello, world!\n")
	end := off + int64(len(dest))
	if end > int64(len(data)) {
		end = int64(len(data))
	}
	return fuse.ReadResultData(data[off:end]), 0
}

func (r *flubberRoot) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	ops := flubberRoot{}
	out.Mode = 0o755
	out.Size = 42

	if name == "." {
		return r.EmbeddedInode(), 0
	}
	if name == ".." {
		if _, p := r.Parent(); p != nil {
			return p, 0
		}
		return r.EmbeddedInode(), 0
	}

	return r.NewInode(ctx, &ops, fs.StableAttr{Mode: syscall.S_IFREG}), 0
}

func (r *flubberRoot) Close(ctx context.Context) {}
