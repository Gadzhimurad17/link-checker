package services

type semathore struct {
	ch chan struct{}
}

func NewSemathore(ch chan struct{}) *semathore {
	return &semathore{
		ch: ch,
	}
}

func (s *semathore) Acquire(n int) {
	for range make([]struct{}, n) {
		s.ch <- struct{}{}
	}
}

func (s *semathore) Release(n int) {
	for range make([]struct{}, n) {
		<-s.ch
	}
}
