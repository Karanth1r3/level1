package main

import "fmt"

func main() {
	val1, val2 := 1, 2
	s1, s2 := "k", "t"
	swap[int](&val1, &val2)
	swap[string](&s1, &s2)
	fmt.Println(val1, val2)
	fmt.Println(s1, s2)
}

// swap function should operate with pointers
// otherwise only copies will swap but original vars will not
func swap[T any](x, y *T) { // perhaps "any" is not the best one for this case, not sure yet
	*x, *y = *y, *x
}
