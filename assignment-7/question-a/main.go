package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	Age  int
	City string
}

func main() {

	emp1 := Employee{
		Id:   101,
		Name: "king",
		Age:  25,
		City: "banglore",
	}

	emp2 := Employee{
		Id:   101,
		Name: "queen",
		Age:  24,
		City: "chennai",
	}

	fmt.Println(emp1)
	fmt.Println(emp2)

}
