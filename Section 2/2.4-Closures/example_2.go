package main

import (
	"fmt"
)

func main() {
	clicker := newClick()
	fmt.Println(clicker()) // nrClicks = 1
	fmt.Println(clicker()) // nrClicks = 2
}

func newClick() func() int {
	nrClicks := 0
	return func() int {
		nrClicks++
		return nrClicks
	}
}
