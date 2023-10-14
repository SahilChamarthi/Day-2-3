package main

import (
	"assignment-8/question-c/model"
	"assignment-8/question-c/person"
)

func main() {

	p1 := model.Person{

		Name: "sahil",
		Age:  22,
	}

	person.PrintPerson(p1)

}
