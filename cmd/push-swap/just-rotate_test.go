package main

import (
	"push-swap/ps"
	"testing"
)

func TestJustRotate(t *testing.T) {
	inputs := []string{
		"2 3 4 5 1",
		"3 4 5 1 2",
		"4 5 1 2 3",
		"5 1 2 3 4",
		"1 2 3 4 5",
	}
	for _, input := range inputs {
		a, _ := ps.NewStack(input)
		b, _ := ps.NewStack("")
		instructions := justRotate(a)
		ps.Run(&a, &b, instructions)
		output := a.GetNumsString()
		if output != "1 2 3 4 5" {
			t.Errorf("\nExpected 1 2 3 4 5, got %s", output)
			t.Errorf("\nInstructions: %s", instructions)
		}
	}
}
