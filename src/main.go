package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"scraping/scrape"
	"scraping/search"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	idsMaxDepth = 5
)

func main() {

	// Starting API
	PrintlnYellow("[Main] Wikipedia Search API Starting...")
	port := "8080"

	//Initialize the banned link
	scrape.InitializeBannedLink()

	// gin instance
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Test Endpoint
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API ready to use!")
	})

	// IDS Endpoint
	r.GET("/ids", func(c *gin.Context) {
		source := c.Query("source")
		target := c.Query("target")

		// Check for bad query
		if source == "" || target == "" {
			PrintlnRed("[Main] Request Failed, Empty Query")
			c.JSON(http.StatusInternalServerError, "")

		} else {
			// Replace empty space with _
			source = strings.Replace(source, " ", "_", -1)
			target = strings.Replace(target, " ", "_", -1)

			StartGreen()
			log.Printf("Starting IDS start=%s target=%s", source, target)
			ResetColor()

			url_init := "/wiki/" + source
			url_end := "/wiki/" + target

			// Request Bucket Cycle
			stopCounter := make(chan struct{})
			defer close(stopCounter)
			go search.ResetRequestCounter(stopCounter)

			// IDS search
			startTime := time.Now()
			result, found, pageVisited := search.IdsStart(url_init, url_end, 5)
			result[len(result)-1] = "/wiki/" + target
			endTime := time.Since(startTime)

			if found {
				result[len(result)-1] = "/wiki/" + target
				PrintResult(result)
			} else {
				StartRed()
				fmt.Printf("[Main] Path not found with max depth of %d\n", idsMaxDepth)
				ResetColor()
			}

			// Result is the path
			// time takken is time of search in millisecond
			// page checked is the amount of page that is checked
			// hops is the amount of travel from the source to the target
			c.JSON(http.StatusOK, gin.H{
				"results":     result,
				"timeTakken":  endTime.Milliseconds(),
				"pageChecked": pageVisited,
				"hops: ":      len(result) - 1,
			})

			// Stop the request counter goroutine after IDS search is finished
			stopCounter <- struct{}{}
		}
	})

	// BFS Endpoint
	r.GET("/bfs", func(c *gin.Context) {
		source := c.Query("source")
		target := c.Query("target")

		// Check for bad query
		if source == "" || target == "" {
			PrintlnRed("[Main] Request Failed, Empty Query")
			c.JSON(http.StatusInternalServerError, "")

		} else {
			// Replace empty space with _
			source = strings.Replace(source, " ", "_", -1)
			target = strings.Replace(target, " ", "_", -1)

			PrintlogGreen("Starting BFS start=" + source + " target=" + target)

			url_init := "/wiki/" + source
			url_end := "/wiki/" + target

			// BFS Search
			startTime := time.Now()
			solutionsPtr := search.BFS(url_init, url_end)
			endTime := time.Since(startTime)

			result := solutionsPtr.GetPaths().GetNodes()
			reverseStringSlice(result)

			if len(result) != 0 {
				PrintResult(result)
			} else {
				StartRed()
				fmt.Printf("[Main] Path not found with max depth of %d\n", idsMaxDepth)
				ResetColor()
			}

			// Result is the path
			// time takken is time of search in millisecond
			// page checked is the amount of page that is checked
			// hops is the amount of travel from the source to the target
			c.JSON(http.StatusOK, gin.H{
				"results":     result,
				"timeTakken":  endTime.Milliseconds(),
				"pageChecked": solutionsPtr.Visited,
				"hops: ":      len(result) - 1,
			})
		}
	})

	PrintlnYellow("[Main] API started")
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)

	defer PrintlnYellow("[Main] API Terminated...")
}

func reverseStringSlice(slice []string) {
	length := len(slice)
	for i := 0; i < length/2; i++ {
		slice[i], slice[length-i-1] = slice[length-i-1], slice[i]
	}
}

// Print Color Functions
func StartYellow() {
	fmt.Print("\x1b[33m")
}

func StartGreen() {
	fmt.Print("\x1b[32m")
}
func StartRed() {
	fmt.Print("\x1b[31m")
}

func ResetColor() {
	fmt.Print("\x1b[0m")
}

func PrintlnYellow(text string) {
	StartYellow()
	fmt.Println(text)
	ResetColor()
}

func PrintlnRed(text string) {
	StartRed()
	fmt.Println(text)
	ResetColor()
}

func PrintlogGreen(text string) {
	StartGreen()
	log.Println(text)
	ResetColor()
}

func PrintResult(text []string) {
	StartGreen()
	fmt.Println(text)
	ResetColor()
}
