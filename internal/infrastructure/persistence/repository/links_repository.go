package repository

import (
	"context"
	"linkcheker/internal/domain/models"
	linksStorage "linkcheker/internal/infrastructure/storage"
)

type LinkRepo struct {
	storage *linksStorage.Links
}

func NewRepository(storage *linksStorage.Links) *LinkRepo {
	return &LinkRepo{
		storage: storage,
	}
}

func (r *LinkRepo) CreateLink(ctx context.Context, link string) error {
	r.storage.StoreLink(link)
	return nil
}

func (r *LinkRepo) CheckStatus(ctx context.Context, link string) (bool, error) {
	_, err := r.storage.LoadLink(link)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *LinkRepo) DeleteLink(ctx context.Context, link string) {
	r.storage.LoadLink(link)
}

func (r *LinkRepo) FindLinks() []models.LinkStat {
	links := r.storage.ListLinks()
	return links
}
