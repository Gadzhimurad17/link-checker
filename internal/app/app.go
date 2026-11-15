package app

import (
	"context"
	"linkcheker/internal/interfaces/http"
	"time"
)

type App struct {
	Server *http.Server
}

func (app *App) Start() {
	app.Server.Start()
}

func (app *App) Stop(ctx context.Context, shutdownPeriod time.Duration, readinessDrainDelay time.Duration) {
	app.Server.Shutdown(ctx, shutdownPeriod, readinessDrainDelay)
}
