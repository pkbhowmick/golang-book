package main

import (
	"fmt"
	"math"
)

// ***** Structs and Methods *******
type Circle struct {
	x, y, r float64
}
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w
}

func (c *Circle) perimeter() float64 {
	return 2.0 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2.0 * (l + w)
}

// ****** Interface ********
type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64

	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var totalPerimeter float64

	for _, s := range shapes {
		totalPerimeter += s.perimeter()
	}
	return totalPerimeter
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	fmt.Println(c.perimeter())

	r := Rectangle{0, 0, 5, 3}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())

	fmt.Println(totalArea(&c, &r))
	fmt.Println(totalPerimeter(&c, &r))
}
