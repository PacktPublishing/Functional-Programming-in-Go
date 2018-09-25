package main

import (
	"fmt"
)

var identity string = ""
var mappend = func(a, b string) string {
	return a + b
}

func main() {
	strings := []string{"Forest", " Whitaker", ", 2018"}

	var concatenated = identity

	for _, s := range strings {
		concatenated = mappend(concatenated, s)
	}
	fmt.Println(concatenated)
}
