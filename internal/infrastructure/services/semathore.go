package services

type semathore struct {
	C chan struct{}
}

func NewSemathore(ch chan struct{}) *semathore {
	return &semathore{
		C: ch,
	}
}

func (s *semathore) Acquire(n int) {
	for range make([]struct{}, n) {
		s.C <- struct{}{}
	}
}

func (s *semathore) Release(n int) {
	for range make([]struct{}, n) {
		<-s.C
	}
}
