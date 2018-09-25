package main

import (
	"fmt"
)

const untypeC = 3

const (
	pi float64 = 3.14159
)

func main() {
	// untyped
	var typedFloat float64 = untypeC
	fmt.Println(typedFloat)

	// pi = 3.15 // error: cannot assign to PI
	fmt.Println(pi)

	const firstname = "Forest"
	const lastname = "Whitaker"

	const name = firstname + " " + lastname

	const newFirstname = "Joe"
	const newName = newFirstname + " " + lastname

	fmt.Println(name, newName)
}
