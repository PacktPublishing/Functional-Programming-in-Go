package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to channel c
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int)
	go sum(numbers[:len(numbers)/2], c) // goroutine 1
	go sum(numbers[len(numbers)/2:], c) // goroutine 2
	x, y := <-c, <-c                    // receive x, y from channel c

	fmt.Println(x, y, x+y)

}
