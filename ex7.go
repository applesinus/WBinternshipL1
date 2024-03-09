package main

import (
	"context"
	"fmt"
	"sync"
)

func set_to_map(your_map map[int]int, ch chan int, ctx context.Context, mu *sync.Mutex, wg *sync.WaitGroup) {
	for {
		select {
		case data := <-ch:
			// locking the map using mutex
			mu.Lock()
			// saving data
			your_map[data] = data
			fmt.Printf("Value %v is saved\n", your_map[data])
			// unlocking map using mutex
			mu.Unlock()
		case <-ctx.Done():
			fmt.Printf("Goroutine has finished\n")
			wg.Done()
			return
		}
	}
}

func Ex7() {
	fmt.Printf("\n==========  Exercise 7:  ==========\n\n")

	// creating a map
	my_map := make(map[int]int)

	// creating a channel to send data
	ch := make(chan int)
	defer close(ch)

	// creating mutex to ensure data cncurrency will not occur
	mu := sync.Mutex{}

	// creating context to stop goroutines
	ctx, cancel := context.WithCancel(context.Background())

	// creating a waitgroup to wait for all workers to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// creating two goroutines
	for i := 0; i < 2; i++ {
		go set_to_map(my_map, ch, ctx, &mu, &wg)
	}

	// sending data to the channel
	for i := 0; i < 15; i++ {
		ch <- i
	}

	// sending a signal for finishing the goroutines
	cancel()

	// waiting for all goroutines to finish
	wg.Wait()

	fmt.Printf("\n===================================\n\n")
}
