package main

import (
	"fmt"
)

// read only variables
const (
	eulerConst = 2.71828 // eulers constant
)

// file block variable declaration
var (
	isProtagonist = true
	isFast        = true
)

func main() {
	name := "Dominic" // short variable declaration
	var lastname string = "Toretto"
	var age = 42

	var friend string // declare
	friend = "Brian"  // and initialize
	fmt.Println(name, lastname, age, friend, eulerConst, isProtagonist)
}
