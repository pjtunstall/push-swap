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

var limit = 9 // Must be sorted in less than 9 instructions.

// Note that this doesn't take account of the pre-checks in main.go that
// deal with the case where the stack is already sorted or can be simply
// into the correct order.
func TestFive(t *testing.T) {
	// // Uncomment this and related lines, and adjust limit to explore stats.
	// fails := 0
	// scores := make([]float64, 0, 120)
	// maximum := 0
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			for k := 1; k <= 5; k++ {
				for l := 1; l <= 5; l++ {
					for m := 1; m <= 5; m++ {
						if (i == j || i == k || i == l || i == m) || (j == k || j == l || j == m) || k == l || k == m || (l == m) {
							continue
						}
						input := fmt.Sprintf("%d %d %d %d %d", i, j, k, l, m)
						a, err := ps.NewStack(input)
						if err != nil {
							t.Errorf("five(%s) failed: %s", input, err)
						}
						b, _ := ps.NewStack("")
						instructions := five(&a, &b)
						if len(instructions) >= limit {
							// // Uncomment to explore stats.
							// scores = append(scores, float64(len(instructions)))
							// fails++
							// if len(instructions) > maximum {
							// 	maximum = len(instructions)
							// }
							t.Errorf("%v took %v instructions to sort\n, not strictly less than %v", input, len(instructions), limit)
						}
						a, _ = ps.NewStack(input)
						b, _ = ps.NewStack("")
						ps.Run(&a, &b, instructions)
						_, sorted := ps.Check(a, b)
						if !sorted {
							split := strings.Split(input, " ")
							in := make([]int, len(split))
							for i, v := range split {
								in[i], err = strconv.Atoi(v)
								if err != nil {
									t.Errorf("five(%s) failed: %s", input, err)
								}
							}
							sort.Ints(in)
							t.Errorf("\nfive(%s) = %s, want %v; instructions: %v", input, a.GetNumsString(), in, instructions)
						}
					}
				}
			}
		}
	}

	// // Uncomment and set limit to 0 to see the mean and standard deviation:
	// t.Errorf("fails: %v", fails)
	// t.Errorf("mean: %v", Average(scores))
	// t.Errorf("standard deviation: %v", StandardDeviation(scores))
	// t.Errorf("maximum: %v", maximum)
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
		num := r.Intn(2000) - 1000 // generates random integers between -1000 and 999
		if !contains(nums, num) {
			nums = append(nums, num)
		}
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), " "), "[]")
}

func TestFiveRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		input := randomFive()
		a, err := ps.NewStack(input)
		if err != nil {
			t.Errorf("five(%s) failed: %s", input, err)
		}
		b, _ := ps.NewStack("")
		instructions := five(&a, &b)
		a, _ = ps.NewStack(input)
		b, _ = ps.NewStack("")
		ps.Run(&a, &b, instructions)
		if len(instructions) >= limit {
			t.Errorf("%v took %v instructions to sort,\nnot strictly less than %v", input, len(instructions), limit)
		}
		_, sorted := ps.Check(a, b)
		if !sorted {
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
