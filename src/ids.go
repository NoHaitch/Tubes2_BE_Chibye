package main

import (
	"fmt"
	"log"
)

func ids(srcTitle string, targetTitle string, limit int, currPath []string) ([]string, bool) {
	if srcTitle == targetTitle {
		return currPath, true
	}

	if limit <= 0 {
		return nil, false
	}

	wikiTitles, err := scrapeWikipediaLinksGocolly(srcTitle)
	if err != nil {
		log.Fatal(err)
	}

	// search for target
	for i := range wikiTitles {
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
		resultPath, found := ids(wikiTitles[i], targetTitle, limit-1, newPath)
		if found {
			return resultPath, true
		}
	}

	return nil, false
}

func idsStart(srcTitle string, targetTitle string, maxDepth int) bool {
	for limit := 0; limit <= maxDepth; limit++ {
		path, found := ids(srcTitle, targetTitle, limit, nil)
		if found {
			fmt.Println("Wikipedia Links:")
			fmt.Printf(" 1. %s\n", srcTitle)
			for i, title := range path {
				fmt.Printf(" %d. %s \n", i+2, title)
			}
			return true
		}
	}
	return false
}
