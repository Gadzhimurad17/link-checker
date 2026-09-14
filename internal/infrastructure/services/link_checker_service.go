package services

import (
	"context"
	semathore "linkcheker/internal/infrastructure"
	link "linkcheker/internal/interfaces/http/dto"
	"net/http"
	"sync"
	"time"
)

type LinkCheckerService struct {
	sem semathore.Semathore
}

func NewLinkCheckerService(maxLinks int) *LinkCheckerService {
	return &LinkCheckerService{
		sem: semathore.Semathore{
			C: make(chan struct{}, maxLinks),
		},
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

			status := CheckLink(ctx, links[i])

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

func CheckLink(parentCtx context.Context, url string) bool {
	ctx, cancel := context.WithTimeout(parentCtx, 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://"+url, nil)
	if err != nil {
		return false
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}

	defer resp.Body.Close()
	return resp.StatusCode >= 200 && resp.StatusCode < 400
}
