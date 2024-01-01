package main

import "fmt"

func main() {

	var (
		arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	)

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
