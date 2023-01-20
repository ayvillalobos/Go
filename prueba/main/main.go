package main

import (
	"fmt"
)

type num int

var x num

func main() {

	fmt.Println(x)
	fmt.Printf("el tipo de x es: %T\n", x)
	x = 42

	fmt.Println(x)
}
