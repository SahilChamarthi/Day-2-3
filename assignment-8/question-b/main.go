package main

import "fmt"

func main() {

	fmt.Println(findFactorial(4))

}

func findFactorial(n int) int {

	if n == 1 {
		return 1
	}

	return n * findFactorial(n-1)

}
