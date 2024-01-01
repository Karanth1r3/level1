package main

import "fmt"

// Task : Implement bitwise set/clear

func main() {
	var val int64 = 16
	fmt.Println(val, "before")
	fmt.Println(setBit(&val, 1))   // setting 1st bit will result in adding 2
	fmt.Println(clearBit(&val, 4)) // 16 is 4th, clearing it will result 18 -16

}

// Sets the bit at pos in the integer n.
func setBit(n *int64, pos uint) int64 {
	*n |= (1 << pos) // OR statement
	return *n
}

// Clears the bit at pos in n.
func clearBit(n *int64, pos uint) int64 {
	mask := int64(^(1 << pos))
	*n &= mask
	return *n
}
