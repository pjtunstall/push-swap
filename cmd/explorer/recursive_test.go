package main

import (
	"fmt"
	"push-swap/ps"
	"strings"
	"testing"
)

// recursiveTest tests the general sorting algorithm on all permutations of n
// numbers.
func recursiveTest(t *testing.T, dice []int, depth, limit, n int) {
	// If we've assigned a value to each die...
	if depth == n {
		// ...and those values are all unique (i.e., we have a valid permutation)...
		if !allUnique(dice) {
			// ...then we can skip this permutation and return early.
			return
		}
		// Convert the dice array to a string, which will be used as input to the sorting algorithm.
		input := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(dice)), " "), "[]")
		a, err := ps.NewStack(input)
		if err != nil {
			t.Errorf("general(%s) failed: %s", input, err)
		}
		b, _ := ps.NewStack("")
		instructions := general(&a, &b)
		if len(instructions) >= limit {
			t.Errorf("more than %v instructions to sort %v numbers:\n%v took %v instructions to sort\n%v", limit-1, n, input, len(instructions), instructions)
		}
		_, sorted := ps.Check(a, b)
		if !sorted {
			expected := make([]int, n)
			for i := range expected {
				expected[i] = i + 1
			}
			t.Errorf("\ngeneral(%s) = %s, want %v", input, a.GetNumsString(), expected)
		}
		return
	}
	// If we haven't assigned a value to each die yet, then for each possible value...
	for i := 1; i <= n; i++ {
		// ...assign that value to the current die...
		dice[depth] = i
		// ...and recursively test the remaining dice.
		recursiveTest(t, dice, depth+1, limit, n)
	}
}

func allUnique(dice []int) bool {
	seen := make(map[int]bool)
	for _, v := range dice {
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}

func TestGeneral(t *testing.T) {
	// 0.183s (1/5 of a second)
	n := 5
	dice := make([]int, n)
	limit := 9
	recursiveTest(t, dice, 0, limit, n)

	// 1.599s (1 and a half seconds)
	n = 6
	dice = make([]int, n)
	limit = 13
	recursiveTest(t, dice, 0, limit, n)

	// // 243.053s (4 minutes)
	// n = 7
	// dice = make([]int, n)
	// limit = 21 // May not need to be this high.
	// recursiveTest(t, dice, 0, limit, n)
}
