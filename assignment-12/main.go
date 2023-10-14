package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct{}

func (d Dog) MakeSound() string { // Dog implementing Animal

	return "dog sounds bow bow"

}

type Cat struct{}

func (c Cat) MakeSound() string { // Cat implementing Animal

	return "cat sounds meow meow"

}

func main() {

	fmt.Println(Cat{}.MakeSound())
	fmt.Println(Dog{}.MakeSound())

}
