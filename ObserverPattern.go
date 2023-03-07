// Observer Design Pattern in Go?
// 1. What ?
// 2. Why ?
// 3. How ?

package main

import (
	"fmt"
)

type Publisher interface {
	Register(subscriber Subscriber)
	NotifyAll(msg string)
}

type Subscriber interface {
	ReactToPublisherMsg(msg string)
}

// Publisher
type publisher struct {
	subscriberList []Subscriber
}

func GetNewPublisher() publisher {
	return publisher{subscriberList: make([]Subscriber, 0)}
}

func (pub *publisher) Register(subs Subscriber) {
	pub.subscriberList = append(pub.subscriberList, subs)
}

func (pub publisher) NotifyAll(msg string) {
	for _, subs := range pub.subscriberList {
		fmt.Println("Publisher notifying Subscriber with Id ", subs.(subscriber).subscriberId) // Type Assertion
		subs.ReactToPublisherMsg(msg)
	}
}

// Subscriber
type subscriber struct {
	subscriberId string
}

func GetNewSubscriber(Id string) subscriber {
	return subscriber{subscriberId: Id}
}

func (s subscriber) ReactToPublisherMsg(msg string) {
	fmt.Println("Subscriber Recieved message", msg, "for subscriber Id ", s.subscriberId)
}

func main() {
	pub := GetNewPublisher()

	subs := GetNewSubscriber("1")
	subs1 := GetNewSubscriber("2")
	pub.Register(subs)
	pub.Register(subs1)
	pub.NotifyAll("Hello notifying subscriber")
}
