package factory

import (
	"linkcheker/internal/config"
	"linkcheker/internal/interfaces/http"
)

type Factory struct {
	cfg *config.Config
}

func NewFactory(cfg *config.Config) *Factory {
	return &Factory{
		cfg: cfg,
	}
}

func (f *Factory) NewHandlers() *http.Handlers {
	// repos
	repo := f.newRepos()
	// services
	linkService := f.newServiсes(repo)
	return http.NewHandlers(linkService)
}
func (f *Factory) NewMiddleware() *http.Middleware {
	return http.NewMiddleware(f.cfg.ApiKey)
}
func (f *Factory) NewHttpServer(handlers *http.Handlers, middleware *http.Middleware) *http.Server {
	host := f.cfg.AppHost
	port := f.cfg.AppPort
	return http.NewServer(host, port, handlers, middleware)
}
