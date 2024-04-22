package solution

import (
	"scraping/path"
	"sync"
)

type Solutions struct {
	sync.RWMutex
	paths   *path.Path
	Visited int
	found   bool
}

func (s *Solutions) Init() *Solutions {
	s.paths = path.New()
	s.found = false
	return s
}

func New() *Solutions {
	return new(Solutions).Init()
}

func (s *Solutions) IsFound() bool {
	s.RLock()
	defer s.RUnlock()
	return s.found
}

func (s *Solutions) SetPaths(p *path.Path) {
	s.Lock()
	defer s.Unlock()
	s.paths = p
	s.found = true
}

func (s *Solutions) Len() int {
	s.RLock()
	defer s.RUnlock()
	return s.paths.Len()
}

func (s *Solutions) GetPaths() *path.Path {
	s.RLock()
	defer s.RUnlock()
	return s.paths
}
