package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	ConcurrencyConfig
	LinksConfig
	AppConfig
	UserConfig
}

func LoadFromEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("can't load config settings from .env, error=%s", err)
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("can't parse config from environment, error=%s", err)
	}

	return cfg
}
