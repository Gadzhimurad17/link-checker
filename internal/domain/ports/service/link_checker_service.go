package service

import (
	"context"
	"linkcheker/internal/domain/models"
)

type LinkCheckerService interface {
	CheckLinks(ctx context.Context, links []string) []models.Link
	PopularLinks(limit int64) []models.LinkStat
}
