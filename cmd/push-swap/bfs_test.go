package main

import (
	"push-swap/ps"
	"testing"
)

func TestBfs(t *testing.T) {
	s := "7 6 5 3 1 2 4"
	beef, found := bfs("7 6 5 3 1 2 4", 13)
	if !found {
		t.Error("Not found")
	}
	a, _ := ps.NewStack(s)
	b, _ := ps.NewStack("")
	ps.Run(&a, &b, beef)
	_, sorted := ps.Check(a, b)
	if !sorted {
		t.Error("Not sorted")
	}
}
