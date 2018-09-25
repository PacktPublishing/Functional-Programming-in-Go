package main

import (
	"fmt"
)

func increase() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	increment := increase()

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3
}
