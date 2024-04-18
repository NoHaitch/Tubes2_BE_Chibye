package main

import (
	"fmt"
	"time"
)

func main() {
	pageTitleStart := "Joko_Anwar"
	pageTitleEnd := "Kampala"
	maxDepth := 3

	startTime := time.Now()

	if idsStart(pageTitleStart, pageTitleEnd, maxDepth) {
		fmt.Println("FOUND")
	} else {
		fmt.Println("NOT FOUND")
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution Time: %f seconds\n", elapsedTime.Seconds())
}
