package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"push-swap/ps"
)

func TestFour(t *testing.T) {
	tests := []string{
		"-999 0 999 1000000001",
		"1 2 4 3",
		"1 3 2 4",
		"1 3 4 2",
		"1 4 2 3",
		"-115 400 360 0",

		"2 1 3 4",
		"2 1 4 3",
		"2 3 1 4",
		"2 3 4 1",
		"2 4 1 3",
		"2 4 3 1",

		"3 1 2 4",
		"3 1 4 2",
		"3 2 1 4",
		"3 2 4 1",
		"3 4 1 2",
		"3 4 2 1",

		"4 1 2 3",
		"4 1 3 2",
		"4 2 1 3",
		"4 2 3 1",
		"4 3 1 2",
		"4 3 2 1",
	}

	for _, tc := range tests {
		expectedStr := strings.Split(tc, " ")
		expectedInt := make([]int, len(expectedStr))
		for i, v := range expectedStr {
			expectedInt[i], _ = strconv.Atoi(v)
		}
		sort.Ints(expectedInt)
		expected := fmt.Sprintf("%v", expectedInt)
		aInit, _ := ps.NewStack(tc)
		bInit, _ := ps.NewStack("")
		instructions := four(aInit, bInit)
		a, _ := ps.NewStack(tc)
		b, _ := ps.NewStack("")
		err := ps.Run(&a, &b, instructions)
		if err != nil {
			t.Errorf("on initial stack %v, `four` returned error: %v", tc, err)
		}
		arrInt := append(a.Nums[a.Top:], a.Nums[:a.Top]...)
		arrStr := make([]string, len(a.Nums))
		problem := false
		for i, v := range arrInt {
			arrStr[i] = strconv.Itoa(v)
			if i == len(arrInt)-1 {
				break
			}
			if v > arrInt[i+1] {
				problem = true
			}
		}
		str := strings.Join(arrStr, " ")
		if problem {
			t.Errorf("expected %v, got %v", expected, str)
		}
	}
}
