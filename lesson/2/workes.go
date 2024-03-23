package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Number of concurrent workers
	concurrency := 5

	// Channel to send integers to workers
	in := make(chan int)

	// Channel to signal that all workers are done
	done := make(chan []byte)

	// Goroutine to generate integers and send them to the 'in' channel
	go func() {
		i := 0
		for {
			in <- i
			i++
		}
	}()

	// Launching multiple worker goroutines
	for x := 0; x < concurrency; x++ {
		go ProcessWorker(in, x)
	}

	// Waiting for the done signal
	<-done
}

// Function to simulate a worker processing integers from the 'in' channel
func ProcessWorker(in chan int, worker int) {
	// Loop to continuously process integers received from the 'in' channel
	for x := range in {
		// Simulating some processing time
		t := time.Duration(rand.Intn(4) * int(time.Second))
		time.Sleep(t)
		// Printing the worker ID and the integer being processed
		fmt.Println("worker", worker, ": ", x)
	}
}
