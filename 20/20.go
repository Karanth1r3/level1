package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "kill fuck repeat"

	fmt.Println(reverseWords(s))
}

// split then add words starting from the last word
func reverseWords(s string) string {
	parts := strings.Split(s, " ")

	var result string
	for i := len(parts) - 1; i >= 0; i-- {
		result += parts[i] + " "
	}

	return result
}
