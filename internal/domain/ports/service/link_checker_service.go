package service

import (
	"context"
	"linkcheker/internal/infrastructure/dto"
)

type LinkCheckerService interface {
	CheckLinks(ctx context.Context, links []string) []dto.LinkResult
}
