package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) AreaOfRectangle() {

	fmt.Println("Area of Rectangle is", r.Width*r.Height)

}

func (r Rectangle) PerimeterOfRectangle() {

	fmt.Println("Perimeter of Rectangle is", (r.Height+r.Width)*2)

}

func main() {

	rec := Rectangle{
		Width:  10,
		Height: 20,
	}

	rec.AreaOfRectangle()
	rec.PerimeterOfRectangle()

}
