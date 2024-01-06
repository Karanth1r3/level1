package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "りんご"
	fmt.Println(Reverse(s))
	fmt.Println(reverseString1(&s))
}

// util
func getRuneSlice(s string) []rune {
	rs := []rune(s)
	return rs
}

func Reverse(s string) string { // this one is from golang lib
	reversed := make([]byte, len(s))
	i := 0

	for len(s) > 0 {
		r, size := utf8.DecodeLastRuneInString(s) // get rune
		s = s[:len(s)-size]                       // subtract last symbol from string
		i += utf8.EncodeRune(reversed[i:], r)     // adds symbol converted to byte (string is []byte under the hood)
	}

	return string(reversed)
}

func reverseString1(s *string) string {

	rs := getRuneSlice(*s)
	// loop through slice starting with first & last rune
	// i++, j-- while not in the middle of the string
	// where all runes are already swapped by that moment
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	result := string(rs)
	return result
}
