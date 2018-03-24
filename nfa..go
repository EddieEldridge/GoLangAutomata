package main

import (
	"fmt"
)

// Struct to store state
type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

// Struct to store fragment of NFA
type nfaFragment struct {
	initial *state
	accept  *state
}

// Function to change our postfix regular expression to an NFA(Non-finite automaton)
// Return a pointer to an NFA struct
func postRegxpToNFA(inputString string) *nfaFragment {

	// Create nfa stack
	nfaStack := []*nfaFragment{}

	// Loop through the regular expression one rune at a time
	for _, r := inputString{		
		
		// Switch statements
		switch r {

			case '.':
				// Get the top element of the stack
				fragment2 := nfaStack[len(nfaStack)-1]

				// Remove the top element off the stack
				nfaStack = nfaStack[:len(nfaStack)-1]

				// Get the top element of the stack
				fragment1 := nfaStack[len(nfaStack)-1]

				// Remove the top element off the stack
				nfaStack = nfaStack[:len(nfaStack)-1]
				
				// Join the two fragments
				fragment1.accept.edge1 = fragment2.initial
		
			case '|':

			case '*':

			// When the character read in is not one of the above characters in our switch statement
			default:
		}
	}

	// Return the bottom item off the stack
	return nfaStack[0]
}

// Main function
func main() {

	// Set our nfa fragment to be equal to whatever regular expression we send in to our postRegxpToNFA function
	nfaFragment := postRegxpToNFA("ab.c*|")

	fmt.Println(nfaFragment)
}
