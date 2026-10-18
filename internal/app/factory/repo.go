package factory

import (
	"linkcheker/internal/domain/ports/repo"
	"linkcheker/internal/infrastructure/persistence/repository"
	"linkcheker/internal/infrastructure/storage"
	"sync"
)

func (f *Factory) newRepos() repo.LinkRepository {
	linkRepo := f.newLinkRepo()
	return linkRepo
}

func (f *Factory) newLinkRepo() *repository.LinkRepo {
	storage := f.newStorage()
	return repository.NewRepository(storage)
}

func (f *Factory) newStorage() *storage.Links {
	links := sync.Map{}
	return storage.NewLinks(&links)
}
