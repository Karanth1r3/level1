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
	*n |= (1 << pos) // original bit + all 0 except the bit on targeted position (x | 1 = 1)
	return *n
}

// Clears the bit at pos in n.
func clearBit(n *int64, pos uint) int64 {
	mask := int64(^(1 << pos)) // create mask with 1 on unchanged bits, 0 on targeted bit
	*n &= mask                 // multiply original number by mask, targeted bit shall change it's value to x & 0, other ones won't change, cuz x & 1 = x
	return *n
}
