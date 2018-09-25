package main

import "fmt"

func fib(n int) chan int {
	c := make(chan int)
	go func() {
		x, y := 0, 1
		for i := 0; i <= n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}()
	return c
}

func main() {
	for i := range fib(3) {
		fmt.Println(i)
	}
}
