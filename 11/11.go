package main

import "fmt"

func main() {
	i1 := []int{5, 3, 6, 8, 15}
	i2 := []int{6, 15, 0, -5, 3, 27}

	s1 := []string{"a", "x", "boom", "solar", "kafka"}
	s2 := []string{"b", "solar", "noname", "nit"}

	fmt.Println(getIntersection[int](i1, i2))
	fmt.Println(getIntersection[string](s1, s2))
}

// returns elements that are present in both inputs
func getIntersection[T comparable](m1, m2 []T) []T { // for == in loop T needs to be comparable
	result := []T{}
	for _, elem := range m1 {
		for _, elem2 := range m2 {
			if elem == elem2 {
				result = append(result, elem)
				break
			}
		}
	}
	return result
}
