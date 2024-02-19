package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alenIncorC/faraway-pow/faraway-tcp-server/config"
	"github.com/alenIncorC/faraway-pow/faraway-tcp-server/internal/app"
	"github.com/alenIncorC/faraway-pow/libs/utils"
)

const ErrConfigInit = "failed config initialization"

func main() {
	ctx, cancel := signal.NotifyContext(context.TODO(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ctx.Done()
		cancel()
	}()

	cfg, err := config.NewConfig(ctx, config.Config{})
	if err != nil {
		utils.FatalApplication(ErrConfigInit, err)
	}

	app.Run(ctx, cfg)
}
