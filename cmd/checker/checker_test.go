// This is not working yet. I'll try to learn more about testing
// by writing tests for smaller components of the program.

package main

import (
	"testing"

	"push-swap/ps"
)

type testCase struct {
	instructions []string
	a            ps.Stack
	b            ps.Stack
	expected     string
}

// Simulating:

// echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5"
// echo -e "sa\npb\nrrr\n" | go run . "0 9 1 8 2 7 3 6 4 5"

// echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"
// echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | go run . "0 9 1 8 2"

func TestRunInstructions(t *testing.T) {
	var resultString string

	testCases := []testCase{
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
		err := run(&tc.a, &tc.b, tc.instructions)
		if err != nil {
			t.Errorf("Test case %d failed. Expected no error, got %s", i, err)
		}
		result := ps.Check(tc.a, tc.b)
		if result {
			resultString = "OK"
		} else {
			resultString = "KO"
		}
		if resultString != tc.expected {
			t.Errorf("Test case %d failed. Expected %s, got %s", i, tc.expected, resultString)
		}
	}
}
