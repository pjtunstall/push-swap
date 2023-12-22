package main

import "testing"

func TestLongestRun(t *testing.T) {
	a := []int{1, 2, 4, 3, 5}
	startIndex, startValue, length := longestRun(a)
	if startIndex != 4 || startValue != 5 || length != 3 {
		t.Error("Wrong result")
	}
}
