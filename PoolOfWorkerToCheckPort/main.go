package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	const (
		target       = "scanme.nmap.org"
		maxPort      = 1024
		numWorkers   = 100
		dialTimeout  = 1 * time.Second
	)

	ports := make(chan int, 100) // job queue
	var wg sync.WaitGroup

	// Worker function: reads ports from the channel and scans them
	worker := func() {
		for p := range ports {
			address := fmt.Sprintf("%s:%d", target, p)
			conn, err := net.DialTimeout("tcp", address, dialTimeout)
			if err == nil {
				fmt.Printf("Port %d is open\n", p)
				_ = conn.Close()
			}
			wg.Done()
		}
	}

	// Start workers
	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	// Enqueue jobs (ports 1..1024)
	for port := 1; port <= maxPort; port++ {
		wg.Add(1)
		ports <- port
	}

	// No more jobs
	close(ports)

	// Wait for all scans to finish
	wg.Wait()
}