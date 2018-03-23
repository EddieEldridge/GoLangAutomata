package main

import "fmt"

func Intopost(infix string) string {

	// Empty array of Runes (https://godoc.org/golang.org/x/text/runes)
	runeArray := []rune{}

	// Empty array of runes to be used as the stack in our program
	stck := []rune{}

	// Return the array of Runes converted into a string
	return string(runeArray)
}

// Main Function
func main() {

	// Print out Infix and Postfix

	// ab.c*.
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ")

	// (a.(b|d))*
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: ")

	// a.(b|d).c*
	fmt.Println("Infix: ", "a.(b|d).c*")
	fmt.Println("Postfix: ")

	// a.(b.b)+.c
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Postfix: ")
}
