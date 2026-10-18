package repo

import (
	"context"
	models "linkcheker/internal/domain/models"
)

type LinkRepository interface {
	CreateLink(ctx context.Context, link *models.Link) error
	CheckStatus(ctx context.Context, link string) (bool, error)
	DeleteLink(ctx context.Context, link string)
}
