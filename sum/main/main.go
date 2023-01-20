package main

import (
	"fmt"
)

func addtion(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) int {
	return a / b
}

func validate(a, b int) int {
	if a > b {
		return a
	}
	return b

}
func calculate(a, b int) {
	addtion := addtion(a, b)
	subtraction := subtraction(a, b)
	multiplication := multiplication(a, b)
	division := division(a, b)
	validate := validate(a, b)

	fmt.Printf("%v, %v, %v, %v, The Number Major Is: %v\n ", addtion, subtraction, multiplication, division, validate)

}

func main() {
	calculate(5, 5)

}
