package main

import (
	"fmt"
	"os"
	"strings"

	".."
)

// Main function
func main() {

	// Variables
	var regXPInput string
	var stringInput string
	var result bool
	var menuKeeper bool = false

	fmt.Println()
	fmt.Println("===============================")
	fmt.Println("ENTER REGXP AS 'EXIT' TO EXIT")
	fmt.Println("===============================")
	fmt.Println()

	// Loop so we don't have to restart the program every time
	for menuKeeper != true {

		// First, prompt for regular expression
		fmt.Print("Please enter a regular expression: ")
		fmt.Scan(&regXPInput)

		// If our input is EXIT, quit the application
		if regXPInput == "EXIT" {
			menuKeeper = true
		} else if menuKeeper == false {

			// Debug (Infix)
			fmt.Println()
			fmt.Println("Infix: ", regXPInput)

			// Send to ShuntingYardAlgo function
			regXPInput = automaton.ShuntingYardAlgo(regXPInput)

			// Debug (Infix)
			fmt.Println("Postfix: ", regXPInput)
			fmt.Println()

			// Second, prompt for string to test against regualar expression
			fmt.Print("Please enter a string: ")
			fmt.Scan(&stringInput)

			// Trim the string
			trimmedStringInput := strings.TrimSuffix(stringInput, "\r\n")

			// Send inputs to postMatch function
			result = automaton.PostMatch(regXPInput, trimmedStringInput)
			// Print results
			fmt.Println()
			fmt.Println("Match: ", result)
			fmt.Println()
		}

	}

	// Exit program
	os.Exit(1)

}
