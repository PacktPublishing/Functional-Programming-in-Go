package main

import (
	"fmt"
	"time"
)

type person struct {
	firstname string
	lastname  string
	age       int
	added     time.Time
}

type Actor struct {
	isProtagonist bool
	person        person // embedding
}

func (a *Actor) String() string {
	return fmt.Sprintf("%s %s is %d years old.", a.person.firstname, a.person.lastname, a.person.age)
}

func (a Actor) walk() string {
	return fmt.Sprintf("Actor: %s %s is walking", a.person.firstname, a.person.lastname)
}

type Director struct {
	firstname     string
	lastname      string
	hasNobelPrice bool
}

func (d Director) walk() string {
	return fmt.Sprintf("Director: %s %s is walking", d.firstname, d.lastname)
}

type walker interface {
	walk() string
}

// emulating a constructor using New()
func New(fn, ln string, age int, isProtagonist bool) *Actor {
	return &Actor{
		isProtagonist: true,
		person: person{
			firstname: fn,
			lastname:  ln,
			age:       age,
			added:     time.Now(),
		},
	}
}

func main() {
	// this is supposed to be a package
	// so it is called actor.New()
	actor := New("Dom", "Toretto", 42, true) // new instance of Actor
	fmt.Println(actor.String())

	director := Director{"John", "Doe", false}

	// using interfaces
	walker := []walker{actor, director}

	for _, v := range walker {
		fmt.Println(v.walk())
	}

}
