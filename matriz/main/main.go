package main

import (
	"fmt"
	"sort"
)

func main() {

	var rows, columns int
	fmt.Println("enter a number")
	fmt.Scan(&rows)
	fmt.Println("enter a number")
	fmt.Scan(&columns)
	var mat = make([][]int, columns)

	for f := 0; f < columns; f++ {
		mat[f] = make([]int, rows)
		for c := 0; c < rows; c++ {
			fmt.Print("Enter the numbers: ")
			fmt.Scan(&mat[f][c])
		}
		fmt.Println(mat)
	}

	for _, v := range mat {
		sort.Ints(v)
		fmt.Println(v)

	}

}

func ordenarMenor(listNum []int, Cant int) {
	tmp := 0
	for x := 0; x < Cant; x++ {
		for y := 0; y < Cant; y++ {
			if listNum[x] < listNum[y] {
				tmp = listNum[y]
				listNum[y] = listNum[x]
				listNum[x] = tmp
			}
		}
	}
}
