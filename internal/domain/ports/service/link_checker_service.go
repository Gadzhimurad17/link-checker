package service

import (
	"context"

	"github.com/Gadzhimurad17/link-checker/internal/domain/models"
)

type LinkCheckerService interface {
	CheckLinks(ctx context.Context, links []string) []models.Link
	PopularLinks(limit int64, offset int64) ([]models.LinkStat, error)
}
