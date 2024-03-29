package main

import (
	"fmt"
)

func main() {
	qnique := GetUnique([]string{"cat", "dog", "cat", "tree", "dog", "cat", "dog", "dog"})
	fmt.Println(qnique)
}

func GetUnique(set []string) []string {
	uniqMap := map[string]struct{}{}
	for _, v := range set {
		uniqMap[v] = struct{}{} // only one unique key is allowed in map => copies will not pass
	}
	result := make([]string, 0, len(uniqMap)) // assingn keys to slice
	for k := range uniqMap {
		result = append(result, k)
	}
	return result
}
