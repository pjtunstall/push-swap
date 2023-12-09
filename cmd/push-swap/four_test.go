package main

import (
	"testing"

	"push-swap/ps"
)

type test struct {
	input string
	want  []string
}

func TestFour(t *testing.T) {
	tests := []test{
		// {"1 2 3 4", []string{}}, <-- Not tested because a sorted stack is
		// never passed to four.
		{"1 2 4 3", []string{"pb", "sa", "ra", "pa"}},
		{"1 3 2 4", []string{"pb", "sa", "pa"}},
		{"1 3 4 2", []string{"pb", "rra", "pa"}},
		{"1 4 2 3", []string{"pb", "ra", "pa"}},
		{"1 4 3 2", []string{"pb", "sa", "rra", "pa"}},

		{"2 1 3 4", []string{"pb", "ra", "pa", "rra"}},
		{"2 1 4 3", []string{"pb", "sa", "rra", "pa", "rra"}},
		{"2 3 1 4", []string{"pb", "sa", "ra", "pa", "rra"}},
		{"2 3 4 1", []string{"pb", "pa", "rra"}}, // Forestall this superfluity.
		// I could move Top to each position in turn and check if it's sorted,
		// then return Top to its original position.
		{"2 4 1 3", []string{"pb", "rra", "pa", "rra"}},
		{"2 4 3 1", []string{"pb", "sa", "pa", "rra"}},

		{"3 1 2 4", []string{"pb", "rra", "pa", "ra", "ra"}},
		// etc.

		/* But how is this a good test? It relies on trust or verifying
		each case oneself by hand, so it's hardly automated. On the other
		hand, automating any of the checks introduces a dependency on
		other parts of the code. Also, doing it this way restricts the
		the function to one solution. What if I improve the function or
		change it an a way that causes it to produce a different solution
		at least as good as the one expected by the test? I guess then the
		new solution could be verified manually and substituted for in the
		test. */
	}

	for _, tc := range tests {
		a, _ := ps.NewStack(tc.input)
		b, _ := ps.NewStack("")
		instructions := four(a, b)
		if len(instructions) != len(tc.want) {
			t.Errorf("On initial stack %v, four gives %v%v, want %v", tc.input, "", instructions, tc.want)
			continue
		}
		for i := range instructions {
			if instructions[i] != tc.want[i] {
				t.Errorf("On initial stack %v, four gives %v%v, want %v", tc.input, "", instructions, tc.want)
			}
		}
	}
}
