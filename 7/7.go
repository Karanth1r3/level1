package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	mp := make(map[int]string)
	mp2 := make(map[string]int) // initialized through make, not only declared
	// loop for concurrent write
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			mu.Lock() // maps don't support concurrent write (will panic without lock)
			writeToMap[int, string](mp, v, strconv.Itoa(v))
			writeToMap[string, int](mp2, strconv.Itoa(v), v)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(mp)
}

// generics for large range of types, no need to write separate function for each case
func writeToMap[T1 comparable, T2 any](m map[T1]T2, k T1, v T2) {
	m[k] = v
}
