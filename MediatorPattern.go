// Mediator Design Pattern

// WHAT
// WHY
// HOW

package main

import "fmt"

type flight interface {
	requestLanding()
	landed()
	permitLanding()
}

type mediator interface {
	canLand(flight) bool
	notify()
}

// Concrete implemetation of flight
type BAFlight struct {
	mediator mediator
}

func (f *BAFlight) requestLanding() {
	if f.mediator.canLand(f) {
		fmt.Println("BAFlight: Is Landing")
	} else {
		fmt.Println("BAFlight: Is Waiting to be Landed")
	}
}

func (f *BAFlight) landed() {
	fmt.Println("BAFlight: Has Landed")
	f.mediator.notify()
}

func (f *BAFlight) permitLanding() {
	fmt.Println("BAFlight: Has been permitted Landing")
}

type JAFlight struct {
	mediator mediator
}

func (f *JAFlight) requestLanding() {
	if f.mediator.canLand(f) {
		fmt.Println("JAFlight: Is Landing")
	} else {
		fmt.Println("JAFlight: Is Waiting to be Landed")
	}
}

func (f *JAFlight) landed() {
	fmt.Println("JAFlight: Has Landed")
	f.mediator.notify()
}

func (f *JAFlight) permitLanding() {
	fmt.Println("JAFlight: Has been permitted Landing")
}

//Concrete implemnetaiton of the mediator
type flightControlRoom struct {
	isRunWayFree bool
	flightQueue  []flight
}

func (f *flightControlRoom) canLand(flt flight) bool {
	if f.isRunWayFree {
		f.isRunWayFree = false
		return true
	}
	f.flightQueue = append(f.flightQueue, flt)
	return false
}

func (f *flightControlRoom) notify() {
	if !f.isRunWayFree {
		f.isRunWayFree = true
	}
	if len(f.flightQueue) > 0 {
		firstFlight := f.flightQueue[0]
		f.flightQueue = f.flightQueue[1:]
		firstFlight.permitLanding()
	}
}

func main() {
	fControlRoom := &flightControlRoom{
		isRunWayFree: true,
	}
	baFlight := &BAFlight{
		mediator: fControlRoom,
	}
	jaFlight := &JAFlight{
		mediator: fControlRoom,
	}

	baFlight.requestLanding()
	jaFlight.requestLanding()
	baFlight.landed()
}
