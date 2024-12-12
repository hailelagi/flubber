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

## Learn more
- https://docs.kernel.org/filesystems/vfs.html
- https://www.kernel.org/doc/html/v6.3/filesystems/fuse.html
