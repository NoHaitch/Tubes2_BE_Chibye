package solution

import (
	"scraping/path"
	"sync"
)

// Struct type for storing the solutions
type Solutions struct {
	sync.RWMutex
	paths   *path.Path
	Visited int
	found   bool
}

// Initialize the Solution
func (s *Solutions) Init() *Solutions {
	s.paths = path.New()
	s.found = false
	return s
}

// Create the new Solutions and return it's pointer
func New() *Solutions {
	return new(Solutions).Init()
}

// Return the condition if the solution is found
func (s *Solutions) IsFound() bool {
	s.RLock()
	defer s.RUnlock()
	return s.found
}

// Set the path as the solutions and found as true
func (s *Solutions) SetPaths(p *path.Path) {
	s.Lock()
	defer s.Unlock()
	s.paths = p
	s.found = true
}

// Return the length of the solution path
func (s *Solutions) Len() int {
	s.RLock()
	defer s.RUnlock()
	return s.paths.Len()
}

// Return the path
func (s *Solutions) GetPaths() *path.Path {
	s.RLock()
	defer s.RUnlock()
	return s.paths
}
