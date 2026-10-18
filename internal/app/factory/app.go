package factory

import (
	"linkcheker/internal/app"
)

func (f *Factory) NewApp() *app.App {
	linkCheckerHandler := f.NewHandlers()

	httpServer := f.NewHttpServer(linkCheckerHandler)

	return &app.App{
		Server: *httpServer,
	}

}
