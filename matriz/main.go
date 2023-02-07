package main

import (
	"fmt"
	"sort"
)

func matriz(columns, rows int) {
	mat := make([][]int, columns)
	var lista []int
	for f := 0; f < columns; f++ {
		mat[f] = make([]int, rows)
		for c := 0; c < rows; c++ {
			var num int
			fmt.Printf("Enter the numbers in posicion: ")
			fmt.Scan(&num)

			lista = append(lista, num)

		}
		sort.Ints(lista)

	}

	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			mat[i][j] = lista[(i*rows)+j]

		}
	}

	for _, v := range mat {

		fmt.Println(v)

	}
}

func main() {

	var rows, columns int
	fmt.Println("enter a number the rows")
	fmt.Scan(&rows)
	fmt.Println("enter a number the columns")
	fmt.Scan(&columns)

	matriz(columns, rows)

}
