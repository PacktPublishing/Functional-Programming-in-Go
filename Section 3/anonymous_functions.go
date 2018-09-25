package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// anonymous function
	anon := func() {
		fmt.Println("I am an anonymous function")
	}
	anon()

	// self invoked anonymous function
	func() {
		fmt.Println("I am a self-invoked anonymous function")
	}()

	// multiple anonymous functions
	x := 10
	functions := []func(){
		func() { x += 5 },
		func() { x -= 5 },
		func() { x /= 2 },
		func() { x *= 5 },
	}
	getRandomFunc(functions)() // change value of x
	log.Println(x)

}

func getRandomFunc(functions []func()) func() {
	return functions[rand.Intn(len(functions))]
}
