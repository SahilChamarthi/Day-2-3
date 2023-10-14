package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode int
}

type Person struct {
	Name          string
	PersonAddress Address
}

func main() {

	person := Person{

		Name:          "sahil",
		PersonAddress: Address{"sr nagar", "kadapa", 516101},
	}

	//fmt.Println(person)
	fmt.Println("Person Name:", person.Name)
	fmt.Println("Person Address:", person.PersonAddress)

}
