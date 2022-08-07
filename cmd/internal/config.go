package internal

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/caarlos0/env"
)

type config struct {
	ApiKey      string      `env:"ApiKey"`
	Environment environment `env:"Env"`
	// Logger *log.Logger
}

// Depedencies
type Deps struct {
	// Config config
	Logger *log.Logger
}

func GetConfig() (*config, error) {
	var cfg = &config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	if !cfg.Environment.isvalid() {
		return nil, errors.New("Enter a valid environment")
	}

	return cfg, nil
}

func GetDeps(log *log.Logger) Deps {
	deps := Deps{
		Logger: log,
	}

	return deps
}
