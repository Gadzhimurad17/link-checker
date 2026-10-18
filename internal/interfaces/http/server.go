package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

type Server struct {
	server   *http.Server
	handlers *Handlers
	//middleware *middleware
}

func NewServer(host string, port int, handlers *Handlers,

// middleware *middleware
) *Server {
	return &Server{
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", host, port),
		},
		handlers: handlers,
		//middleware: middleware,
	}
}

func (s *Server) Start() {
	l := slog.With(slog.String("server", "http"))

	mux := http.NewServeMux()
	s.RegisterRoutes(mux)

	// Подключаем маршруты к HTTP-серверу
	s.server.Handler = mux

	l.Info("Starting server", slog.String("addr", s.server.Addr))

	if err := s.server.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("can't start http server, error=%s", err)
	}
}

func (s *Server) Close(ctx context.Context) {
	slog.InfoContext(ctx, "Stopping http server on", slog.String("address", s.server.Addr))

	s.server.Shutdown(ctx)
}
