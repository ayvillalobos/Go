package main

import (
	"fmt"
	"strings"
)

func main() {
	var phrasal string
	var word string
	fmt.Println("Enter a text:")

	fmt.Scan(&phrasal)
	fmt.Println("Enter a word:")
	fmt.Scan(&word)

	err := countWord(phrasal, word)
	fmt.Println(err)
	/*var refString = "maria es linda maria maria"
	var word = "maria"
	palabras := strings.Split(refString, " ")
	for _, palabra := range palabras {
		fmt.Println(palabra)
	}*/

}

func countWord(phrasal, word string) string {

	a := strings.Split(phrasal, ",")

	fmt.Println(a)
	fmt.Sprintln(a)
	for _, v := range a {
		if v == word {
			count := strings.Count(phrasal, word)
			return fmt.Sprintln("The word", word, "appears ", count, "times in the text.")
		}

	}

	return fmt.Sprintln("That word is not in the text.")

}
