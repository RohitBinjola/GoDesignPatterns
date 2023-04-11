// Momento Design Pattern

// WHAT
// WHY
// HOW

package main

import "fmt"

type distance struct {
	value int
}

func newDistance() *distance {
	return &distance{
		value: 100,
	}
}

func (d *distance) travelled(dis int) {
	d.value = d.value - dis
}

func (d *distance) print() {
	fmt.Println("value of Distance is : ", d.value)
}

// momento implementation
type distanceMomento struct {
	distance distance
}

func (dm *distanceMomento) restore() distance {
	return dm.distance
}

// originator implementation
type distanceOriginator struct {
	distance distance
}

func newDistanceOriginator(distance distance) *distanceOriginator {
	return &distanceOriginator{
		distance: distance,
	}
}

func (do *distanceOriginator) saveState() *distanceMomento {
	return &distanceMomento{
		distance: do.distance,
	}
}

func (do *distanceOriginator) getState() distance {
	return do.distance
}

func (do *distanceOriginator) setState(distance distance) {
	do.distance = distance
}

func (do *distanceOriginator) restoreFromMomento(dm *distanceMomento) {
	do.distance = dm.restore()
}

// careTeker Implementation
type careTaker struct {
	momentos []distanceMomento
}

func newCareTaker() *careTaker {
	return &careTaker{
		momentos: make([]distanceMomento, 0),
	}
}

func (c *careTaker) addMomento(dm distanceMomento) {
	c.momentos = append(c.momentos, dm)
}

func (c *careTaker) getLastIndexMomento() distanceMomento {
	lastIndex := len(c.momentos) - 1
	lastMomento := c.momentos[lastIndex]
	c.momentos = c.momentos[:lastIndex]
	return lastMomento
}

func main() {
	dis := newDistance() // 100
	careTaker := newCareTaker()
	originator := newDistanceOriginator(*dis)

	dis.print() // 100

	// first Momento
	m1 := originator.saveState()
	careTaker.addMomento(*m1) // 100

	dis2 := originator.getState()
	dis2.print() // 100
	dis2.travelled(50)
	dis2.print() // 50

	// second Momento
	originator.setState(dis2)
	m2 := originator.saveState()
	careTaker.addMomento(*m2) // 50

	dis2.travelled(25)
	dis2.print() // 25

	// restoration
	lastMomento := careTaker.getLastIndexMomento()
	originator.restoreFromMomento(&lastMomento)
	dis3 := originator.getState()
	dis3.print() // 50

	lastMomento2 := careTaker.getLastIndexMomento()
	originator.restoreFromMomento(&lastMomento2)
	dis4 := originator.getState()
	dis4.print() // 100
}
