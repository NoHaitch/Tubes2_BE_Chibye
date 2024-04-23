package queue

import (
	// "container/list"
	"scraping/path"
	"sync"
)

type Vertice struct {
	Url  string
	Path *path.Path
}

func (v *Vertice) GetURL() string {
	return v.Url
}

func (v *Vertice) GetPath() *path.Path {
	return v.Path
}

func (v *Vertice) GetLevel() int {
	return v.Path.Len()
}

type ListVertice struct {
	sync.RWMutex
	list   []Vertice
	length int
}

func (l *ListVertice) Init() *ListVertice {
	l.list = []Vertice{}
	l.length = 0
	return l
}

func NewListVertice() *ListVertice {
	return new(ListVertice).Init()
}

func (l *ListVertice) Add(node Vertice) {
	l.Lock()
	defer l.Unlock()
	l.list = append(l.list, node)
	l.length++
}

func (l *ListVertice) GetListVertice() []Vertice {
	l.RLock()
	defer l.RUnlock()
	return l.list
}

func (l *ListVertice) Len() int {
	l.RLock()
	defer l.RUnlock()
	return l.length
}

// type Queue struct {
// 	sync.Mutex
// 	list      *list.List
// 	currLevel int
// }

// func (q *Queue) Init() *Queue {
// 	q.list = list.New()
// 	q.currLevel = 0
// 	return q
// }

// func New() *Queue {
// 	return new(Queue).Init()
// }

// func (q *Queue) IsEmpty() bool {
// 	q.Lock()
// 	defer q.Unlock()
// 	return q.list.Len() == 0
// }

// func (q *Queue) Enqueue(link Vertice) {
// 	q.Lock()
// 	defer q.Unlock()
// 	q.list.PushBack(link)
// }

// func (q *Queue) Dequeue() Vertice {
// 	q.Lock()
// 	defer q.Unlock()
// 	front := q.list.Front()
// 	if front == nil {
// 		return Vertice{}
// 	}
// 	q.list.Remove(front)
// 	return front.Value.(Vertice)
// }

// func (q *Queue) Front() Vertice {
// 	q.Lock()
// 	defer q.Unlock()
// 	return q.list.Front().Value.(Vertice)
// }

// func (q *Queue) GetCurrentLevel() int {
// 	q.Lock()
// 	defer q.Unlock()
// 	return q.currLevel
// }

// func (q *Queue) SetCurrentLevel(newLevel int) {
// 	q.Lock()
// 	defer q.Unlock()
// 	q.currLevel = newLevel
// }

// func (q *Queue) Len() int {
// 	return q.list.Len()
// }
