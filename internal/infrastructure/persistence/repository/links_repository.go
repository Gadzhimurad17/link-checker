package repository

import (
	"context"

	linksStorage "github.com/Gadzhimurad17/link-checker/internal/infrastructure/storage"

	"github.com/Gadzhimurad17/link-checker/internal/domain/models"
)

type LinkRepo struct {
	storage *linksStorage.Links
}

func NewRepository(storage *linksStorage.Links) *LinkRepo {
	return &LinkRepo{
		storage: storage,
	}
}

func (r *LinkRepo) CreateLink(ctx context.Context, link string) {
	r.storage.StoreLink(link)
}

func (r *LinkRepo) CheckStatus(ctx context.Context, link string) (bool, error) {
	_, err := r.storage.LoadLink(link)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *LinkRepo) DeleteLink(ctx context.Context, link string) {
	r.storage.DeleteLink(link)
}

func (r *LinkRepo) FindLinks() []models.LinkStat {
	links := r.storage.ListLinks()
	return links
}
