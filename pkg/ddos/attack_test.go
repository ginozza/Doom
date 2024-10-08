package attack_test

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	ddos "github.com/ginozza/doom/pkg/ddos"
)

// getFreePort attempts to find an available free port within the user port range.
func getFreePort() int {
	// Port range for the search
	for port := 1024; port <= 65535; port++ {
		address := fmt.Sprintf("localhost:%d", port)
		// Try to listen on the port
		listener, err := net.Listen("tcp", address)
		if err != nil {
			// The port is already in use, try the next one
			continue
		}
		// Close the listener after finding a free port
		listener.Close()
		return port
	}
	// If no free port is found, return 0
	return 0
}

func TestNewDDoS(t *testing.T) {
	d, err := ddos.New("http://127.0.0.1", 1)
	if err != nil {
		t.Error("Cannot create a new ddos structure. Error = ", err)
	}
	if d == nil {
		t.Error("Cannot create a new ddos structure")
	}
}

func TestDDoS(t *testing.T) {
	port := getFreePort()
	if port == 0 {
		t.Fatalf("Cannot find a free TCP port")
	}
	createServer(port, t)

	url := "http://127.0.0.1:" + strconv.Itoa(port)
	d, err := ddos.New(url, 1000)
	if err != nil {
		t.Error("Cannot create a new ddos structure")
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	success, amount := d.Result()
	if success == 0 || amount == 0 {
		t.Errorf("Negative result of DDoS attack.\n"+
			"Success requests = %v.\n"+
			"Amount requests = %v", success, amount)
	}
	t.Logf("Statistic: %d %d", success, amount)
}

// Create a simple Go server
func createServer(port int, t *testing.T) {
	// Channel to receive server errors
	errCh := make(chan error, 1)

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		})
		err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
		errCh <- err // Send the error to the channel
	}()

	// Wait for a short period to ensure the server starts
	time.Sleep(time.Millisecond * 100)

	// Check if there were any errors starting the server
	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("Server is down. %v", err)
		}
	default:
		// No errors, the server is running
	}
}

func TestWorkers(t *testing.T) {
	_, err := ddos.New("127.0.0.1", 0)
	if err == nil {
		t.Error("Cannot create a new ddos structure")
	}
}

func TestUrl(t *testing.T) {
	_, err := ddos.New("some_strange_host", 1)
	if err == nil {
		t.Error("Cannot create a new ddos structure")
	}
}

func ExampleNew() {
	workers := 1000
	d, err := ddos.New("http://127.0.0.1:80", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Fprintf(os.Stdout, "DDoS attack server: http://127.0.0.1:80\n")
	// Output:
	// DDoS attack server: http://127.0.0.1:80
}
