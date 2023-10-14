package main

import (
	"fmt"
)

func divide(numerator, denominator int) int {
	if denominator == 0 {
		panic("denominator cannot be zero")
	}
	return numerator / denominator
}

func safeDivide(numerator, denominator int) int {

	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(rec)
			fmt.Println("panic occured and handled gracefully")
		}
	}()

	return divide(numerator, denominator)
}

func main() {

	numerator, denominator := 10, 0

	if numerator == 0 && denominator == 0 {
		fmt.Println("both numerator and denominator shouldn't be zero")
		return
	}

	fmt.Println(safeDivide(numerator, denominator))

}
