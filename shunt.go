package main

import "fmt"

func Intopost(infix string) string {

	// Map to that we can map characters (e.g *(),.|) to integers
	// We choose these numbers due to the fact that certain characters have precedence over others
	characterMap := map[rune]int{'*': 10, '.': 9, '|': 8}

	// Empty array of Runes (https://godoc.org/golang.org/x/text/runes)
	runeArray := []rune{}

	// Empty array of runes to be used as the stack in our program
	stack := []rune{}

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
