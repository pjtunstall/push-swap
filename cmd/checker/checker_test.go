// This is not working yet. I'll try to learn more about testing
// by writing tests for smaller components of the program.

package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		args         []string
		instructions string
		expected     string
	}{
		{
			args:         []string{"push-swap", "0 9 1 8 2 7 3 6 4 5"},
			instructions: "sa\npb\nrrr\n",
			expected:     "KO\n",
		},
		{
			args:         []string{"push-swap", "0 9 1 8 2"},
			instructions: "pb\nra\npb\nra\nsa\nra\npa\npa\n",
			expected:     "OK\n",
		},
	}

	for i, tc := range testCases {
		t.Logf("Running test case %d", i) // Log which test case is running
		var outBuf bytes.Buffer
		in := strings.NewReader(tc.instructions)
		_, err := run(in, &outBuf, false, tc.args)
		if err != nil {
			t.Errorf("Test case %d failed with error: %v", i, err)
			continue
		}
		result := outBuf.String()
		if result != tc.expected {
			t.Errorf("Test case %d failed. Expected %s, got %s", i, tc.expected, result)
		}
	}
}
