package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	printlnYellow("[Main] Wikipedia Search API Starting...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// gin instance
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Define handlers
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/search", func(c *gin.Context) {
		source := c.Query("source")
		target := c.Query("target")

		log.Printf("Starting IDS start=%s target=%s", source, target)

		result, found, timeTakken, pageVisited := idsStart(source, target, 4, time.Now())

		if !found {
			log.Println("Search Failed")
		}

		// Result is the path
		// TimeTakken is time of search in milisecond
		c.JSON(http.StatusOK, gin.H{
			"results":     result,
			"timeTakken":  timeTakken,
			"pageVisited": pageVisited,
		})
	})

	printlnYellow("[Main] API started")
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)

	defer printlnYellow("[Main] API Terminated...")

	// pageTitleStart := "lion"
	// pageTitleEnd := "hen"
	// maxDepth := 3

	// startTime := time.Now()

	// if idsStart(pageTitleStart, pageTitleEnd, maxDepth) {
	// 	fmt.Println("FOUND")
	// } else {
	// 	fmt.Println("NOT FOUND")
	// }

	// elapsedTime := time.Since(startTime)
	// fmt.Printf("Execution Time: %f seconds\n", elapsedTime.Seconds())

}

func startYellow() {
	fmt.Print("\x1b[33m")
}

func resetColor() {
	fmt.Print("\x1b[0m")
}

func printlnYellow(text string) {
	startYellow()
	fmt.Println(text)
	resetColor()
}
