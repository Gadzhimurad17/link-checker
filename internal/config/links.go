package config

type LinksConfig struct {
	MinLinksAmount int `env:"MIN_LINKS_AMOUNT"`
	MaxLinksAmount int `env:"MAX_LINKS_AMOUNT"`
}
