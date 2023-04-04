// Command Design Pattern

// WHAT
// WHY
// HOW

package main

import "fmt"

//Command interface
type command interface {
	execute()
}

type device interface { // fired on a device
	on()
	off()
}

//Concrete class/object implementing command interface
type onCommand struct { // To switch on the device
	device device
}

func (r *onCommand) setDevice(device device) {
	r.device = device
}

func (loc onCommand) execute() {
	loc.device.on()
}

type offCommand struct { // To switch off the device
	device device
}

func (r *offCommand) setDevice(device device) {
	r.device = device
}

func (loc *offCommand) execute() {
	loc.device.off()
}

// Concrete implementation of device interface
// Request Object/Receiver i.e. device where commands will work
type lightBulb struct {
}

func (lb lightBulb) on() {
	fmt.Println("Light bulb ON!")
}

func (lb lightBulb) off() {
	fmt.Println("Light bulb OFF!")
}

//Invoker
type remote struct {
	cmd command
}

func (r *remote) setCommand(cmd command) {
	r.cmd = cmd
}

func (r *remote) pressButton() {
	r.cmd.execute()
}

func main() {
	lightBulb := &lightBulb{}

	onCommand := &onCommand{}
	onCommand.setDevice(lightBulb)

	offCommand := &offCommand{}
	offCommand.setDevice(lightBulb)

	oncmd := &remote{}
	oncmd.setCommand(onCommand)
	oncmd.pressButton()

	offcmd := &remote{}
	offcmd.setCommand(offCommand)
	offcmd.pressButton()

}
