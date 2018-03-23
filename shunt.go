package main

import "fmt"

func Intopost(inputString string) string {

	// Map to that we can map characters (e.g *(),.|) to integers
	// We choose these numbers due to the fact that certain characters have precedence over others
	characterMap := map[rune]int{'*': 10, '.': 9, '|': 8}

	// Empty array of Runes (https://godoc.org/golang.org/x/text/runes)
	runeArray := []rune{}

	// Empty array of runes to be used as the stack in our program
	stack := []rune{}

	// For loop to loop over our input (inputString string)
	// Using range in our loop on our string will cause it to be converted to an array of Runes
	for _, r := range inputString {

		// Switch statement for our Shunting algorithm
		switch {

		// Character cases
		case r == '(':
			stack = append(stack, r)

		case r == ')':

			// While the last element of our stack is not '(')
			for stack[len(stack)-1] != '(' {

				// Append the last element of the stack to our runeArray
				runeArray = append(runeArray, stack[len(stack)-1])

				// Remove the last element from the stack
				stack = stack[:len(stack)-1]
			}

		// If 'r' is not contained in our characterMap, a 0 will be returned
		// When a number greater than 0 is returned (i.e 'r' is not contained in our characterMap, append our runeArray with 'r')
		case characterMap[r] > 0:

		default:
			runeArray = append(runeArray, r)

		}

	}

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
