// A test to compare the performance or turk and orion.
// This one does actually work. Move to cmd/push-swap and
// call it from the start of main.

package main

// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"push-swap/ps"
// )

// func ovt() {
// 	for n := 8; n < 101; n++ {
// 		turkScores := make([]float64, 0, 100)
// 		orionScores := make([]float64, 0, 100)
// 		tally := 0
// 		for i := 0; i < 10000; i++ {
// 			input := randomInput(n)
// 			// orion modifies the stacks.
// 			a, _ := ps.NewStack(input)
// 			b, _ := ps.NewStack("")
// 			orion := bucket3(&a, &b)

// 			// turk modifies the stack
// 			a, _ = ps.NewStack(input)
// 			b, _ = ps.NewStack("")
// 			turk := turk(&a, &b)

// 			turkScores = append(turkScores, float64(len(turk)))
// 			orionScores = append(orionScores, float64(len(orion)))

// 			if len(turk) < len(orion) {
// 				tally++
// 				// fmt.Printf("Test failed for input: %s. Length of `turk` is %v, less than length of `orion` %v\n", input, len(turk), len(orion))
// 			}
// 		}
// 		fmt.Println(n, tally)
// 		fmt.Println("turk mean:", mean(turkScores))
// 		fmt.Println("turk standard deviation:", std(turkScores))
// 		fmt.Println("orion mean:", mean(orionScores))
// 		fmt.Println("orion standard deviation:", std(orionScores))
// 	}
// }

// func turk(a, b *ps.Stack) []string {
// 	result := []string{}
// 	ps.Px(b, a)
// 	ps.Px(b, a)
// 	result = append(result, "pb", "pb")

// 	nums := b.GetNumsSlice()
// 	if nums[0] < nums[1] {
// 		ps.Sx(b)
// 		result = append(result, "sb")
// 	}

// 	result = append(result, insert(a, b, 3, true)...)

// 	_, rotatable := three(a.Nums)
// 	if !rotatable {
// 		ps.Sx(a)
// 		result = append(result, "sa")

// 	}

// 	result = append(result, insert(b, a, 0, false)...)

// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, justRotate(*a))

// 	return result
// }

// func randomInput(n int) string {
// 	if n < 1 {
// 		return ""
// 	}
// 	var result string
// 	arr := make([]int, n)
// 	for i := range arr {
// 		arr[i] = i + 1
// 	}
// 	rand.Shuffle(n, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
// 	for i, v := range arr {
// 		result += fmt.Sprintf("%d", v)
// 		if i != n-1 {
// 			result += " "
// 		}
// 	}
// 	return result
// }

// func mean(data []float64) float64 {
// 	var sum float64
// 	for i := 0; i < len(data); i++ {
// 		sum += data[i]
// 	}
// 	return sum / float64(len(data))
// }

// func variance(data []float64) float64 {
// 	m := mean(data)
// 	var d []float64
// 	for i := 0; i < len(data); i++ {
// 		d = append(d, (data[i]-m)*(data[i]-m))
// 	}
// 	return mean(d)
// }
// func std(data []float64) float64 {
// 	return math.Sqrt(variance(data))
// }
