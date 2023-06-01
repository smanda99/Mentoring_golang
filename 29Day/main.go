package main

import "fmt"

type Shape interface {
	Area() float64
	Peremeter() float64
}

// Circle

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14

}

func (c Circle) Peremeter() float64 {
	return c.Radius

}

// Rectangle
type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return 0.5 * r.Length * r.Width

}

func (c Rectangle) Peremeter() float64 {
	return 2*c.Length + 2*c.Width
}

func main() {
	c := Circle{Radius: 5.9}
	r := Rectangle{Length: 15, Width: 32}

	// fmt.Println(c.Area())
	// fmt.Println(r.Area())

	shape := []Shape{c, r}

	for _, d := range shape {
		fmt.Println(d.Area())
		fmt.Println(d.Peremeter())
	}

	// shape := []Shape{c,r}

}
