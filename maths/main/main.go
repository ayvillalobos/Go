package main

import (
	"fmt"
)

func addition(a, b int) int {
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

func validate(a, b int) string {
	if a == b {
		return "equal numbers"
	}

	return fmt.Sprintf("The greater of %v and %v, is: %v", a, b, greater(a, b))

}

func greater(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculate(a, b int) {

	addition := addition(a, b)
	subtraction := subtraction(a, b)
	multiplication := multiplication(a, b)
	division := division(a, b)
	fmt.Printf("Addition: %v, Subtraction: %v, Multiplication: %v, Division: %v \n", addition, subtraction, multiplication, division)
	fmt.Println(validate(a, b))
}

func main() {
	calculate(9, 5)

}
