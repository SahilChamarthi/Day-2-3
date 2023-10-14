package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {

	return 3.14 * (c.Radius * c.Radius)

}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {

	return r.Height * r.Width

}

func main() {

	circle := Circle{
		Radius: 30,
	}

	rectangle := Rectangle{
		Height: 20,
		Width:  10,
	}

	fmt.Println("Area of Circle is", circle.Area())
	fmt.Println("Area of Rectangle is", rectangle.Area())

}
