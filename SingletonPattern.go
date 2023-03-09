// Singleton Design Pattern in Go?
// 1. What ?
// 2. Why ?
// 3. How ?

package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

type config struct {
	// Config variables
}

var counter int = 1
var singleConfigInstance *config

func getConfigInstance() *config {
	if singleConfigInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if singleConfigInstance == nil {
			fmt.Println("Creting Single Instance Now, and Counter is ", counter)
			singleConfigInstance = &config{}
			counter = counter + 1
		} else {
			fmt.Println("Single Instance already created-1, returning that one")
		}
	} else {
		fmt.Println("Single Instance already created-2, returning the same")
	}
	return singleConfigInstance
}

func main() {
	for i := 0; i < 100; i++ {
		go getConfigInstance()
	}
	fmt.Scanln()
}
