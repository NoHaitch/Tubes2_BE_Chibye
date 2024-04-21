package main

import (
	"log"
	"time"
)

func ids(source string, target string, limit int, maxDepth int, currPath []string, pageVisitedPtr *int) ([]string, bool) {
	// log.Printf("current Page: %s\n", source)
	if source == target {
		return currPath, true
	}

	if limit <= 0 {
		return nil, false
	}

	(*pageVisitedPtr)++
	wikiTitles, err := scrapeWikipediaLinksGocolly(source)
	if err != nil {
		log.Fatal(err)
	}

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
		resultPath, found := ids(wikiTitles[i], target, limit-1, maxDepth, newPath, pageVisitedPtr)
		if found {
			return resultPath, true
		}
	}

	return nil, false
}

func idsStart(source string, target string, maxDepth int, startTime time.Time) ([]string, bool, int, int) {
	pageVisited := 0
	for limit := 0; limit <= maxDepth; limit++ {
		path, found := ids(source, target, limit, maxDepth, []string{source}, &pageVisited)
		if found {
			return path, true, int(time.Since(startTime).Milliseconds()), pageVisited
		}
	}
	return nil, false, int(time.Since(startTime).Milliseconds()), pageVisited
}
