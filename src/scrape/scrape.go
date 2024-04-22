package scrape

import (
	"fmt"
	"runtime"
	"scraping/path"
	"scraping/queue"
	"scraping/solution"
	"scraping/visit"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

// Visit the page and scrape all the body.
// Parsing the body and select the link that contains "/wiki/" prefix.
// Check the link in the visited map.
// If the link doesn't exist in the visited map add it into localVisited and links slice.
// Return links slice
func ExtractPage(visited *visit.Visited, localVisited *visit.Visited, url string) []string {

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.AllowURLRevisit = true

	links := []string{}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Collect links with href attribute starting with "/wiki/"
		if strings.HasPrefix(link, "/wiki/") {
			// Mark the link as visited
			if !visited.Check(link) {
				links = append(links, link)
				localVisited.Add(link)
			}
			visited.Add(link)
		}
	})

	time.Sleep(time.Millisecond)

	err := c.Visit("https://en.wikipedia.org" + url)
	if err != nil {
		fmt.Println("Error to visit URL: " + url)
		return []string{}
	}

	return links
}

// This function call the ExtractPage function.
// If the endURL (link destination) exist in the localVisited map, set it as solutions.
// If not, add all the unvisited link from links slice to the line queue and set true for the value of its map.
func ProcessPage(node *queue.Vertice, line *queue.ListVertice, visited *visit.Visited, solution *solution.Solutions, endURL string, wg *sync.WaitGroup, counter chan int) {

	localVisited := visit.New()

	links := ExtractPage(visited, localVisited, node.GetURL())
	if localVisited.Check(endURL) {
		pathLines := *node.Path.Add(node.GetURL())
		pathLines = *pathLines.Add(endURL)
		solution.SetPaths(&pathLines)
	} else {
		for _, info := range links {
			if !visited.IsVisited(info) {
				line.Add(queue.Vertice{Url: info, Path: node.Path.Add(node.GetURL())})
				visited.SetVisited(info, true)
			}
		}
	}

	defer wg.Done()
	select {
	case counter <- 0:
	default:
	}
}

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
				go ProcessPage(&nodes[i], &layer[depth+1], visited, solution, last, &wg, counter)
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
