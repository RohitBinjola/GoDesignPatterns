// Chain of Responsibility Design Pattern

// WHAT
// WHY
// HOW

package main

import "fmt"

type step interface {
	execute()     // Func to be executed next.
	setNext(step) // Func to set the next func to be executed.
}

type openFlap struct {
	next step
}

type pressStartButton struct {
	next step
}

type enterPassword struct {
	next step
}

// Open Flap

func (r *openFlap) execute() {
	fmt.Println("Opening Laptop Flap")
	r.next.execute()
}

func (r *openFlap) setNext(next step) {
	r.next = next
}

// Press Button Below

func (d *pressStartButton) execute() {
	fmt.Println("Pressing start button")
	d.next.execute()
}

func (d *pressStartButton) setNext(next step) {
	d.next = next
}

// Enter Password

func (m *enterPassword) execute() {
	fmt.Println("Entering password for Laptop")
	fmt.Println("Laptop Opened!")
}

func (m *enterPassword) setNext(next step) {
	m.next = next
}

func main() {
	//Set next for enterPassword step
	enterPassword := &enterPassword{}

	//Set next for pressStartButton step
	pressStartButton := &pressStartButton{}
	pressStartButton.setNext(enterPassword)

	//Set next for openFlap step
	openFlap := &openFlap{}
	openFlap.setNext(pressStartButton)

	openFlap.execute()
}
