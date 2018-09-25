package main

import (
	"fmt"
)

func main() {
	// setting up the key:value pairs
	key1 := []byte("name")
	v1 := "Forest Whitaker"

	// new hashmap
	hashMap := New()
	hashMap.Insert(key1, v1)

	// getting values by key
	mapValue, _ := hashMap.Get(key1)
	fmt.Println(mapValue)

}
