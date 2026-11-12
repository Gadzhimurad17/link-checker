package http

import (
	"linkcheker/internal/domain/ports/service"
	"sync/atomic"
)

type Handlers struct {
	service        service.LinkCheckerService
	isShuttingDown atomic.Bool
}

func NewHandlers(service service.LinkCheckerService) *Handlers {
	return &Handlers{
		service: service,
	}
}
