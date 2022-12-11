package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	name       string
	sideLength float64
}
type triangle struct {
	name   string
	height float64
	base   float64
}

func main() {
	sq := square{
		name:       "square",
		sideLength: 2,
	}
	t := triangle{
		name:   "triangle",
		height: 3,
		base:   3,
	}

	print(sq.name)
	printArea(sq)
	print(t.name)
	printArea(t)
}

func printArea(s shape) {
	a := s.getArea()
	fmt.Println(" has the area of", a)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
