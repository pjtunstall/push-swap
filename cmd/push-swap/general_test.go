package main

import (
	"fmt"
	"push-swap/ps"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestSix(t *testing.T) {
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			for k := 1; k <= 6; k++ {
				for l := 1; l <= 6; l++ {
					for m := 1; m <= 6; m++ {
						for n := 1; n <= 6; n++ {
							if (i == j || i == k || i == l || i == m || i == n) || (j == k || j == l || j == m || j == n) || k == l || k == m || k == n || (l == m || l == n) || m == n {
								continue
							}
							input := fmt.Sprintf("%d %d %d %d %d %d", i, j, k, l, m, n)
							a, err := ps.NewStack(input)
							if err != nil {
								t.Errorf("general(%s) failed: %s", input, err)
							}
							b, _ := ps.NewStack("")
							instructions := general(&a, &b)
							if len(instructions) > 12 {
								t.Errorf("%v took more than 12 instructions to sort", input)
							}
							_, sorted := ps.Check(a, b)
							if !sorted {
								split := strings.Split(input, " ")
								in := make([]int, len(split))
								for i, v := range split {
									in[i], err = strconv.Atoi(v)
									if err != nil {
										t.Errorf("general(%s) failed: %s", input, err)
									}
								}
								sort.Ints(in)
								t.Errorf("\ngeneral(%s) = %s, want %v", input, a.GetNumsString(), in)
							}
						}
					}
				}
			}
		}
	}
}
