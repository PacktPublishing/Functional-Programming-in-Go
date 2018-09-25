package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func isMinor(age int) bool {
	if age >= 18 {
		return false
	}

	return true
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("isMinor(20): %v\n", isMinor(20))     // pure
	fmt.Printf("rand.Intn(10): %v\n", rand.Intn(10)) // impure

	square := math.Sqrt(4)                   // pure
	fmt.Printf("math.Sqrt(4): %v\n", square) // impure (fmt.Printf)
}
