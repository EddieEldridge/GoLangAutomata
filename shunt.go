package main

import "fmt"

// Shunting-Yard Algorithm (https://en.wikipedia.org/wiki/Shunting-yard_algorithm)
func ShuntingYardAlgo(inputString string) string {

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

			// While the top element of our stack is not '('
			for stack[len(stack)-1] != '(' {

				// Append the top element of the stack to our runeArray
				runeArray = append(runeArray, stack[len(stack)-1])

				// Remove the top element from the stack
				stack = stack[:len(stack)-1]
			}

			// If 'r' is not contained in our characterMap, a 0 will be returned
			// When a number greater than 0 is returned (i.e 'r' is not contained in our characterMap, append our runeArray with 'r')
		case characterMap[r] > 0:

			// While the stack still has elements in AND while the precedence of
			// the current character being read is less than the precedence of the character at the top of the stack
			for len(stack) > 0 && characterMap[r] <= characterMap[stack[len(stack)-1]] {

				// Append the top element of the stack to our runeArray
				runeArray = append(runeArray, stack[len(stack)-1])

				// Remove the top element from the stack
				stack = stack[:len(stack)-1]
			}
			// Append our the character being read in to the stack
			stack = append(stack, r)

		// If the string being read doesn't contain one of our special case characters, append our runeArray stack with the characters
		default:
			runeArray = append(runeArray, r)

		}

	}

	for len(stack) > 0 {

		// Append the top element of the stack to our runeArray
		runeArray = append(runeArray, stack[len(stack)-1])

		// Remove the top element from the stack
		stack = stack[:len(stack)-1]
	}

	// Return the array of Runes converted into a string
	return string(runeArray)
}

// Main Function

func printResults() {

	// Print out Infix
	// Run the algorithm on our infix and print the result as Postfix:

	// ab.c*.
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", ShuntingYardAlgo("a.b.c*"))

	// (a.(b|d))*
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: ", ShuntingYardAlgo("(a.(b|d))*"))

	// a.(b|d).c*
	fmt.Println("Infix: ", "a.(b|d).c*")
	fmt.Println("Postfix: ", ShuntingYardAlgo("a.(b|d).c*"))

	// a.(b.b)+.c
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", ShuntingYardAlgo("a.(b.b)+.c"))
}
