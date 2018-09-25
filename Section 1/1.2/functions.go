package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	minNumber := min(3, 20)
	fmt.Println(minNumber)
}
