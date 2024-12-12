# flubber
A simple FUSE filesystem built on s3

## Installation

probably want to run this inside a docker container or an emulation of linux, or linux, or on [macfuse](https://osxfuse.github.io/) on macOS.

```
 docker build --tag docker-flubber .
```

```
docker run -p 8080:8080 docker-flubber
```

## Learn more
- https://docs.kernel.org/filesystems/vfs.html
- https://www.kernel.org/doc/html/v6.3/filesystems/fuse.html
