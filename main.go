package main

import (
	"flag"
	"log"

	"github.com/hailelagi/flubber/cmd"
	"github.com/hailelagi/flubber/fuse"
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

var _ = (fs.NodeGetattrer)((*fuse.HelloRoot)(nil))
var _ = (fs.NodeOnAdder)((*fuse.HelloRoot)(nil))

func main() {
	cmd.Execute()

	debug := flag.Bool("debug", false, "print debug data")

	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n  hello MOUNTPOINT")
	}

	opts := &fs.Options{}
	opts.Debug = *debug

	server, err := fs.Mount(flag.Arg(0), &fuse.HelloRoot{}, opts)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	server.Wait()
}
