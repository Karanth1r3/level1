package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func Distance(p1, p2 Point) float64 {
	var dx, dy float64
	//
	dx = p1.x - p2.x // suitable because of power of 2 in return
	dy = p1.y - p2.y

	return math.Sqrt(dx*dx + dy*dy)
}

func main() {

	p1 := NewPoint(5, 5)
	p2 := NewPoint(0, 0)

	fmt.Println(Distance(p1, p2))
}
