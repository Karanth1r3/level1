package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
}

func (c *counter) inc() {
	c.count++
}

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	counter := counter{0}
	//defer fmt.Println(counter.count, "defer") not sure how this one behaves...
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			mu.Lock() // sync is required to prevent race
			counter.inc()
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}
