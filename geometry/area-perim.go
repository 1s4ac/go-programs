package main

import (
	"fmt"
	"math"
)

type geom interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

type triangle struct {
	base, height, l float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

func (t triangle) perim() float64 {
	return t.base + t.height + t.l
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geom) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perimeter:", g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	t := triangle{base: 10, height: 6, l: 12}

	measure(r)
	measure(c)
	measure(t)
}
