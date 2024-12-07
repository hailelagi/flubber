package main

import (
	"github.com/hailelagi/flubber/cmd"
	"github.com/hailelagi/flubber/fuse"
)

func main() {
	cmd.Execute()
	fuse.Mount()
}
