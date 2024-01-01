package main

import (
	"cmp"
	"fmt"
)

func main() {
	arr := []int{5, 6, 8, 15, 15, 25, 30}
	fmt.Println(binarySearch[int](arr, 15))
}

// return slice of indexes of found elements
func binarySearch[T cmp.Ordered](sortedArr []T, val T) []int {
	result := []int{}
	firstFound := -1
	left := 0
	right := len(sortedArr) - 1 // initial bounds

	for left <= right { // while all /= 2 sectors arent' checked
		mid := (left + right) / 2
		//	fmt.Println(left, mid, right)
		if sortedArr[mid] == val {
			firstFound = mid
			result = append(result, firstFound)
			break
		} else if sortedArr[mid] > val {
			right = mid - 1
		} else if sortedArr[mid] < val {
			left = mid + 1
		}
	}

	if firstFound == -1 { // if not found => no reason to search for another indexes
		return result
	}

	lr, rr := firstFound, firstFound
	for lr > 0 {
		lr--
		if sortedArr[lr] != val {
			break
		}
		result = append(result, lr)
	}
	for rr < len(sortedArr)-1 {
		rr++
		if sortedArr[rr] != val {
			break
		}
		result = append(result, rr)
	}
	return result
}
