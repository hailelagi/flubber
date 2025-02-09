package fs

import (
	"context"
	"syscall"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

/*
** API design considerations
how to map a stateless protocol over a stateful api?
volume identifier, an inode number, and a generation number
*/

/*
// ~stolen~ borrowed from polarFS
int     pfs_mount(const char *volname, int host_id)
int     pfs_umount(const char *volname)
int     pfs_mount_growfs(const char *volname)

int     pfs_creat(const char *volpath, mode_t mode)
int     pfs_open(const char *volpath, int flags, mode_t mode)
int     pfs_close(int fd)
ssize_t pfs_read(int fd, void *buf, size_t len)
ssize_t pfs_write(int fd, const void *buf, size_t len)
off_t   pfs_lseek(int fd, off_t offset, int whence)
ssize_t pfs_pread(int fd, void *buf, size_t len, off_t offset)
ssize_t pfs_pwrite(int fd, const void *buf, size_t len, off_t offset)
int     pfs_stat(const char *volpath, struct stat *buf)
int     pfs_fstat(int fd, struct stat *buf)
int     pfs_posix_fallocate(int fd, off_t offset, off_t len)
int     pfs_unlink(const char *volpath)
int     pfs_rename(const char *oldvolpath, const char *newvolpath)
int     pfs_truncate(const char *volpath, off_t len)
int     pfs_ftruncate(int fd, off_t len)
int     pfs_access(const char *volpath, int amode)

int     pfs_mkdir(const char *volpath, mode_t mode)
DIR*    pfs_opendir(const char *volpath)
struct dirent *pfs_readdir(DIR *dir)
int     pfs_readdir_r(DIR *dir, struct dirent *entry,

	struct dirent **result)

int     pfs_closedir(DIR *dir)
int     pfs_rmdir(const char *volpath)
int     pfs_chdir(const char *volpath)
int     pfs_getcwd(char *buf)
*/
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
				Data: []byte(""),
				Attr: fuse.Attr{
					Mode: 0644,
				},
			}, fs.StableAttr{Ino: 2})
		r.AddChild("hello.txt", ch, false)
	}
}

func (r *flubberRoot) Getattr(ctx context.Context, fh fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
	out.Mode = 0755
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
	out.Mode = 0755
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
