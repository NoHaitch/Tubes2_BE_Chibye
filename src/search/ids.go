package search

import (
	"fmt"
	"log"
	"scraping/scrape"
	"scraping/visit"
	"time"
)

var (
	tokenBucket = make(chan struct{}, 1000)
)

// Start IDS Search
func IdsStart(url_start string, url_end string, maxDepth int) ([]string, bool, int) {
	cachePage := visit.New()
	pageVisited := 0

	path := []string{}

	// IDS recursive for every depth
	for limit := 0; limit <= maxDepth; limit++ {
		log.Println("Depth:", limit)
		Ids(url_start, url_end, limit, []string{url_start}, &pageVisited, cachePage, &path)

		// if a solution is found, stop checking
		if len(path) != 0 {
			return path, true, pageVisited
		}

		// reset page count for each limit
		pageVisited = 0
	}

	return nil, false, pageVisited
}

// IDS Recursion with Worker Pool
func Ids(source string, target string, limit int, currPath []string, visitCounter *int, cachePage *visit.Visited, path *[]string) {
	// A solution is already found
	if len(*path) != 0 {
		return
	}

	// target found
	if source == target {
		*path = currPath
		return
	}

	// limit reached
	if limit <= 0 {
		return
	}

	// Check if Request Bucket is full
	// Currently will never happen, last time used was bugged
	select {
	case tokenBucket <- struct{}{}:
	default:
		time.Sleep(time.Millisecond)
		Ids(source, target, limit, currPath, visitCounter, cachePage, path)
		<-tokenBucket // Release token
		return
	}

	// add url to cache map
	if !cachePage.IsVisited(source) {
		cachePage.SetVisited(source, true)
	}

	(*visitCounter)++

	// Scrape
	wikiTitles := scrape.ExtractPageIDS(source, 1)

	// Create a channel to receive results from workers
	results := make(chan struct{})
	defer close(results)

	// Number of workers
	numWorkers := 15 + limit

	// Calculate workload per worker
	workload := len(wikiTitles) / numWorkers

	// Start workers
	for i := 0; i < numWorkers; i++ {
		start := i * workload
		end := (i + 1) * workload
		if i == numWorkers-1 {
			end = len(wikiTitles)
		}

		go func(titles []string) {
			defer func() { results <- struct{}{} }()
			for _, title := range titles {
				// ids recursion
				Ids(title, target, limit-1, append(currPath, title), visitCounter, cachePage, path)
			}
		}(wikiTitles[start:end])
	}

	// Wait for all workers to finish
	for i := 0; i < numWorkers; i++ {
		<-results
	}
}

func ResetRequestCounter(stop chan struct{}) {
	for {
		select {
		case <-stop:
			return
		default:
			fmt.Println("counter: ", len(tokenBucket))
			for len(tokenBucket) > 0 {
				<-tokenBucket
			}

			time.Sleep(time.Second)
		}
	}
}
