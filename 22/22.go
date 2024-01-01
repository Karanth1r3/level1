package main

import (
	"fmt"
	"math/big"
)

func main() {

	var a1, a2, sum big.Int

	a1.Exp(big.NewInt(3), big.NewInt(21), nil) // 3^21
	a2.Exp(big.NewInt(2), big.NewInt(22), nil) // 2^22

	fmt.Println(sum.Add(&a1, &a2))

	fmt.Println(sum.Sub(&a2, &a1))

	fmt.Println(sum.Mul(&a1, &a2))

	fmt.Println(sum.Div(&a2, &a1))

}
