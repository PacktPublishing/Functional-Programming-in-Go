package main

import (
	"fmt"
	"math"
)

func inPowerN(b float64) func(float64) float64 {
	return func(e float64) float64 { return math.Pow(b, e) }
}

type MethodE int

func (m MethodE) Method(b int) int {
	return int(m) + b
}

func main() {
	// currying math.Pow function
	twoInPowerOf := inPowerN(2)
	twoInPowerOfThree := twoInPowerOf(3)
	fmt.Println("2^3 =", twoInPowerOfThree)

	var m MethodE = 5

	// currying using method value
	curr := m.Method

	// uncurrying using method expressions
	uncurr := MethodE.Method

	fmt.Println("5 + 1 =", m.Method(1)) // standard

	fmt.Println("5 + 2 =", curr(2))

	fmt.Println("5 + 3 =", uncurr(m, 3))
	fmt.Println("5 + 4 =", uncurr(MethodE(5), 4))
}
