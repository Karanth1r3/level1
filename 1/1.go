package main

// parent struct
type Human struct {
	x, y int
}

func (h *Human) move() {
	h.x++
	h.y++
}

// child struct
type Action struct {
	Human // for embedding there should be type only without name, otherwise it will be a field/variable
}

func main() {
	a := Action{
		Human{25, 25},
	}
	a.move() // call method of embedded type (if there was a field in child struct, it would not be possible, only a.human.move)
}
