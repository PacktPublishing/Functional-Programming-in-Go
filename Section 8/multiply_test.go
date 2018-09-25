package main

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	total := Multiply(5, 5)
	if total != 25 {
		t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestMultiplyTable(t *testing.T) {
	tables := []struct {
		x int
		y int
		m int
	}{
		{1, 1, 1},
		{1, 2, 2},
		{2, 2, 4},
		{5, 2, 10},
	}

	for _, table := range tables {
		total := Multiply(table.x, table.y)
		if total != table.m {
			t.Errorf("%d*%d != %d, but is %d", table.x, table.y, table.m, total)
		}
	}
}
