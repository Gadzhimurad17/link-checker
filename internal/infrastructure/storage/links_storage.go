package storage

import (
	"errors"
	"sync"
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
	l.links.Store(link, true)
}

func (l *Links) LoadLink(link string) (any, error) {
	value, ok := l.links.Load(link)
	if !ok {
		return nil, errors.New("link doesnt exist")
	}
	return value, nil
}

func (l *Links) DeleteLink(link string) {
	l.links.Delete(link)
}
