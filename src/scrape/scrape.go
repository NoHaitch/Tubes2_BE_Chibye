package scrape

import (
	"fmt"
	"os"
	"scraping/queue"
	"scraping/solution"
	"scraping/visit"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

var (
	BfsCacheDir = "./cache"
)

// Visit the page and scrape all the body.
// Parsing the body and select the link that contains "/wiki/" prefix.
// Check the link in the visited map.
// If the link doesn't exist in the visited map add it into localVisited and links slice.
// Return links slice
func ExtractPage(visited *visit.Visited, localVisited *visit.Visited, url string) []string {

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.AllowURLRevisit = true

	links := []string{}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Collect links with href attribute starting with "/wiki/"
		if strings.HasPrefix(link, "/wiki/") {
			// Mark the link as visited
			if !visited.Check(link) {
				links = append(links, link)
				localVisited.Add(link)
			}
			visited.Add(link)
		}
	})

	time.Sleep(time.Millisecond)

	err := c.Visit("https://en.wikipedia.org" + url)
	if err != nil {
		fmt.Println("Error to visit URL: " + url)
		return []string{}
	}

	return links
}

// This function call the ExtractPage function.
// If the endURL (link destination) exist in the localVisited map, set it as solutions.
// If not, add all the unvisited link from links slice to the line queue and set true for the value of its map.
func ProcessPage(node *queue.Vertice, line *queue.ListVertice, visited *visit.Visited, solution *solution.Solutions, endURL string, wg *sync.WaitGroup, counter chan int) {

	localVisited := visit.New()

	links := ExtractPage(visited, localVisited, node.GetURL())
	if localVisited.Check(endURL) {
		pathLines := *node.Path.Add(node.GetURL())
		pathLines = *pathLines.Add(endURL)
		solution.SetPaths(&pathLines)
	} else {
		for _, info := range links {
			if !visited.IsVisited(info) {
				line.Add(queue.Vertice{Url: info, Path: node.Path.Add(node.GetURL())})
				visited.SetVisited(info, true)
			}
		}
	}

	defer wg.Done()
	select {
	case counter <- 0:
	default:
	}
}

// Visit the page and scrape all the body.
// Parsing the body and select the link that contains "/wiki/" prefix and non important articles.
// Cache the page into BfsCacheDir, which will be deleted when the ids search finished
func ExtractPageIDS(url string, try int) []string {

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.CacheDir(BfsCacheDir),
	)

	c.AllowURLRevisit = true

	// List of links
	links := []string{}

	// Get href links
	c.OnHTML("#mw-content-text a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")

		// Only get Wikipedia links, ignore .jpg links, ignore Wikipedia template
		if strings.HasPrefix(href, "/wiki/") && !strings.HasPrefix(href, "/wiki/File:") &&
			!strings.HasPrefix(href, "/wiki/Template:") && !strings.HasPrefix(href, "/wiki/Help:") &&
			!strings.HasPrefix(href, "/wiki/Special:") && !strings.HasPrefix(href, "/wiki/Template_talk:") &&
			!strings.HasPrefix(href, "/wiki/Category:") && !strings.HasPrefix(href, "/wiki/Wikipedia:") {

			links = append(links, href)
		}
	})

	// Visit the URL
	err := c.Visit("https://en.wikipedia.org" + url)
	if err != nil {
		fmt.Println("Error to visit URL: " + url)
		if try > 1 {
			return nil
		} else {
			time.Sleep(time.Millisecond)
			return ExtractPageIDS(url, try+1)
		}
	}

	return links
}

// Remove Cache used by IDS search
func ClearCache() {
	os.RemoveAll(BfsCacheDir)
}
