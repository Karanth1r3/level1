package main

// Task : conveyor of numbers, first writes to channel from array, second is for writing ar[..] * 2

import "fmt"

func main() {

	var (
		arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	)

	//channels are not buffered, so when the first one receives value, it's sent to the second one after that; same for every number in array, one by one
	c := readArr(arr)  // creating first channel, first element of pipeline (readArr returns readonly channel)
	out := fillSqrs(c) // second channel, second element, fillsqrs returns rdonly chnl

	//fmt.Println(<-out)
	for k := range out {
		fmt.Println(k)
	}
	//go fillSqrs(sqCh)

}

func readArr(ar [10]int) <-chan int {
	out := make(chan int)
	go func() {
		for _, elem := range ar {
			out <- elem
		}
		close(out)
	}()
	return out
}

func fillSqrs(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range inCh {
			out <- x * x
		}
		close(out)
	}()
	return out
}
