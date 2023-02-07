package main

import (
	"fmt"
	"unicode"
)

func validatePassword(password string) string {
	var validateUppercase bool
	var validateLowercase bool
	var numberPresent bool
	var validateSpecialChar bool
	minPassLength := 8
	var pass int

	for _, v := range password {
		switch {
		case unicode.IsNumber(v):
			numberPresent = true
			pass++
		case unicode.IsUpper(v):
			validateUppercase = true
			pass++
		case unicode.IsLower(v):
			validateLowercase = true
			pass++
		case unicode.IsPunct(v):
			validateSpecialChar = true
			pass++

		}
	}

	if !validateLowercase || !validateUppercase || !numberPresent || !validateSpecialChar {
		return fmt.Sprintf("Invalide Password")
	}

	if !(minPassLength <= pass) {
		return fmt.Sprintf("Invalide Password")
	}

	return fmt.Sprint("Great Password")
}

func main() {
	fmt.Println("enter your password")
	var password string
	fmt.Scan(&password)
	err := validatePassword(password)
	fmt.Println(err)
}
