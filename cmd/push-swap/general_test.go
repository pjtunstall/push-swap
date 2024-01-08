package main

import (
	"fmt"
	"math"
	"math/rand"
	"push-swap/ps"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// 10000 times tests that 100 random numbers are sorted in less than 700 instructions,
// and 6! times tests that all permutations of 1-6 are sorted in less than 9 instructions:
func TestGeneral(t *testing.T) {
	limit := 700 // Must be under this.

	// // Uncomment this and related lines, and adjust limit to explore stats.
	// fails := 0
	// scores := make([]float64, 0, 10000)

	for i := 0; i < 10000; i++ {
		hundred := hundredRandomNumbers()
		a, err := ps.NewStack(hundred)
		if err != nil {
			t.Errorf("general(%s) failed: %s", hundred, err)
		}
		b, _ := ps.NewStack("")
		instructions := general(&a, &b)
		a, _ = ps.NewStack(hundred)
		a, _ = ps.NewStack(rank(a.Nums))
		original := a.GetNumsSlice()
		b, _ = ps.NewStack("")
		ps.Run(&a, &b, instructions)
		_, sorted := ps.Check(a, b)
		if !sorted {
			t.Errorf("not sorted\noriginal:%v\nresult:%v", original, a.GetNumsSlice())
		}

		if len(instructions) >= limit {
			// // Uncomment to explore performance stats.
			// // This fails count assumed they are actually being sorted.
			// scores = append(scores, float64(len(instructions)))
			// fails++
			t.Errorf("%v instructions--that's a bit much", len(instructions))
		}
	}

	// // Uncomment and set limit to 0 to see the mean and standard deviation:
	// t.Errorf("fails: %v", fails)
	// t.Errorf("mean: %v", Average(scores))
	// t.Errorf("standard deviation: %v", StandardDeviation(scores))

	// Stack size 6:

	// // Uncomment this and related lines, and adjust limit to explore stats.
	// fails = 0
	// scores = make([]float64, 0, 720)

	// Check that the result is sorted for all permutations of 1-6.
	// Note that this unit test for general doesn't take into account
	// pre-checks in main.go that deal with the case where the stack
	// is already sorted.
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			for k := 1; k <= 6; k++ {
				for l := 1; l <= 6; l++ {
					for m := 1; m <= 6; m++ {
						for n := 1; n <= 6; n++ {
							if (i == j || i == k || i == l || i == m || i == n) || (j == k || j == l || j == m || j == n) || (k == l || k == m || k == n) || (l == m || l == n) || (m == n) {
								continue
							}
							limit := 13 // Must be under this.
							input := fmt.Sprintf("%d %d %d %d %d %d", i, j, k, l, m, n)
							a, err := ps.NewStack(input)
							if err != nil {
								t.Errorf("general(%s) failed: %s", input, err)
							}
							b, _ := ps.NewStack("")
							instructions := general(&a, &b)
							if len(instructions) >= limit {
								// // Uncomment to explore stats.
								// scores = append(scores, float64(len(instructions)))
								// fails++
								t.Errorf("more than %v instructions to sort 6 numbers:\n%v took %v instructions to sort\n%v", limit-1, input, len(instructions), instructions)
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

	// // Uncomment and set limit to 0 to see the mean and standard deviation:
	// if true {
	// 	t.Errorf("fails: %v", fails)
	// 	t.Errorf("mean: %v", Average(scores))
	// 	t.Errorf("standard deviation: %v", StandardDeviation(scores))
	// }

	// Test with different stack sizes. Note that 6, 7, and 8
	// take longer than 9+ because these are the sizes for
	// which we run a BFS to check for better solutions that
	// don't use any pushes. 100 iterations is manageable for 7,
	// but, for 8, try running the test repeatedly with, say, 10
	// iterations.
	for i := 0; i < 10; i++ {
		s := giveMeRandom(7)
		a, err := ps.NewStack(s)
		if err != nil {
			t.Errorf("general(%s) failed: %s", s, err)
		}
		b, _ := ps.NewStack("")
		instructions := general(&a, &b)
		a, _ = ps.NewStack(s)
		b, _ = ps.NewStack("")
		ps.Run(&a, &b, instructions)
		_, sorted := ps.Check(a, b)
		if !sorted {
			t.Errorf("not sorted: want %v, got %v", s, a.GetNumsString())
		}
	}
}

func giveMeRandom(n int) string {
	if n < 1 {
		return ""
	}
	var result string
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}
	rand.Shuffle(n, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	for i, v := range arr {
		result += fmt.Sprintf("%d", v)
		if i != n-1 {
			result += " "
		}
	}
	return result
}

// The first 100 primes greater than 100:
var primes = []int{101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997}

// Generate a string of a hundred random numbers:
func hundredRandomNumbers() string {
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

func Average(data []float64) float64 {
	var sum float64
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	return sum / float64(len(data))
}

func Variance(data []float64) float64 {
	m := Average(data)
	var d []float64
	for i := 0; i < len(data); i++ {
		d = append(d, (data[i]-m)*(data[i]-m))
	}
	return Average(d)
}
func StandardDeviation(data []float64) float64 {
	return math.Sqrt(Variance(data))
}
