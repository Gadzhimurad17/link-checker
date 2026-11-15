package config

import "time"

type ServerConfig struct {
	ShutdownPeriod      time.Duration `env:"SHUTDOWN_PERIOD"`
	ReadinessDrainDelay time.Duration `env:"READINESS_DRAIN_DELAY"`
}
