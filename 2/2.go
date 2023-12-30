package main

import (
	"fmt"
	"sync"
)

func main() {

	var (
		arr = [5]int{2, 4, 6, 8, 10}
		wg  sync.WaitGroup
	)

	for _, elem := range arr {
		// increase goroutine counter
		wg.Add(1)
		go func(x int) {
			fmt.Println(x * x)
			// decrease counter
			wg.Done()
		}(elem)
	}
	// wait for whole group to finish
	wg.Wait()
}
