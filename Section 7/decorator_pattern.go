package main

import (
	"fmt"
	"strings"
)

type stringModifier func(string) string

func toLowercase(sm stringModifier) stringModifier {
	return func(s string) string {
		lower := strings.ToLower(s)
		return sm(lower)
	}
}

func main() {
	name := "Forest Whitaker"

	var fn stringModifier = func(s string) string {
		return s
	}
	fmt.Println(fn(name)) // default value

	// using the decorator
	fn = toLowercase(fn)
	fmt.Println(fn(name)) // output default values again ?

}
