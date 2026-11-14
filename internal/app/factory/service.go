package factory

import (
	"linkcheker/internal/domain/ports/repo"
	"linkcheker/internal/domain/ports/service"
	"linkcheker/internal/infrastructure/services"
)

func (f *Factory) newServices(repo repo.LinkRepository) service.LinkCheckerService {
	return services.NewLinkCheckerService(f.cfg.MaxLinksAmount, repo)
}
