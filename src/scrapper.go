package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func scrapeWikipediaLinksGoquery(pageTitle string) ([]string, error) {
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

func scrapeWikipediaLinksGocolly(pageTitle string) ([]string, error) {
	// create source link
	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", strings.ReplaceAll(pageTitle, " ", "_"))

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

	// Error handling
	c.OnError(func(r *colly.Response, err error) {
		return
	})

	// Visit the URL
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return wikiTitles, nil
}

func testSpeed(pageTitle string) {
	goqueryStartTime := time.Now()

	_, err := scrapeWikipediaLinksGoquery(pageTitle)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	goqueryElapsedTime := time.Since(goqueryStartTime)
	fmt.Printf("Total Execution Time Goquery: %d Milliseconds\n", goqueryElapsedTime.Milliseconds())

	gocollyStartTime := time.Now()

	_, err = scrapeWikipediaLinksGocolly(pageTitle)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	gocollyElapsedTime := time.Since(gocollyStartTime)
	fmt.Printf("Total Execution Time Gocolly: %d Milliseconds\n", gocollyElapsedTime.Milliseconds())

}
