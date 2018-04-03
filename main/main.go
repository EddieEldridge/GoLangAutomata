package main

import (
	"fmt"

	".."
)

// Need to trim off \n and \r from our stringInput
// (Thank you Cian Gannon for help)
func TrimString(s string) string {

	// If the length of the string is greater than 0
	if len(s) > 0 {

		// Remove the last two characters '\n' and '\r' from our string
		s = s[:len(s)-2]
	}

	return s
}

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
	fmt.Println()
	fmt.Println("RegExp: ", regXPInput)
	fmt.Println()

	// Send to ShuntingYardAlgo function
	regXPInput = automaton.ShuntingYardAlgo(regXPInput)

	// Second, prompt for string to test against regualar expression
	fmt.Print("Please enter a string: ")
	fmt.Scan(&stringInput)

	// Trim the string
	trimmedStringInput := TrimString(stringInput)

	// Send inputs to postMatch function
	result = automaton.PostMatch(regXPInput, trimmedStringInput)

	// Print results
	fmt.Println()
	fmt.Println("Match: ", result)
	fmt.Println()

}
