package main

import (
	"cmp"
	"fmt"
)

func main() {
	arr := []int{500, 0, 73, 85, 9}
	(quickSort[int](arr))
	fmt.Println(arr)
}

func quickSort[T cmp.Ordered](arr []T) {
	qSort(arr, 0, len(arr)-1)
}

func qSort[T cmp.Ordered](arr []T, left, right int) {
	if left < right { // if not sorted yet
		p := partition[T](arr, left, right) // find pivot point

		qSort[T](arr, left, p-1)  // sort left sublists to the current pivot
		qSort[T](arr, p+1, right) // sort right sublists to the current pivot
	}
}

// return pivot element index and later divide arr to the left (all values will end up being < pivot)
// & right of the pivot (all values shall be > pivot)
func partition[T cmp.Ordered](arr []T, left, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j <= right-1; j++ { // from the left side to the end
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] //
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1] //
	return i + 1
}
