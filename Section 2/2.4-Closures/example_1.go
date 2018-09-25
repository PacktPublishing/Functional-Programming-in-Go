package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{20, 16, 18, 14, 10, 20, 24, 45, 50, 35, 36}
	sort.Ints(ages)
	fmt.Println("sorted ages: ", ages)

	index := sort.Search(len(ages), func(i int) bool {
		return ages[i] >= 18
	})

	fmt.Println(">= 18 years old:", ages[index:])
}
