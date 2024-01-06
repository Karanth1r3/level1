package main

import (
	"math/rand"
	"strings"
)

var substr string // is global, gc won't touch it until end of execute

func createHugeString(x int) string {
	result := []byte{}
	for i := 0; i < x; i++ {
		result = append(result, byte(rand.Intn(26)+'a'))
	}
	return string(result)
}

func part() {
	s := createHugeString(1 << 10)
	// s will be gc'ed after a smaller copy (of size [:x] instead of full) will be assigned to substr
	substr = strings.Clone(s[:100])
	//substr = s[:100]   // whole s will be allocated in memory this way (substr will contain link to substring)
}

func main() {
	part()
}
