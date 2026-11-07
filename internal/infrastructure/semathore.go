package infrastructure

type Semathore struct {
	C chan struct{}
}

func (s *Semathore) Acquire(n int) {
	for range make([]struct{}, n) {
		s.C <- struct{}{}
	}
}

func (s *Semathore) Release(n int) {
	for range make([]struct{}, n) {
		<-s.C
	}
}
