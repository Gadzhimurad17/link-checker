package service

import (
	"context"
	"linkcheker/internal/domain/models"
	"linkcheker/internal/infrastructure/dto"
)

type LinkCheckerService interface {
	CheckLinks(ctx context.Context, links []string) []dto.LinkResult
	PopularLinks(limit int64) []models.LinkStat
}
