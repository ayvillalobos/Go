package main

import (
	"fmt"
	"sort"
)

func main() {

	var rows, columns int
	fmt.Println("ingrese un numero")
	fmt.Scan(&rows)
	fmt.Println("ingrese un numero")
	fmt.Scan(&columns)
	var mat = make([][]int, columns)

	for f := 0; f < columns; f++ {
		mat[f] = make([]int, rows)
		for c := 0; c < rows; c++ {
			fmt.Print("Ingrese los numeros: ")
			fmt.Scan(&mat[f][c])

		}
	}
	for _, v := range mat {
		sort.Ints(v)

		fmt.Println(v)

	}

}
