// Prototype Design Pattern in Go?
// 1. What ?
// 2. Why ?
// 3. How ? (Structs, Interfaces and Receiver Functions)

package main

import (
	"fmt"
)

type ShapeType int // Define Custom type here

const (
	CircleType ShapeType = 1
	SquareType ShapeType = 2
)

type Shape interface {
	GetId() ShapeType // Get the Shape Id
	PrintTypeProp()   // used for printing proerty values of the Shape
	Clone() Shape     // used for getting DeepCopy
}

// Implement Circle
type Circle struct {
	Id            ShapeType
	Radius        int
	Diameter      int
	Circumference int
}

func NewCircle(radius, diameter, circumference int) Circle {
	return Circle{CircleType, radius, diameter, circumference}
}

func (c Circle) GetId() ShapeType {
	return c.Id
}

func (c Circle) Clone() Shape { // Prototyping
	return NewCircle(c.Radius, c.Diameter, c.Circumference)
}

func (c Circle) PrintTypeProp() {
	fmt.Println("Circle Properties Radius, Diameter, Circumference:", c.Radius, c.Diameter, c.Circumference)
}

type Square struct {
	Id      ShapeType
	Length  int
	Breadth int
}

func NewSquare(Length, Breadth int) Square {
	return Square{SquareType, Length, Breadth}
}

func (s Square) GetId() ShapeType {
	return s.Id
}

func (s Square) Clone() Shape {
	return NewSquare(s.Length, s.Breadth)
}

func (s Square) PrintTypeProp() {
	fmt.Println("Square Properties Length, Breadth: ", s.Length, s.Breadth)
}

var RegistryList = make(map[int]Shape) // Prototype registry

func loadToRegistry() {
	circle := NewCircle(50, 40, 15)
	RegistryList[int(circle.GetId())] = circle // Adding circle to registry

	square := NewSquare(50, 40)
	RegistryList[int(square.GetId())] = square // Adding square to registry
}

func main() {
	loadToRegistry() // Load new objects data to Registry

	square := RegistryList[int(SquareType)]
	sq, ok := square.(Square) // Type Assertion
	if ok {
		fmt.Println("Old Properties:")
		sq.PrintTypeProp()
		newSquare := sq.Clone() // Prototype
		fmt.Println("Cloned object Properties:")
		newSquare.PrintTypeProp()
	}

	circle := RegistryList[int(CircleType)]
	cr, ok := circle.(Circle) // Type Assertion
	if ok {
		fmt.Println("Old Properties:")
		cr.PrintTypeProp()
		newCircle := cr.Clone().(Circle) // Prototype i.e. Cloning existing object
		newCircle.Radius = 35            // Changing property of the prototype without effecting the original object
		fmt.Println("Cloned object Changed Properties:")
		newCircle.PrintTypeProp()
	}
}
