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

func TestFive(t *testing.T) {
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
