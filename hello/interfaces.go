package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length, height float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length * r.height)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rectangle{length: 3, height: 5}
	c := circle{radius: 7}

	measure(r)
	measure(c)
}
