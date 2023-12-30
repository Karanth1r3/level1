package main

import "fmt"

type (
	printer struct{}
	adapter struct{ printer }
	output  interface {
		print(int)
	}
)

// not compatible with printer
func writeOutput(out output) {
	out.print(8)
}

// translate int to string (initial format to compatible)
func (a *adapter) print(val int) {
	a.printer.print(fmt.Sprint(val))
}

func (p *printer) print(s string) {
	fmt.Println(s)
}

func main() {
	var pp = printer{}
	var aa = adapter{pp}

	writeOutput(&aa)
}
