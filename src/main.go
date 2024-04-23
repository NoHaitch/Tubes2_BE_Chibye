package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"scraping/scrape"
	"scraping/search"

	"github.com/gin-gonic/gin"
)

func main() {
	// Clear Cache
	scrape.ClearCache()

	// Starting API
	PrintlnYellow("[Main] Wikipedia Search API Starting...")
	port := "8080"

	// gin instance
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Test Endpoint
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API ready to use!")
	})

	// IDS Endpoint
	r.GET("/ids", func(c *gin.Context) {
		source := c.Query("source")
		target := c.Query("target")

		log.Printf("Starting IDS start=%s target=%s", source, target)
		url_init := "/wiki/" + source
		url_end := "/wiki/" + target

		startTime := time.Now()

		stopCounter := make(chan struct{})
		defer close(stopCounter)

		go search.ResetRequestCounter(stopCounter)

		result, found, pageVisited := search.IdsStart(url_init, url_end, 5)
		endTime := time.Since(startTime)

		if !found {
			log.Println("Search Failed")
		}

		// Result is the path
		// TimeTakken is time of search in milisecond
		c.JSON(http.StatusOK, gin.H{
			"results":     result,
			"timeTakken":  endTime.Milliseconds(),
			"pageVisited": pageVisited,
		})

		// Stop the request counter goroutine after IDS search is finished
		stopCounter <- struct{}{}
	})

	// BFS Endpoint
	r.GET("/bfs", func(c *gin.Context) {
		source := c.Query("source")
		target := c.Query("target")

		log.Printf("Starting BFS start=%s target=%s", source, target)
		url_init := "/wiki/" + source
		url_end := "/wiki/" + target

		startTime := time.Now()
		solutionsPtr := search.BFS(url_init, url_end)
		endTime := time.Since(startTime)

		// Result is the path
		// TimeTakken is time of search in milisecond
		c.JSON(http.StatusOK, gin.H{
			"results":     solutionsPtr.GetPaths(),
			"timeTakken":  endTime.Milliseconds(),
			"pageVisited": solutionsPtr.Visited,
		})
	})

	PrintlnYellow("[Main] API started")
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)

	defer PrintlnYellow("[Main] API Terminated...")
}

func StartYellow() {
	fmt.Print("\x1b[33m")
}

func ResetColor() {
	fmt.Print("\x1b[0m")
}

func PrintlnYellow(text string) {
	StartYellow()
	fmt.Println(text)
	ResetColor()
}
