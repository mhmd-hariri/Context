package main

import (
	"context"
	"fmt"
	"time"
)

// Simulate a function that does some work and respects the context for cancellation.
func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("work canceled:", ctx.Err())
			return
		default:
			// Simulate work
			fmt.Println("working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// Create a context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())

	// Start a goroutine that does some work
	go doWork(ctx)

	// Let the work run for 2 seconds
	time.Sleep(2 * time.Second)

	// Cancel the context to stop the goroutine
	cancel()

	// Give some time for the goroutine to print the cancellation message
	time.Sleep(1 * time.Second)
}
