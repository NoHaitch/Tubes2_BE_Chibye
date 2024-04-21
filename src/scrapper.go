package main

import (
	"strings"

	"github.com/gocolly/colly"
)

/* ASUMSI pageTitle sudah benar tanpa spaci */
func scrapeWikipediaLinksGocolly(pageTitle string) ([]string, error) {
	// create source link
	url := "https://en.wikipedia.org/wiki/" + pageTitle

	// Initialize a new Colly collector
	c := colly.NewCollector()

	// List of links
	var wikiTitles []string

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

	// Visit the URL
	c.Visit(url)

	return wikiTitles, nil
}
