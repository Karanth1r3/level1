package main

import "fmt"

func main() {
	v1 := []int{1, 2, 3}
	//fmt.Println(v1)
	v2 := []string{"a", "b", "c"}
	v1 = removeUnordered[int](0, v1)
	v2 = remove[string](1, v2)
	fmt.Println(v1, v2)
}

func removeUnordered[T any](i int, slc []T) []T { // T any to make one function for all incoming types
	slc[i] = slc[len(slc)-1]
	return slc[:len(slc)-1]
}

func remove[T any](i int, slc []T) []T {
	slc = append(slc[:i], slc[i+1:]...)
	return slc
}
