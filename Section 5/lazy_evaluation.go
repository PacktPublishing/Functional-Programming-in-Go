package main

import (
	"fmt"
	"sync"
)

type LazyInt func() int

func Make(f func() int) LazyInt {
	var v int
	var once sync.Once

	return func() int {
		once.Do(func() {
			v = f()
			f = nil // garbage collect f
		})
		return v
	}
}

func main() {
	n := Make(func() int { return 20 }) // different computations
	fmt.Println(n())                    // calculates 20
	// reused n, the calculated value, it doesnt compute it again
	fmt.Println(n() + 40)
}
