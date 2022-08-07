package internal

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

func NewLogger(cfg config) (*log.Logger, error) {
	logger := log.New()
	switch  cfg.Environment {
	case local, staging:
		logger.SetFormatter(&log.TextFormatter{})

	case production:
		logger.SetFormatter(&log.JSONFormatter{})

	default:
		return nil, errors.New("Unable to set logger service")
		
	}

	return logger, nil
}