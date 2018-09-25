package main

import "fmt"

func showName(f func(string), name string) {
	f(name)
}

func printInConsole(name string) {
	fmt.Printf("The name is %s.\n", name)
}

// a function (getClicker) that returns a function (which returns an int)
func getClicker() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// 1. function as parameter
	showName(printInConsole, "Forest")

	// 2. returning a function
	click := getClicker()
	fmt.Println(click()) // 1
	fmt.Println(click()) // 2
	fmt.Println(click()) // 3

}
