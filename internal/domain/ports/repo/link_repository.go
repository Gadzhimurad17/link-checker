package repo

import (
	"context"
	"linkcheker/internal/domain/models"
)

type LinkRepository interface {
	CreateLink(ctx context.Context, link string)
	CheckStatus(ctx context.Context, link string) (bool, error)
	DeleteLink(ctx context.Context, link string)
	FindLinks() []models.LinkStat
}
