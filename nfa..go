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
func postRegxpToNFA(inputString string) {

}

func main() {

	// Set our nfa fragment to be equal to whatever regular expression we send in to our postRegxpToNFA function
	nfaFragment := postRegxpToNFA("ab.c*|")

	fmt.Println(nfaFragment)
}
