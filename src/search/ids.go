package search

import (
	"fmt"
	"scraping/scrape"
	"time"
)

var (
	maxRequestPerSecond = 200
	tokenBucket         = make(chan struct{}, maxRequestPerSecond)
)

// Start IDS Search
func IDS(url_start string, url_end string, maxDepth int) ([]string, bool, int) {
	pageVisited := 0

	path := []string{}

	// IDS recursive for every depth
	for limit := 0; limit <= maxDepth; limit++ {
		fmt.Println("Depth:", limit)
		DLS(url_start, url_end, limit, []string{url_start}, &pageVisited, &path)

		// if a solution is found, stop checking
		if len(path) != 0 {
			return path, true, pageVisited
		}

		// reset page count for each limit
		pageVisited = 0
	}

	return nil, false, pageVisited
}

// DLS Recursion with Worker Pool
func DLS(source string, target string, limit int, currPath []string, visitCounter *int, path *[]string) {

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
		time.Sleep(time.Millisecond * 100)
		DLS(source, target, limit, currPath, visitCounter, path)
		<-tokenBucket // Release token
		return
	}

	(*visitCounter)++

	// Scrape
	wikiTitles := scrape.ExtractPageIDS(source, 1)

	// Create a channel to receive results from workers
	results := make(chan struct{})
	defer close(results)

	// Number of workers
	numWorkers := 20 + (limit * 3)

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
				DLS(title, target, limit-1, append(currPath, title), visitCounter, path)
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
			// Debugging
			// fmt.Println("counter: ", len(tokenBucket))

			for len(tokenBucket) > 0 {
				<-tokenBucket
			}

			time.Sleep(time.Second)
		}
	}
}
