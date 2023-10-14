package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func UpdateAge(s *Student) {

	s.Age += 1

}

func main() {

	student := Student{

		Name: "sahil",
		Age:  21,
	}

	fmt.Println("Before updating Age")
	fmt.Println(student)

	UpdateAge(&student)

	fmt.Println("After updating Age")
	fmt.Println(student)

}
