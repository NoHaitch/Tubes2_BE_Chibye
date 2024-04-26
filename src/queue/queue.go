package queue

import (
	"scraping/path"
	"sync"
)

// Struct type for Node in BFS
type Vertice struct {
	Url  string
	Path *path.Path
}

// Getter for URL
func (v *Vertice) GetURL() string {
	return v.Url
}

// Getter for Path
func (v *Vertice) GetPath() *path.Path {
	return v.Path
}

// Return the length of the path as Level
func (v *Vertice) GetLevel() int {
	return v.Path.Len()
}

// This list vertice is used for the concurrent process
// Using RWMutex to avoid the race condition
// Save the list of vertice which is processed
type ListVertice struct {
	sync.RWMutex
	list   []Vertice
	length int
}

// Initialize the ListVertice
func (l *ListVertice) Init() *ListVertice {
	l.list = []Vertice{}
	l.length = 0
	return l
}

// Initialize the new ListVertice and return it's pointer
func NewListVertice() *ListVertice {
	return new(ListVertice).Init()
}

// Add the vertice to the end of the list
func (l *ListVertice) Add(node Vertice) {
	l.Lock()
	defer l.Unlock()
	l.list = append(l.list, node)
	l.length++
}

// Getter for list of Vertice
func (l *ListVertice) GetListVertice() []Vertice {
	l.RLock()
	defer l.RUnlock()
	return l.list
}

// Return the number of Vertice stored
func (l *ListVertice) Len() int {
	l.RLock()
	defer l.RUnlock()
	return l.length
}
