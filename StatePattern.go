// State Design Pattern

// WHAT
// WHY
// HOW

package main

import "fmt"

// state
type tvState interface {
	state()
}

//concrete implementation of state
type on struct{}

func (o *on) state() { // Implementing the Behaviour of ON state
	fmt.Println("TV is ON!")
}

type off struct{}

func (o *off) state() { // Implementing the Behaviour of OFF state
	fmt.Println("TV is OFF!")
}

// Context
type stateContext struct {
	currentTvState tvState
}

func getContext() *stateContext {
	return &stateContext{
		currentTvState: &off{},
	}
}

func (sc *stateContext) setState(state tvState) {
	sc.currentTvState = state
}

func (sc *stateContext) getState() {
	sc.currentTvState.state()
}

//Client
func main() {
	tvContext := getContext() // Default state is OFF
	tvContext.getState()      // Get the state as OFF
	tvContext.setState(&on{}) // Change the current state to ON
	tvContext.getState()      // Get the state as ON

}
