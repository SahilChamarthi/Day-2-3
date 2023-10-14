package main

import "fmt"

func sum(n ...int) int {

	totalSum := 0

	for _, v := range n {
		totalSum += v
	}

	return totalSum

}

func main() {

	slc1 := []int{3, 1, 4}
	slc2 := []int{3, 1, 4, 5}
	slc3 := []int{3, 1, 4, 5, 2}

	fmt.Println(slc1, sum(slc1...))
	fmt.Println(slc2, sum(slc2...))
	fmt.Println(slc3, sum(slc3...))

}
