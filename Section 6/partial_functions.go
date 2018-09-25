package main

import "fmt"

func addSome(x, y, z int) int {
	return x + y + z
}

func addBy(x int) func(int, int) int {
	return func(y, z int) int {
		return addSome(x, y, z)
	}
}

func main() {
	addBy10 := addBy(10)
	fmt.Println(addBy10(4, 6)) // 20
}
