package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
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
		{"3 5 2 4 1", []string{"pb", "pb", "rra", "pa", "rra", "pa", "rra", "rra"}},
		{"1 5 2 4 3", []string{"pb", "pb", "sa", "ra", "pa", "ra", "pa"}},
		{"3 2 1 4 5", []string{"pb", "pb", "rr", "pa", "pa", "rra"}},
	}
	for _, tc := range tests {
		a, err := ps.NewStack(tc.input)
		if err != nil {
			t.Errorf("five(%s) failed: %s", tc.input, err)
		}
		b, _ := ps.NewStack("")
		// fmt.Println(a.GetNumsString())
		got := five(&a, &b)

		_, sorted := ps.Check(a, b)
		if !sorted {
			fmt.Println(tc.input)
			fmt.Println(a.GetNumsString())
			fmt.Println(got)
			fmt.Println()
			fmt.Println()
			t.Errorf("\nfive(%s) = %s, want %s", tc.input, got, tc.want)
		}
	}

	for i := 0; i < 100; i++ {
		input := randomFive()
		a, err := ps.NewStack(input)
		if err != nil {
			t.Errorf("five(%s) failed: %s", input, err)
		}
		b, _ := ps.NewStack("")
		instructions := five(&a, &b)
		if len(instructions) > 12 {
			t.Errorf("%v took more than 12 instructions to sort", input)
		}
		_, sorted := ps.Check(a, b)
		if !sorted {
			fmt.Println(input)
			fmt.Println(a.GetNumsString())
			fmt.Println(instructions)
			fmt.Println()

			split := strings.Split(input, " ")
			in := make([]int, len(split))
			for i, v := range split {
				in[i], _ = strconv.Atoi(v)
			}
			sort.Ints(in)
			t.Errorf("\nfive(%s) = %s, want %v", input, a.GetNumsString(), in)
		}
	}
}
