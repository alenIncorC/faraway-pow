package app

import (
	"context"

	"github.com/alenIncorC/faraway-pow/faraway-tcp-server/config"
	"github.com/alenIncorC/faraway-pow/faraway-tcp-server/internal/repository"
	"github.com/alenIncorC/faraway-pow/faraway-tcp-server/internal/usecase"
	"github.com/alenIncorC/faraway-pow/faraway-tcp-server/pkg/challenger/challenge"
	"github.com/alenIncorC/faraway-pow/faraway-tcp-server/pkg/verifier/verification"
	"github.com/alenIncorC/faraway-pow/libs/utils"
)

const (
	ErrVerifierInit   = "failed to initialize pow verification"
	ErrChallengerInit = "failed to initialize pow challenger"
	ErrRunServer      = "failed server run"
)

// RunServer started server application
func Run(ctx context.Context, cfg *config.Config) {
	verifier, err := verification.NewPOWPOWVerifier(cfg.Pow.Difficulty)
	if err != nil {
		utils.FatalApplication(ErrVerifierInit, err)
	}

	challenger, err := challenge.NewChallenge(cfg.Pow.Difficulty)
	if err != nil {
		utils.FatalApplication(ErrChallengerInit, err)
	}

	repo := repository.NewRepositories()
	server := usecase.NewServer(cfg, verifier, challenger, repo)

	if err = server.Run(ctx); err != nil {
		utils.FatalApplication(ErrRunServer, err)
	}
}
