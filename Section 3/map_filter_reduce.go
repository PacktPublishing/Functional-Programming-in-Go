package main

import "fmt"

// returns the index of the number in the slice if found
// otherwise returns -1
func index(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// returns true if number is in the slice
func include(vs []int, t int) bool {
	return index(vs, t) >= 0
}

// filter returns a new slice
// that satisfies function f()
func Filter(vs []int, f func(int) bool) []int {
	vsf := make([]int, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// map returns a new slice after applying
// function f into each element of the slice
func Map(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var weights = []int{60, 80, 64, 101, 92} // in kg

	fmt.Printf("the index of number 40 is: %d. \n", index(weights, 40))
	fmt.Printf("number 80 exists in the slice: %v. \n", include(weights, 80))

	// convert kg weights into pound
	fmt.Printf("in pounds -> %v. \n", Map(weights, func(w int) int {
		return w * 2
	}))

	// find numbers bigger than 80
	fmt.Printf("bigger than 80 -> %v. \n", Filter(weights, func(w int) bool {
		return w > 80
	}))
}
