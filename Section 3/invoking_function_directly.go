package main

import (
	"fmt"
)

func main() {
	age := func() int {
		return 10
	}() // declare and invoke at the same time

	fmt.Printf("age old scope: %d \n", age)

	// age := 20 // error: no new variables on left side of :=

	// creating their own scope
	func() {
		age := 20
		fmt.Printf("age new scope: %d \n", age)

		// creating a struct inside an anonymous function
		type actor struct{ name string }
		a := actor{"Forest Whitaker"}
		fmt.Println(a)
	}()

	// garbage collection friendly
	func() {
		numbers := make([]int, 0)
		for i := 1; i <= 1000; i++ {
			numbers = append(numbers, i)
		}
		// fmt.Println(numbers)
	}() // function is finished executing

}
