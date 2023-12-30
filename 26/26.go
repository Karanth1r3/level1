package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "aaab"
	s2 := "rofl"
	fmt.Printf("First string is Unique? : %t \n", chckUnique(s1))
	fmt.Printf("Second string is  Unique? : %t ", chckUnique(s2))
}

// sorting requires runes
func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

// sort based on rune symbol index in table
func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

// when string is sorted and converted to lowercase =>
// if 2 elems

func chckUnique(s string) bool {
	unique := true
	s = strings.ToLower(s)
	r := StringToRuneSlice(s)

	for i := 0; i < len(r)-1; i++ {

		fmt.Println(r[i], r[i+1])
		if r[i] == r[i+1] {
			unique = false
		}
	}
	return unique
}
