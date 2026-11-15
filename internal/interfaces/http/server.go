package http

import (
	"context"
	"errors"
	"fmt"
	"linkcheker/internal/interfaces/http/handlers"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

type Server struct {
	server         *http.Server
	handlers       *handlers.Handlers
	middleware     *Middleware
	cancelChan     chan struct{}
	isShuttingDown atomic.Bool
}

func NewServer(host string, port int, handlers *handlers.Handlers, middleware *Middleware) *Server {
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
	l := slog.With(slog.String("server", "http"))

	mux := http.NewServeMux()
	s.RegisterRoutes(mux)

	// Подключаем маршруты к HTTP-серверу
	handler := s.middleware.LogRequests(mux)
	s.server.Handler = s.WithGracefulShutdown(handler)

	l.Info("Starting server", slog.String("addr", s.server.Addr))

	if err := s.server.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("can't start http server, error=%s", err)
	}
}

func (s *Server) Close(ctx context.Context) {
	slog.InfoContext(ctx, "Stopping http server on", slog.String("address", s.server.Addr))

	if err := s.server.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "Failed to shutdown HTTP server",
			slog.Any("error: ", err),
		)

		close(s.cancelChan)

		if err := s.server.Close(); err != nil {
			slog.ErrorContext(ctx, "Failed to close HTTP server",
				slog.Any("error: ", err))
		}
	}
}

func (s *Server) Shutdown(ctx context.Context, shutdownPeriod time.Duration, readinessDrainDelay time.Duration) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	stop()

	s.isShuttingDown.Store(true)
	time.Sleep(readinessDrainDelay)

	shutdownCtx, cancel := context.WithTimeout(
		context.WithoutCancel(ctx),
		shutdownPeriod,
	)
	defer cancel()

	s.Close(shutdownCtx)
}

func (s *Server) ReadinessCheck(w http.ResponseWriter, r *http.Request) {
	if s.isShuttingDown.Load() {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) WithGracefulShutdown(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx, cancel := s.cancel(req.Context())
		defer cancel()

		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}

func (s *Server) cancel(parentCtx context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parentCtx)

	go func() {
		select {
		case <-s.cancelChan:
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}
