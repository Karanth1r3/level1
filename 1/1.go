package main

type Human struct {
	x, y int
}

func (h *Human) move() {
	h.x++
	h.y++
}

type Action struct {
	Human // for embedding there should be type only without name, otherwise it will be a field/variable
}

func main() {
	a := Action{
		Human{25, 25},
	}
	a.move() // call method of embedded type
}
