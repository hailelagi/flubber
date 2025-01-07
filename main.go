//go:build linux || darwin

package main

import (
	"log"

	"github.com/hailelagi/flubber/cmd"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	sugar := logger.Sugar()

	if err := cmd.Execute(); err != nil {
		sugar.Fatal(err)
	}
}
