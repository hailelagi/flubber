//go:build linux || darwin

package main

import (
	"log"

	"github.com/hailelagi/flubber/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
