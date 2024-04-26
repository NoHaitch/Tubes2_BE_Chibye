package visit

import (
	"sync"
)

// Work just like a set.
// Using the RWMutex to avoid race condition because this is used for concurrent processes.
type Visited struct {
	sync.RWMutex
	visited map[string]bool
}

// Initialize the Visited
func New() *Visited {
	return &Visited{
		visited: make(map[string]bool),
	}
}

// Create new Visited and return it's pointer
func (vis *Visited) Add(url string) {
	vis.Lock()
	defer vis.Unlock()
	if _, exists := vis.visited[url]; !exists {
		vis.visited[url] = false
	}
}

// Set the value of the visited map refered of it' key
func (vis *Visited) SetVisited(url string, cond bool) {
	vis.Lock()
	defer vis.Unlock()
	vis.visited[url] = cond
}

// Check the existence url in the Visited
func (vis *Visited) Check(url string) bool {
	vis.RLock()
	defer vis.RUnlock()
	_, exists := vis.visited[url]
	return exists
}

// Return the value of url key from visited
func (vis *Visited) IsVisited(url string) bool {
	vis.RLock()
	defer vis.RUnlock()
	return vis.visited[url]
}

// Return the number of url inside the Visited
func (vis *Visited) Len() int {
	vis.RLock()
	defer vis.RUnlock()
	return len(vis.visited)
}
