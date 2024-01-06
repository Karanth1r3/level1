package main

import "fmt"

func main() {

	var (
		a = 1
		s = "abc"
		f = 15.1
		b = true
	)

	fmt.Println(typeof(a))
	fmt.Println(typeof(s))
	fmt.Println(typeof(f))
	fmt.Println(typeof(b))

	fmt.Println(typeofSC(a))
	fmt.Println(typeofSC(s))
	fmt.Println(typeofSC(b))
}

// with sprintf with %T spec
func typeof(v interface{}) string {
	return fmt.Sprintf("type of variable : %T", v) // %T stands for var type
}

// with switch-case
func typeofSC(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	//... etc
	default:
		return "unknown"
	}
}

// also seen option with reflect but the task is about interface typeof
