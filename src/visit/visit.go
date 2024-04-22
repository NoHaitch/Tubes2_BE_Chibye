package visit

import (
	"sync"
)

type Visited struct {
	sync.RWMutex
	visited map[string]bool
}

func New() *Visited {
	return &Visited{
		visited: make(map[string]bool),
	}
}

func (vis *Visited) Add(url string) {
	vis.Lock()
	defer vis.Unlock()
	if _, exists := vis.visited[url]; !exists {
		vis.visited[url] = false
	}
}

func (vis *Visited) SetVisited(url string, cond bool) {
	vis.Lock()
	defer vis.Unlock()
	vis.visited[url] = cond
}

func (vis *Visited) Check(url string) bool {
	vis.RLock()
	defer vis.RUnlock()
	_, exists := vis.visited[url]
	return exists
}

func (vis *Visited) IsVisited(url string) bool {
	vis.RLock()
	defer vis.RUnlock()
	return vis.visited[url]
}

func (vis *Visited) Len() int {
	vis.RLock()
	defer vis.RUnlock()
	return len(vis.visited)
}
