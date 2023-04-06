// Iterator Design Pattern

// WHAT
// WHY
// HOW

package main

import "fmt"

type User struct {
	name string
	age  int
}
type iterator interface {
	hasNext() bool
	next()
	getCurrentItem() *User
}

type iterable interface {
	getIterator() iterator
}

// Concrete implementation of iterator interface
type userIterator struct {
	index int
	users []*User
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *userIterator) next() {
	if u.hasNext() {
		u.index++
	}
}

func (u *userIterator) getCurrentItem() *User {
	if u.hasNext() {
		return u.users[u.index]
	}
	return nil
}

// Concrete implementation of iterable interface
type userIterableCollection struct {
	users []*User
}

func (u *userIterableCollection) getIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

func main() {
	usr1 := &User{
		name: "Mark",
		age:  30,
	}

	usr2 := &User{
		name: "Anderson",
		age:  35,
	}

	userIterableCollection := &userIterableCollection{
		users: []*User{usr1, usr2},
	}

	iterator := userIterableCollection.getIterator()

	for iterator.hasNext() {
		elm := iterator.getCurrentItem()
		fmt.Println("Current Item is ", elm)
		iterator.next()
	}
}
