package config

import (
	"context"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Addr      string        `env:"SERVER_ADDR,default=0.0.0.0:8000"`
	KeepAlive time.Duration `env:"SERVER_KEEP_ALIVE,default=120s"`
	Deadline  time.Duration `env:"SERVER_DEADLINE,default=120s"`
	Pow       PowConfig     `env:",prefix=POW_"`
}

type PowConfig struct {
	Difficulty int `env:"DIFFICULTY,default=4"`
}

// NewConfig generic for creates a new client or server config
func NewConfig[C any](ctx context.Context, config C) (*C, error) {
	if err := envconfig.Process(ctx, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
