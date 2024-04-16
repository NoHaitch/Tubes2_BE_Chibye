package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	pageTitle := "Joko_Widodo"
	wikiLinks, err := scrapeWikipediaLinks(pageTitle)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wikipedia Links:")
	for _, link := range wikiLinks {
		fmt.Println(link)
	}
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
	var wikiLinks []string

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
				wikiLinks = append(wikiLinks, "https://en.wikipedia.org"+href)
			}
		}
	})

	return wikiLinks, nil
}
