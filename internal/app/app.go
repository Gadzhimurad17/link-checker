package app

import (
	"context"
	"linkcheker/internal/interfaces/http"
)

type App struct {
	Server *http.Server
}

func (app *App) Start() {
	app.Server.Start()
}

func (app *App) Stop(ctx context.Context) {
	app.Server.Close(ctx)
}
