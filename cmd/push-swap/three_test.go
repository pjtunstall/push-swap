package main

import (
	"testing"
)

type ExpectedOutput struct {
	instructions []string
	rot          bool
}

type testCaseThree struct {
	input []int
	want  ExpectedOutput
}

func TestThree(t *testing.T) {
	tests := []testCaseThree{
		{[]int{1, 2, 3}, ExpectedOutput{[]string{}, true}},
		{[]int{3, 1, 2}, ExpectedOutput{[]string{"ra"}, true}},
		{[]int{2, 3, 1}, ExpectedOutput{[]string{"rra"}, true}},
		{[]int{2, 1, 3}, ExpectedOutput{[]string{"sa"}, false}},
		{[]int{1, 3, 2}, ExpectedOutput{[]string{"sa", "ra"}, false}},
		{[]int{3, 2, 1}, ExpectedOutput{[]string{"sa", "rra"}, false}},
	}

	for _, tc := range tests {
		got, rot := three(tc.input)
		want := tc.want
		if len(got) != len(want.instructions) || rot != want.rot {
			t.Errorf("three(%v) = %v, want %v", tc.input, got, want)
		}
		for i := range got {
			if got[i] != want.instructions[i] {
				t.Errorf("three(%v) = %v, want %v", tc.input, got, want)
			}
		}
	}
}
