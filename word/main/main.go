package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter a text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	phrasal := scanner.Text()
	fmt.Println("Enter a word ")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	word := scanner2.Text()

	err := countWord(phrasal, word)
	fmt.Println(err)

}

func countWord(phrasal, word string) string {

	a := strings.Split(phrasal, " ")

	for _, v := range a {
		if v == word {
			count := strings.Count(phrasal, word)
			return fmt.Sprintln("The word", word, "appears ", count, "times in the text.")
		}

	}

	return fmt.Sprintln("That word is not in the text.")

}
