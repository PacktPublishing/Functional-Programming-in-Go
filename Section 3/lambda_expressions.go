package main

import "fmt"

func areaOfCircle(radius float64) float64 {
	area := 3.14159 * radius * radius
	return area
}

func main() {
	// emulating a lambda expression using an anonymous function
	area := func(x float64) float64 { return areaOfCircle(x) }

	fmt.Println(area(10))

	// adding two numbers
	add := func(x, y float64) float64 { return x + y }
	fmt.Println(add(3, 4))
}
