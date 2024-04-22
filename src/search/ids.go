package search

import (
	"log"
	"scraping/scrape"
	"scraping/visit"
	"sync"
	"time"
)

var (
	max_request_per_second = 180
	RequestCounter         int
	stopRequestCounter     bool
	counterMutex           sync.Mutex
)

func IdsStart(url_start string, url_end string, maxDepth int) ([]string, bool, int) {
	defer scrape.ClearCache()

	// runtime.GOMAXPROCS(4)

	visitedPage := visit.New()
	pageVisited := 0

	for limit := 0; limit <= maxDepth; limit++ {
		log.Println("Depth: ", limit)
		path, found := Ids(url_start, url_end, limit, []string{url_start}, &pageVisited, visitedPage)
		if found {
			return path, true, pageVisited
		}
	}
	return nil, false, pageVisited
}

func Ids(source string, target string, limit int, currPath []string, visitCounter *int, visitedPage *visit.Visited) ([]string, bool) {
	if source == target {
		return currPath, true
	}

	if limit <= 0 {
		return nil, false
	}

	(*visitCounter)++

	// visited
	if !visitedPage.IsVisited(source) {
		IncrementRequestCounter()
		visitedPage.SetVisited(source, true)
	}

	wikiTitles := scrape.ExtractPageIDS(source)

	// search for target
	for i := range wikiTitles {

		// Stop link looping
		seen := false
		for _, link := range currPath {
			if wikiTitles[i] == link {
				seen = true
				break
			}
		}

		if seen {
			continue
		}

		newPath := append(currPath, wikiTitles[i])
		resultPath, found := Ids(wikiTitles[i], target, limit-1, newPath, visitCounter, visitedPage)
		if found {
			return resultPath, true
		}
	}

	return nil, false
}

func ResetRequestCounter() {
	for {
		counterMutex.Lock()
		RequestCounter = 0
		counterMutex.Unlock()

		time.Sleep(time.Second)
	}
}

func IncrementRequestCounter() {
	counterMutex.Lock()
	defer counterMutex.Unlock()
	RequestCounter++
}
