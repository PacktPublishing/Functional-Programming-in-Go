package main

import (
	"fmt"
	"functional-programming-in-go/4_immutability_and_monads/immutable_struct/movie"
)

func main() {
	m := movie.NewMovie("The Last Stand", 2013)
	fmt.Println(m.name)
	fmt.Printf(m.String())
}
