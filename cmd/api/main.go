package main

import (
	f "linkcheker/internal/app/factory"
	"linkcheker/internal/config"
)

func main() {
	config := config.LoadFromEnv()
	factory := f.NewFactory(config)
	app := factory.NewApp()
	app.Start()
}
