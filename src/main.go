package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	pageTitleStart := "Joko_Anwar"
	pageTitleEnd := "Kampala"

	startTime := time.Now()

	if idsStart(pageTitleStart, pageTitleEnd, 3) {
		fmt.Println("FOUND")
	} else {
		fmt.Println("NOT FOUND")
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution Time: %f seconds\n", elapsedTime.Seconds())
}

func scrapeWikipediaLinks(pageTitle string) ([]string, error) {
	// create source link
	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", strings.ReplaceAll(pageTitle, " ", "_"))

	// Fetch Page
	resp, err := http.Get(url)
	if err != nil {
		// Fetching Error
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch page: %s", resp.Status)
	}

	// Parse HTTP response to goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// List of links
	var wikiTitles []string

	// selects all <a> element
	doc.Find("#mw-content-text a").Each(func(i int, s *goquery.Selection) {

		// if <a> element have href link
		href, exists := s.Attr("href")
		if exists {

			// Only get wikipedia links, igonore .jpg links, igonore wikipedia template
			if strings.HasPrefix(href, "/wiki/") && !strings.HasPrefix(href, "/wiki/File:") &&
				!strings.HasPrefix(href, "/wiki/Template:") && !strings.HasPrefix(href, "/wiki/Help:") &&
				!strings.HasPrefix(href, "/wiki/Special:") && !strings.HasPrefix(href, "/wiki/Template_talk:") &&
				!strings.HasPrefix(href, "/wiki/Category:") && !strings.HasPrefix(href, "/wiki/Wikipedia:") {

				// Remove prefix
				pageTitle := strings.TrimPrefix(href, "/wiki/")
				wikiTitles = append(wikiTitles, pageTitle)
			}
		}
	})

	return wikiTitles, nil
}

func ids(srcTitle string, targetTitle string, limit int, currPath []string) ([]string, bool) {
	fmt.Printf("> %d %s\n", limit, srcTitle)
	if srcTitle == targetTitle {
		return currPath, true
	}

	if limit <= 0 {
		return nil, false
	}

	wikiTitles, err := scrapeWikipediaLinks(srcTitle)
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

	fmt.Println(" == Finish a ids")

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
