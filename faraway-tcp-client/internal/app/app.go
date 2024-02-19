package app

import (
	"context"
	"log"

	"github.com/alenIncorC/faraway-pow/faraway-tcp-client/config"
	"github.com/alenIncorC/faraway-pow/faraway-tcp-client/internal/usecase"
	"github.com/alenIncorC/faraway-pow/faraway-tcp-client/pkg/solver/hashcash"
)

func Run(ctx context.Context, cfg *config.Config) {
	powProvider := hashcash.NewSolver()

	client := usecase.NewClient(cfg, powProvider)
	if err := client.Start(ctx, cfg.RequestCount); err != nil {
		log.Fatalf("config init error occured %s", err)
	}
}
