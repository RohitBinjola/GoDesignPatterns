// Builder Design Pattern

// What ?
// Why ?
// How ?
// Pros and Cons ?

package main

import "fmt"

// type houseBuilder struct {
// 	window string
// 	door   string
// 	wall   string
// 	floor  int
// }

// func getNewHouseBuilder() houseBuilder {
// 	return houseBuilder{}
// }

// func (b *houseBuilder) buildWall() *houseBuilder {
// 	b.wall = "Wooden Wall"
// 	return b
// }

// func (b *houseBuilder) buildWindow() *houseBuilder {
// 	b.window = "Wooden Window"
// 	return b
// }

// func (b *houseBuilder) buildDoor() *houseBuilder {
// 	b.door = "Wooden Door"
// 	return b
// }

// func (b *houseBuilder) buildNumFloor() *houseBuilder {
// 	b.floor = 2
// 	return b
// }

// func main() {
// 	houseBuilder := getNewHouseBuilder()
// 	houseBuilder.buildDoor().buildNumFloor().buildWall().buildWindow()
// 	fmt.Println(houseBuilder)
// }

//Product
type house struct {
	wall   string
	window string
	door   string
	floor  int
}

// Builder
type houseBuilder struct {
	window string
	door   string
	wall   string
	floor  int
}

func gethouseBuilder() houseBuilder {
	return houseBuilder{}
}

func (b *houseBuilder) buildWall() {
	b.wall = "Wooden Wall"
}

func (b *houseBuilder) buildWindow() {
	b.window = "Wooden Window"
}

func (b *houseBuilder) buildDoor() {
	b.door = "Wooden Door"
}

func (b *houseBuilder) buildNumFloor() {
	b.floor = 2
}

func (b *houseBuilder) buildHouse() house {
	return house{
		wall:   b.wall,
		door:   b.door,
		window: b.window,
		floor:  b.floor,
	}
}

// Director
type director struct {
	builder houseBuilder
}

func newDirector(b houseBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) buildHouse() house {
	d.builder.buildWall()
	d.builder.buildDoor()
	d.builder.buildWindow()
	d.builder.buildNumFloor()
	return d.builder.buildHouse()
}

func main() {
	houseBuilder := gethouseBuilder()

	director := newDirector(houseBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Wall is : %s\n", normalHouse.wall)
	fmt.Printf("Normal House Door is : %s\n", normalHouse.door)
	fmt.Printf("Normal House Window is : %s\n", normalHouse.window)
	fmt.Printf("Normal House Num Floor is : %d\n", normalHouse.floor)
}
