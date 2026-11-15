package main

import (
	"context"
	f "linkcheker/internal/app/factory"
	"linkcheker/internal/config"
)

func main() {
	cfg := config.LoadFromEnv()
	factory := f.NewFactory(cfg)
	app := factory.NewApp()
	ctx := context.Background()

	go app.Start()
	app.Stop(ctx, cfg.ShutdownPeriod, cfg.ReadinessDrainDelay)
}
