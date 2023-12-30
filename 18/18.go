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
	defer fmt.Println(counter.count, "defer")
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter.inc()
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}
