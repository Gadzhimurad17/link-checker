package factory

import (
	"sync"

	"github.com/Gadzhimurad17/link-checker/internal/domain/ports/repo"

	"github.com/Gadzhimurad17/link-checker/internal/infrastructure/persistence/repository"
	"github.com/Gadzhimurad17/link-checker/internal/infrastructure/storage"
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
