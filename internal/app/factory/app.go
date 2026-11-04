package factory

import (
	"linkcheker/internal/app"
)

func (f *Factory) NewApp() *app.App {
	linkCheckerHandler := f.NewHandlers()
	httpMiddleware := f.NewMiddleware()
	httpServer := f.NewHttpServer(linkCheckerHandler, httpMiddleware)

	return &app.App{
		Server: *httpServer,
	}

}
