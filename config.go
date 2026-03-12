package main

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
)

var c Config

type Config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
}

func loadConfig(ctx context.Context) error {
	logrus.Info("Loading configuration...")

	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	return nil
}
