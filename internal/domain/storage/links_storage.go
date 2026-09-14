package storage

import "sync"

type Links struct {
	links sync.Map
}

func (l *Links) Save(link string) {
	l.links.Store(link, true)
}
