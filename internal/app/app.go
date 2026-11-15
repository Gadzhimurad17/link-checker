package app

import (
	"context"
	"time"

	"github.com/Gadzhimurad17/link-checker/internal/interfaces/http"
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
