package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gnarlyman/dbpractice/internal/dbpractice"
)

func main() {
	app := dbpractice.NewDbPractice()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go app.Start(ctx)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh

	app.Stop(context.Background())
}
