package main

import (
	"fmt"
	"time"
)

type actor struct {
	protagonist bool
	person      person // embedded type
}

type person struct {
	firstname string
	lastname  string
	age       int
	dob       time.Time
}

func main() {
	p := person{"Dominic", "Toretto", 42, time.Date(1976, time.August, 29, 0, 0, 0, 0, time.UTC)}
	a := actor{true, p}

	// fmt.Println(a)
	fmt.Printf("Firstname is %s. \n", a.person.firstname)
	fmt.Printf("Date of birth is %s. \n", a.person.dob.Format("2006-01-02"))

	// anonymous struct
	fmt.Println(struct {
		age int
	}{42})
}
