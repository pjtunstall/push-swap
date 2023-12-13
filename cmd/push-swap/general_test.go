package main

import (
	"fmt"
	"math/rand"
	"push-swap/ps"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestGeneral(t *testing.T) {
	// Check length of result for 100 random numbers 100 times:
	for i := 0; i < 100; i++ {
		hundred := hundredRandomNumbers()
		a, err := ps.NewStack(hundred)
		if err != nil {
			t.Errorf("general(%s) failed: %s", hundred, err)
		}
		b, _ := ps.NewStack("")
		instructions := general(&a, &b)
		_, sorted := ps.Check(a, b)
		if !sorted {
			t.Errorf("not sorted")
		}
		if len(instructions) >= 700 {
			t.Errorf("%v instructions--that's a bit much", len(instructions))
		}
	}

	// Check that the result is sorted for all permutations of 1-6:
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
							general(&a, &b)
							// instructions := general(&a, &b)
							// if len(instructions) > 12 {
							// 	t.Errorf("%v took more than 12 instructions to sort", input)
							// }
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

// Generate a string of a hundred random numbers:
func hundredRandomNumbers() string {
	// The first 100 primes greater than 100:
	primes := []int{101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997}
	var result string
	var arr [100]int

	rand.Shuffle(len(primes), func(i, j int) { primes[i], primes[j] = primes[j], primes[i] })

	// Choose 100 random numbers between -100 and 100 and multiply
	// them by the first 100 primes greater than 100 to ensure
	// that they are all unique:
	for i := 0; i < 100; i++ {
		r := randInt(-99, 100) * primes[i]
		if r == 0 {
			i--
			continue
		}
		arr[i] = r
	}

	for i, v := range arr {
		result += fmt.Sprintf("%d", v)
		if i != 99 {
			result += " "
		}
	}

	return result
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
