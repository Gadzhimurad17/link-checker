package handlers

import "github.com/Gadzhimurad17/link-checker/internal/domain/ports/service"

type Handlers struct {
	service service.LinkCheckerService
}

func NewHandlers(service service.LinkCheckerService) *Handlers {
	return &Handlers{
		service: service,
	}
}
