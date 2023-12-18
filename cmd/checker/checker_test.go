package main

import (
	"testing"

	"push-swap/ps"
)

// See also cmd/push-swap/main_test.go for a more complete test of the
// whole push-swap program against the checker program.

type testCase struct {
	instructions []string
	a            ps.Stack
	b            ps.Stack
	expected     string
}

// Simulating:

// checker % echo -e "sa\nrra\npb\n" | go run . "3 2 1 0"
// checker % echo -e "sa\nrra\npb\n" | ./checker "3 2 1 0"

// echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5"
// echo -e "sa\npb\nrrr\n" | go run . "0 9 1 8 2 7 3 6 4 5"

// echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"
// echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | go run . "0 9 1 8 2"

func TestRunInstructions(t *testing.T) {
	var result string

	testCases := []testCase{
		{
			instructions: []string{"sa", "rra", "pb"},
			a:            ps.Stack{Top: 0, Nums: []int{3, 2, 1, 0}},
			expected:     "KO",
		},
		{
			instructions: []string{"sa", "pb", "rrr"},
			a:            ps.Stack{Top: 0, Nums: []int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5}},
			expected:     "KO",
		},
		{
			instructions: []string{"pb", "ra", "pb", "ra", "sa", "ra", "pa", "pa"},
			a:            ps.Stack{Top: 0, Nums: []int{0, 9, 1, 8, 2}},
			expected:     "OK",
		},
	}

	for i, tc := range testCases {
		t.Logf("Running test case %d", i)
		err := ps.Run(&tc.a, &tc.b, tc.instructions)
		if err != nil {
			t.Errorf("Test case %d failed. Expected no error, got %s", i, err)
		}
		_, sorted := ps.Check(tc.a, tc.b)
		if sorted {
			result = "OK"
		} else {
			result = "KO"
		}
		if result != tc.expected {
			t.Errorf("Test case %d failed. Expected %s, got %s", i, tc.expected, result)
		}
	}
}
