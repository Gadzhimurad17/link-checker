package storage

import (
	"errors"
	"linkcheker/internal/domain/models"
	"sync"
	"sync/atomic"
)

type Links struct {
	links *sync.Map
}

func NewLinks(links *sync.Map) *Links {
	return &Links{
		links: links,
	}
}

func (l *Links) StoreLink(link string) {
	value, _ := l.links.LoadOrStore(link, new(atomic.Uint64))
	counter := value.(*atomic.Uint64)
	counter.Add(1)
}

func (l *Links) LoadLink(link string) (models.LinkStat, error) {
	value, ok := l.links.Load(link)
	if !ok {
		return models.LinkStat{}, errors.New("link doesn't exist")
	}

	counter := value.(*atomic.Uint64)
	return models.LinkStat{
		URL:    link,
		Amount: counter.Load(),
	}, nil
}

func (l *Links) DeleteLink(link string) {
	l.links.Delete(link)
}

func (l *Links) ListLinks() []models.LinkStat {
	links := make([]models.LinkStat, 0)

	l.links.Range(func(key, value any) bool {
		counter := value.(*atomic.Uint64)
		links = append(links, models.LinkStat{
			URL:    key.(string),
			Amount: counter.Load(),
		})
		return true
	})

	return links
}
