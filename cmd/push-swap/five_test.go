package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"push-swap/ps"
)

type testCase struct {
	input string
	want  []string
}

func contains(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

func randomFive() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	nums := make([]int, 0, 5)
	for len(nums) < 5 {
		num := r.Intn(100) // generates random integers between 0 to 99
		if !contains(nums, num) {
			nums = append(nums, num)
		}
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), " "), "[]")
}

func TestFive(t *testing.T) {
	tests := []testCase{
		{"1 5 2 4 3", []string{"pb", "pb", "sa", "ra", "pa", "ra", "pa"}},
	}
	for _, tc := range tests {
		a, err := ps.NewStack(tc.input)
		if err != nil {
			t.Errorf("five(%s) failed: %s", tc.input, err)
		}
		b, _ := ps.NewStack("")
		got := five(&a, &b)
		if len(got) != len(tc.want) {
			t.Errorf("five(%s) = %s, want %s", tc.input, got, tc.want)
		} else {
			for i, v := range got {
				if v != tc.want[i] {
					t.Errorf("five(%s) = %s, want %s", tc.input, got, tc.want)
				}
			}
		}
	}

	for i := 0; i < 1000; i++ {
		input := randomFive()
		a, err := ps.NewStack(input)
		if err != nil {
			t.Errorf("five(%s) failed: %s", input, err)
		}
		b, _ := ps.NewStack("")
		got := five(&a, &b)
		if len(got) > 12 {
			t.Errorf("Took more than 12 instructions to sort.")
		}
	}
}
