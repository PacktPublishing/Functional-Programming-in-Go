package main

import (
	"fmt"
)

func main() {
	getNumberOfUsers := func() int {
		return 20
	}

	fmt.Println(getNumberOfUsers())
}
