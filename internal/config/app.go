package config

type AppConfig struct {
	AppHost string `env:"APP_HOST"`
	AppPort int    `env:"APP_PORT"`
}
