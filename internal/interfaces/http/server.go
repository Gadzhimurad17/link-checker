package http

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

type Server struct {
	server     *http.Server
	handlers   *Handlers
	middleware *Middleware
	cancelChan chan struct{}
}

func NewServer(host string, port int, handlers *Handlers, middleware *Middleware) *Server {
	return &Server{
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", host, port),
		},
		handlers:   handlers,
		middleware: middleware,
		cancelChan: make(chan struct{}),
	}
}

func (s *Server) Start() {
	go func() {
		l := slog.With(slog.String("server", "http"))

		mux := http.NewServeMux()
		s.RegisterRoutes(mux)

		// Подключаем маршруты к HTTP-серверу
		handler := s.middleware.LogRequests(mux)
		s.server.Handler = s.middleware.WithGracefulShutdown(handler, s.cancelChan)

		l.Info("Starting server", slog.String("addr", s.server.Addr))

		if err := s.server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("can't start http server, error=%s", err)
		}
	}()
}
