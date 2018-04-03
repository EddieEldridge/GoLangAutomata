package main

import (
	"fmt"

	".."
)

// Main function
func main() {

	// Variables
	var regXPInput string
	var stringInput string
	var result bool

	// First, prompt for regular expression
	fmt.Print("Please enter a regular expression: ")
	fmt.Scan(&regXPInput)

	// Debug
	fmt.Println("RegExp: ", regXPInput)

	// Send to ShuntingYardAlgo function
	regXPInput = automaton.ShuntingYardAlgo(regXPInput)

	// Second, prompt for string to test against regualar expression
	fmt.Print("Please enter a string: ")
	fmt.Scan(&stringInput)

	// Send inputs to postMatch function
	result = automaton.PostMatch(regXPInput, stringInput)

	// Print results
	fmt.Println()
	fmt.Println("Match: ", result)
	fmt.Println()

}
