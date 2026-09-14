package repository

import (
	linksStorage "linkcheker/internal/domain/storage"
)

type LinkRepository struct {
	storage *linksStorage.Links
}

func NewRepository(storage *linksStorage.Links) *LinkRepository {
	return &LinkRepository{
		storage: storage,
	}
}
