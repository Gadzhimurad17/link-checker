package http

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"
	"time"
)

const (
	shutdownPeriod      = 15 * time.Second
	readinessDrainDelay = 5 * time.Second
)

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

func (s *Server) Shutdown(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	stop()

	s.handlers.isShuttingDown.Store(true)
	time.Sleep(readinessDrainDelay)

	shutdownCtx, cancel := context.WithTimeout(
		context.WithoutCancel(ctx),
		shutdownPeriod,
	)
	defer cancel()

	s.Close(shutdownCtx)
}
