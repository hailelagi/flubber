# flubber
A simple FUSE filesystem built on s3

## Installation

probably want to run this on linux or a good emulation of linux or on [macfuse](https://osxfuse.github.io/) on macOS.

Alternatively(not recommended):
```zsh
 docker build --tag docker-flubber .
```

```zsh
docker run -it --device=/dev/fuse --cap-add SYS_ADMIN \
--security-opt apparmor=unconfined -p 8080:8080 docker-flubber
```

s3 via minio:
```zsh
podman run -p 9000:9000 -p 9001:9001 \
  quay.io/minio/minio server /data --console-address ":9001"
```

## Learn more
- https://docs.kernel.org/filesystems/vfs.html
- https://www.kernel.org/doc/html/v6.3/filesystems/fuse.html

FUSE API:
- https://libfuse.github.io/doxygen/
- https://pkg.go.dev/github.com/hanwen/go-fuse/v2/fuse
- https://pkg.go.dev/github.com/hanwen/go-fuse/v2/fuse
