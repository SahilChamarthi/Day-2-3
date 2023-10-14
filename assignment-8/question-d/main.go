package main

import (
	"assignment-8/question-c/model"
	"assignment-8/question-d/shape"
)

func main() {

	circle := model.Circle{

		Radius: 30,
	}

	shape.AreaOfCircle(circle.Radius)

}
