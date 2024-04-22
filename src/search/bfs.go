package search

import (
	"fmt"
	"runtime"
	"scraping/path"
	"scraping/queue"
	"scraping/scrape"
	"scraping/solution"
	"scraping/visit"
	"sync"
)

// Search the nearest path to the destination link from start link using BFS algorithm.
func BFS(url_init, url_end string) *solution.Solutions {
	// initiate the queue
	layer := []queue.ListVertice{}

	runtime.GOMAXPROCS(11)

	init := url_init
	last := url_end

	// initiate the solution
	solution := solution.New()

	depth := 0
	// initiate the layer
	layer = append(layer, *queue.NewListVertice())
	layer[depth].Add(queue.Vertice{Url: init, Path: path.New()})

	maxProcs := 200

	counter := make(chan int, maxProcs-1)

	cc := 0

	var wg sync.WaitGroup

	visited := visit.New()

	visited.SetVisited(url_init, true)

	for {
		iter := layer[depth].Len()
		nodes := layer[depth].GetListVertice()
		layer = append(layer, *queue.NewListVertice())
		fmt.Println("Layer:", depth)
		for i := 0; i < iter; i++ {
			if !solution.IsFound() {
				cc++
				wg.Add(1)
				go scrape.ProcessPage(&nodes[i], &layer[depth+1], visited, solution, last, &wg, counter)
			} else {
				wg.Wait()
				close(counter)
				break
			}
			if (i+1)%maxProcs == 0 || i+1 == iter {
				wg.Wait()
				close(counter)
				counter = make(chan int, maxProcs-1)
			}
		}
		depth++
		wg.Wait()
		if solution.IsFound() {
			break
		}
	}

	wg.Wait()

	solution.Visited = cc

	return solution
}
