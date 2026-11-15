package factory

import "github.com/Gadzhimurad17/link-checker/internal/app"

func (f *Factory) NewApp() *app.App {
	linkCheckerHandler := f.NewHandlers()
	httpMiddleware := f.NewMiddleware()
	httpServer := f.NewHttpServer(linkCheckerHandler, httpMiddleware)

	return &app.App{
		Server: httpServer,
	}
}
