package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// creating a struct with counter and internal mutex
type counter struct {
	count int64
}

// creating increment function for the struct
func (c *counter) increment() {
	// incrementing the counter with atomic to ensure data race will not occur
	atomic.AddInt64(&c.count, 1)

	// Less efficient way to do it:
	// First try I used a mutex
	/*
		c.mutex.Lock()
		c.count++
		c.mutex.Unlock()
	*/
}

// creating getCount method for the struct
func (c *counter) getCount() int64 {
	// First try I used a mutex again
	/*
		c.mutex.Lock()
		defer c.mutex.Unlock()
		return c.count
	*/
	return atomic.LoadInt64(&c.count)
}

func Ex18() {
	fmt.Printf("\n==========  Exercise 18:  ==========\n\n")

	// creating a counter
	var c counter

	// creating some goroutines
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}

	// waiting for goroutines to finish
	wg.Wait()

	// printing the count
	fmt.Printf("Count: %d\n", c.getCount())

	fmt.Printf("\n====================================\n\n")
}
