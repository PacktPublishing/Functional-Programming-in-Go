package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/go-functional/core/functor"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	firstNumber := getInt()
	secondNumber := getInt()
	log.Printf("created numbers %s and %s", firstNumber, secondNumber)

	intMapperFunc := func(i int) int {
		return i + 10
	}
	mappedOptionalInt1 := firstNumber.Map(intMapperFunc)
	mappedOptionalInt2 := secondNumber.Map(intMapperFunc)

	log.Printf("mapped optional ints %s and %s", mappedOptionalInt1, mappedOptionalInt2)

}

func getInt() functor.OptionalIntFunctor {
	if rand.Int()%2 == 0 {
		return functor.SomeInt(rand.Int())
	}
	return functor.EmptyInt()
}
