package main

import (
	"fmt"
	"sync"
)

func main() {
	/*  For testing purposes
	for i := 0; i < 1000; i++ {
		x := lesson3()
		if x != 220 {
			fmt.Println("Wrong answer")
		}
	}
	*/
	fmt.Println(lesson3())
}

func lesson3() int {

	var (
		array = [5]int{2, 4, 6, 8, 10}
		sum   int
		sumMu sync.Mutex // for locking to prevent data race
		wg    sync.WaitGroup
	)

	for _, elem := range array {
		wg.Add(1)
		go func(x int) {
			sumMu.Lock() // without lock order of calculations & output is not guaranteed
			sum += x * x
			sumMu.Unlock()
			wg.Done()
		}(elem)
	}

	wg.Wait()
	//fmt.Println(sum)
	return sum
}
