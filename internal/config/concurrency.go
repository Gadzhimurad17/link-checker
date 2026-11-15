package config

import "time"

type ConcurrencyConfig struct {
	MaxConcurrency   int           `env:"MAX_CONCURRENCY"`
	MaxLinkCheckTime time.Duration `env:"LINK_TIMEOUT"`
}
