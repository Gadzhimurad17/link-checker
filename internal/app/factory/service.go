package factory

import (
	"github.com/Gadzhimurad17/link-checker/internal/domain/ports/repo"

	"github.com/Gadzhimurad17/link-checker/internal/domain/ports/service"
	"github.com/Gadzhimurad17/link-checker/internal/infrastructure/services"
)

func (f *Factory) newServices(repo repo.LinkRepository) service.LinkCheckerService {
	return services.NewLinkCheckerService(f.cfg.MaxLinksAmount, repo)
}
