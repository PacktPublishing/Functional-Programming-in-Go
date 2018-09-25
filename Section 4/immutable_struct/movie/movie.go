package movie

import (
	"fmt"
)

type movie struct {
	name        string // non-exported
	releaseYear uint   // non-exported
}

// initialize
func NewMovie(name string, releaseYear uint) *movie {
	m := &movie{
		name:        name,
		releaseYear: releaseYear,
	}

	return m
}

func (m *movie) String() string {
	return fmt.Sprintf("%s was released in %d.\n", m.name, m.releaseYear)
}
