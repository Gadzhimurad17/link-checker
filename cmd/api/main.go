package main

import (
	"context"

	f "github.com/Gadzhimurad17/link-checker/internal/app/factory"

	"github.com/Gadzhimurad17/link-checker/internal/config"
)

func main() {
	cfg := config.LoadFromEnv()
	factory := f.NewFactory(cfg)
	app := factory.NewApp()
	ctx := context.Background()

	go app.Start()
	app.Stop(ctx, cfg.ShutdownPeriod, cfg.ReadinessDrainDelay)
}
