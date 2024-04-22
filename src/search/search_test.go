package search_test

import (
	"fmt"
	"os"
	"runtime"
	"scraping/search"
	"sync"
	"testing"
)

func printGoroutines() {
	buf := make([]byte, 1<<20)
	runtime.Stack(buf, true)
	fmt.Fprintf(os.Stderr, "=== Goroutine stack traces ===\n%s\n", buf)
}

func TestResetRequestCounterDuplication(t *testing.T) {
	// Start the ResetRequestCounter goroutine
	go search.ResetRequestCounter()

	// Use a WaitGroup to wait for all spawned goroutines to finish
	var wg sync.WaitGroup
	wg.Add(5) // Number of goroutines spawned

	// Call IdsStart multiple times
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			search.IdsStart("John_Cena", "url_end", 5)
		}()
	}

	// Wait for all spawned goroutines to finish
	wg.Wait()

	// Print the stack traces of all goroutines
	printGoroutines()

	// Now check the number of goroutines running ResetRequestCounter
	// Assuming you have access to the runtime package for this check
	numGoroutines := runtime.NumGoroutine()
	expectedGoroutines := 1 // One for the main test goroutine
	if numGoroutines != expectedGoroutines {
		t.Errorf("Expected %d goroutines but found %d", expectedGoroutines, numGoroutines)
	}
}
