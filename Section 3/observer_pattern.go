package main

import "fmt"

//Listener is a
type Listener struct {
	ID int
}

//ListenerInterface is an
type ListenerInterface interface {
	execute(m string)
}

func (l *Listener) execute(m string) {
	fmt.Printf("%q message receiver for id %d \n", m, l.ID)
}

//Subject is an
type Subject struct {
	listeners []ListenerInterface
}

//AddListener is a
func (s *Subject) addListener(l ListenerInterface) {
	s.listeners = append(s.listeners, l)
}

func (s *Subject) notify(m string) {
	for _, l := range s.listeners {
		if l != nil {
			l.execute(m)
		}
	}
}

var iter int

func newListener() *Listener {
	l := Listener{iter}
	iter++
	return &l
}
func main() {
	iter = 0
	s := Subject{listeners: make([]ListenerInterface, 0)}

	l := newListener()
	s.addListener(l)

	for i := 0; i < 5; i++ {
		l = newListener()
		s.addListener(l)
	}

	s.notify("Hello")
	s.notify("Goobye")
}
