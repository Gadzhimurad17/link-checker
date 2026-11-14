package services

import (
	"context"
	"linkcheker/internal/domain/models"
	"linkcheker/internal/domain/ports/repo"
	semathore "linkcheker/internal/infrastructure"
	link "linkcheker/internal/infrastructure/dto"
	"log/slog"
	"net/http"
	"sort"
	"sync"
	"time"
)

const linkCheckTimeout = 2 * time.Second

type LinkCheckerService struct {
	sem      semathore.Semathore
	linkRepo repo.LinkRepository
}

func NewLinkCheckerService(maxLinks int, linkRepo repo.LinkRepository) *LinkCheckerService {
	return &LinkCheckerService{

		sem: semathore.Semathore{
			C: make(chan struct{}, maxLinks),
		},
		linkRepo: linkRepo,
	}
}

func (s *LinkCheckerService) CheckLinks(ctx context.Context, links []string) []link.LinkResult {
	resultch := make(chan link.LinkResult)
	var wg sync.WaitGroup
	for i := range links {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			s.sem.Acquire(1)
			defer s.sem.Release(1)
			url := links[i]

			s.linkRepo.CreateLink(ctx, url)

			status := checkLink(ctx, links[i])

			resultch <- link.LinkResult{
				URL:    links[i],
				Status: status,
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultch)
	}()
	var results []link.LinkResult

	for val := range resultch {
		results = append(results, val)
	}
	return results
}

func checkLink(parentCtx context.Context, url string) bool {
	ctx, cancel := context.WithTimeout(parentCtx, linkCheckTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.ErrorContext(ctx, "Failed to close response body",
				slog.Any("error", err),
			)
		}
	}()

	return resp.StatusCode >= 200 && resp.StatusCode < 400
}

func (s *LinkCheckerService) PopularLinks(limit int64) []models.LinkStat {
	links := s.linkRepo.FindLinks()

	sort.Slice(links, func(i, j int) bool {
		if links[i].Amount == links[j].Amount {
			return links[i].URL < links[j].URL
		}
		return links[i].Amount > links[j].Amount
	})

	links = links[:limit]
	return links
}
