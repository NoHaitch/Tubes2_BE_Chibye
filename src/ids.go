package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var (
	RequestCounter int
	cacheDir       = "./cache"
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
	defer clearCache()
	pageVisited := 0
	for limit := 0; limit <= maxDepth; limit++ {
		path, found := ids(source, target, limit, maxDepth, []string{source}, &pageVisited)
		if found {
			return path, true, int(time.Since(startTime).Milliseconds()), pageVisited
		}
	}
	return nil, false, int(time.Since(startTime).Milliseconds()), pageVisited
}

/* ASUMSI pageTitle sudah benar tanpa spaci */
func scrapeWikipediaLinksGocolly(pageTitle string) ([]string, error) {
	// create source link
	url := "https://en.wikipedia.org/wiki/" + pageTitle

	// Initialize a new Colly collector
	c := colly.NewCollector(
		colly.CacheDir(cacheDir),
	)

	// List of links
	var wikiTitles []string

	// // Request Counter
	// c.OnRequest(func(r *colly.Request) {
	// 	if RequestCounter >= 200 {
	// 		// If the request counter reaches the limit, stop making additional requests
	// 		log.Println("Request limit reached. Stopping further requests.")
	// 		r.Abort()
	// 		return
	// 	}
	// 	RequestCounter++
	// })

	// On HTML element found, we extract links
	c.OnHTML("#mw-content-text a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")

		// Only get Wikipedia links, ignore .jpg links, ignore Wikipedia template
		if strings.HasPrefix(href, "/wiki/") && !strings.HasPrefix(href, "/wiki/File:") &&
			!strings.HasPrefix(href, "/wiki/Template:") && !strings.HasPrefix(href, "/wiki/Help:") &&
			!strings.HasPrefix(href, "/wiki/Special:") && !strings.HasPrefix(href, "/wiki/Template_talk:") &&
			!strings.HasPrefix(href, "/wiki/Category:") && !strings.HasPrefix(href, "/wiki/Wikipedia:") {

			// Remove prefix
			pageTitle := strings.TrimPrefix(href, "/wiki/")
			wikiTitles = append(wikiTitles, pageTitle)
		}
	})

	// Request Counter
	c.OnRequest(func(r *colly.Request) {
		RequestCounter++
	})

	// Visit the URL
	c.Visit(url)

	return wikiTitles, nil
}

func resetRequestCounter() {
	for {
		time.Sleep(time.Second)
		RequestCounter = 0
	}
}

func clearCache() error {
	err := os.RemoveAll(cacheDir)
	if err != nil {
		return err
	}
	log.Println("Cache cleared successfully")
	return nil
}
