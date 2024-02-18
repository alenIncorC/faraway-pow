package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alenIncorC/faraway-test/faraway-tcp-client/config"
	"github.com/alenIncorC/faraway-test/faraway-tcp-client/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.TODO(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ctx.Done()
		cancel()
	}()

	cfg, err := config.NewConfig(ctx, config.Config{})
	if err != nil {
		log.Fatalf("config init error occured %s", err)
	}

	app.Run(ctx, cfg)
}
