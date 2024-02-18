package config

import (
	"context"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	ServerAddr   string        `env:"SERVER_ADDR,default=127.0.0.0:8000"`
	RequestCount int           `env:"CLIENT_REQUEST_COUNT,default=100"`
	KeepAlive    time.Duration `env:"CLIENT_KEEP_ALIVE,default=15s"`
	Pow          PowConfig     `env:",prefix=POW_"`
}

type PowConfig struct {
	Difficulty int `env:"DIFFICULTY,default=4"`
}

func NewConfig[C any](ctx context.Context, config C) (*C, error) {
	if err := envconfig.Process(ctx, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
