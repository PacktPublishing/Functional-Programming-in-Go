package main

import (
	"fmt"
)

type actor struct {
	protagonist bool
}

func (a actor) talk() string {
	return "Actor is talking"
}

type driver struct {
	isGood bool
}

func (d driver) talk() string {
	return "Driver is talking"
}

// to be used by actor and driver
// we define 'methods sets' inside the interface
// instead of fields
type humanizer interface {
	talk() string
}

func talkAll(h []humanizer) {
	for _, v := range h {
		fmt.Println(v.talk())
	}

}

func main() {
	a := actor{true}
	d := driver{true}

	h := []humanizer{a, d}
	talkAll(h)
}
