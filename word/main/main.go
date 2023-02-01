package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var phrasal string
	//var word string
	/*
		fmt.Print("Enter a text:")

		fmt.Scan(&phrasal)
		fmt.Println("Enter a word")
		fmt.Scan(&word)*/
	fmt.Print("Enter a text:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)
	fmt.Println("Enter a word")
	//fmt.Scan(&word)
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	texto := scanner.Text()
	fmt.Println(texto)

	err := countWord(text, texto)
	fmt.Println(err)

	/*var refString = "maria es linda maria maria"
	var word = "maria"
	palabras := strings.Split(refString, " ")
	for _, palabra := range palabras {
		fmt.Println(palabra)
	}*/

}

func countWord(phrasal, word string) string {

	a := strings.Split(phrasal, " ")

	for _, v := range a {
		if v == word {
			count := strings.Count(phrasal, word)
			return fmt.Sprintln("The word", word, "appears ", count, "times in the text.")
		}
		fmt.Println(v)
		fmt.Println("posicion", len(v))
	}

	return fmt.Sprintln("That word is not in the text.")

}
