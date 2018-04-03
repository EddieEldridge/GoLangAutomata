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
	for _, r := range inputString {

		// Switch statements
		switch r {

		case '.':

			// Step 1: Remove 2 NFA fragments from the top of the stacks
			// Get the top element of the stack
			fragment2 := nfaStack[len(nfaStack)-1]

			// Remove the top element off the stack
			nfaStack = nfaStack[:len(nfaStack)-1]

			// Get the top element of the stack
			fragment1 := nfaStack[len(nfaStack)-1]

			// Remove the top element off the stack
			nfaStack = nfaStack[:len(nfaStack)-1]

			// Step 2: Join the accept state of the first fragment to the initial state of the second fragment
			fragment1.accept.edge1 = fragment2.initial

			// Step 3: Add a new fragment to the nfaStrack, from step 2.
			nfaStack = append(nfaStack, &nfaFragment{initial: fragment1.initial, accept: fragment2.accept})

		case '|':

			// Step 1: Remove 2 NFA fragments from the top of the stacks
			// Get the top element of the stack
			fragment2 := nfaStack[len(nfaStack)-1]

			// Remove the top element off the stack
			nfaStack = nfaStack[:len(nfaStack)-1]

			// Get the top element of the stack
			fragment1 := nfaStack[len(nfaStack)-1]

			// Remove the top element off the stack
			nfaStack = nfaStack[:len(nfaStack)-1]

			// Accept states
			accept := state{}
			initial := state{edge1: fragment1.initial, edge2: fragment2.initial}

			fragment1.accept.edge1 = &accept
			fragment2.accept.edge1 = &accept

			// Step 2: Add a new fragment to the nfaStrack, from step 2.
			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})

		case '*':

			// Step 1: Remove 2 NFA fragments from the top of the stacks
			// Get the top element of the stack
			frag := nfaStack[len(nfaStack)-1]

			// Remove the top element off the stack
			nfaStack = nfaStack[:len(nfaStack)-1]

			// States
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			// Step 2: Add a new fragment to the nfaStrack, from step 2.
			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})

		// When the character read in is not one of the above characters in our switch statement
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			// Push to the stack
			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})
		}
	}

	// Return the bottom item off the stack
	return nfaStack[0]
}

// Helper function for our postMatch function
func addState(stateList []*state, s *state, a *state) []*state {

	// Append our list of states
	stateList = append(stateList, s)

	if s != a && s.symbol == 0 {

		// Add to our stateList
		stateList = addState(stateList, s.edge1, a)

		if s.edge2 != nil {

			// Add to our stateList
			stateList = addState(stateList, s.edge2, a)
		}
	}

	return stateList
}

func postMatch(postFixRegExp string, s string) bool {

	// Variable to hold the status of our match
	matchingStatus := false

	// Use postRegxpToNFA function on our regular expression
	postNFA := postRegxpToNFA(postFixRegExp)

	// Pointers to our array of states
	currentState := []*state{}
	nextState := []*state{}

	// addState function
	currentState = addState(currentState[:], postNFA.initial, postNFA.accept)

	for _, rune := range s {

		// Loop through all currentStates
		for _, c := range currentState {

			// Check they are labeled by the character s
			if c.symbol == rune {
				nextState = addState(nextState[:], c.edge1, postNFA.accept)
			}
		}

		// Replace currentState with nextState and clear nextState array
		currentState, nextState = nextState, []*state{}
	}

	// After we have our currenState array set up, loop through
	for _, c := range currentState {
		if c == postNFA.accept {
			matchingStatus = true
			break
		}
	}

	// Return result of operation
	return matchingStatus
}

// Main function
func main() {

	// Set our nfa fragment to be equal to whatever regular expression we send in to our postRegxpToNFA function
	//nfaFragment := postRegxpToNFA("ab.c*|", "cccc")

	fmt.Println(postMatch("ab.c*|", ""))
}
