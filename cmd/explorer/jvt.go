package main

// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"push-swap/ps"
// )

// // Comparing Ali's "turk" algorithm to variants of Jamie Dawson's.
// func jvt() {
// 	for n := 500; n < 501; n++ {
// 		turkScores := make([]float64, 0, 100)
// 		j5Scores := make([]float64, 0, 100)
// 		tally := 0
// 		for i := 0; i < 100; i++ {
// 			input := randomInput(n)
// 			// j5 modifies the stacks.
// 			a, _ := ps.NewStack(input)
// 			b, _ := ps.NewStack("")
// 			j5 := bucket(&a, &b, 26)

// 			// turk modifies the stack
// 			a, _ = ps.NewStack(input)
// 			b, _ = ps.NewStack("")
// 			// turk := turk(&a, &b)
// 			turk := bucket(&a, &b, 27)

// 			turkScores = append(turkScores, float64(len(turk)))
// 			j5Scores = append(j5Scores, float64(len(j5)))

// 			if len(turk) < len(j5) {
// 				tally++
// 				// fmt.Printf("Test failed for input: %s. Length of `turk` is %v, less than length of `orion` %v\n", input, len(turk), len(orion))
// 			}
// 		}
// 		fmt.Println(n, tally)
// 		fmt.Println("turk mean:", mean(turkScores))
// 		fmt.Println("turk standard deviation:", std(turkScores))
// 		fmt.Println("other mean:", mean(j5Scores))
// 		fmt.Println("other standard deviation:", std(j5Scores))
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

// func bucket(a, b *ps.Stack, size int) []string {
// 	result := []string{}
// 	n := len(a.Nums)
// 	*a, _ = ps.NewStack(rank(a.Nums))
// 	for k := size; k <= n-n%size; k += size {
// 		for len(b.Nums) < k {
// 			A := a.GetNumsSlice()
// 			B := b.GetNumsSlice()
// 			for i := 0; i <= len(a.Nums)/2; i++ {
// 				rotsA := 0
// 				rotsB := 0
// 				upB := true
// 				if A[i] <= k {
// 					for j := i; j > 0; j-- {
// 						rotsA++
// 					}
// 					if len(b.Nums) > 0 {
// 						iMin, m := ps.MinInt(B)
// 						iMax, _ := ps.MaxInt(B)
// 						targetIndex := iMin
// 						if A[i] < m {
// 							targetIndex = iMax
// 						} else {
// 							for j := 0; j < len(B); j++ {
// 								if B[j] < A[i] && B[j] > B[targetIndex] {
// 									targetIndex = j
// 								}
// 							}
// 						}
// 						if targetIndex > len(B)/2 {
// 							upB = false
// 							for j := targetIndex; j < len(B); j++ {
// 								rotsB++
// 							}
// 						} else {
// 							for j := targetIndex; j > 0; j-- {
// 								rotsB++
// 							}
// 						}
// 					}
// 					if upB {
// 						shared := 0
// 						shared = min(rotsA, rotsB)
// 						rotsA -= shared
// 						rotsB -= shared
// 						for j := 0; j < shared; j++ {
// 							ps.Rr(a, b)
// 							result = append(result, "rr")
// 						}
// 					}
// 					for j := 0; j < rotsB; j++ {
// 						if upB {
// 							ps.Rx(b)
// 							result = append(result, "rb")
// 						} else {
// 							ps.Rrx(b)
// 							result = append(result, "rrb")
// 						}
// 					}
// 					for j := 0; j < rotsA; j++ {
// 						ps.Rx(a)
// 						result = append(result, "ra")
// 					}
// 					ps.Px(b, a)
// 					result = append(result, "pb")
// 					break
// 				}

// 				if A[len(A)-i-1] <= k {
// 					for j := len(A) - i - 1; j < len(A); j++ {
// 						rotsA++
// 					}
// 					if len(b.Nums) > 0 {
// 						iMin, m := ps.MinInt(B)
// 						iMax, _ := ps.MaxInt(B)
// 						targetIndex := iMin
// 						if A[len(A)-i-1] < m {
// 							targetIndex = iMax
// 						} else {
// 							for j := 0; j < len(B); j++ {
// 								if B[j] < A[len(A)-i-1] && B[j] > B[targetIndex] {
// 									targetIndex = j
// 								}
// 							}
// 						}
// 						if targetIndex > len(B)/2 {
// 							upB = false
// 							for j := targetIndex; j < len(B); j++ {
// 								rotsB++
// 							}
// 						} else {
// 							for j := targetIndex; j > 0; j-- {
// 								rotsB++
// 							}
// 						}
// 					}
// 					if !upB {
// 						shared := 0
// 						shared = min(rotsA, rotsB)
// 						rotsA -= shared
// 						rotsB -= shared
// 						for j := 0; j < shared; j++ {
// 							ps.Rrr(a, b)
// 							result = append(result, "rrr")
// 						}
// 					}
// 					for j := 0; j < rotsB; j++ {
// 						if upB {
// 							ps.Rx(b)
// 							result = append(result, "rb")
// 						} else {
// 							ps.Rrx(b)
// 							result = append(result, "rrb")
// 						}
// 					}
// 					for j := 0; j < rotsA; j++ {
// 						ps.Rrx(a)
// 						result = append(result, "rra")
// 					}
// 					ps.Px(b, a)
// 					result = append(result, "pb")
// 					break
// 				}
// 			}
// 		}
// 	}

// 	if len(a.Nums) == 0 {
// 		ps.Px(a, b)
// 		result = result[:len(result)-1]
// 	} else {
// 		for len(a.Nums) > 1 {
// 			ps.Px(b, a)
// 			result = append(result, "pb")
// 		}
// 	}

// 	result = append(result, insert(b, a, 0, false)...)
// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, justRotate(*a))
// 	return result
// }

// // Jamie Dawson's but with descending order on B, and the additional
// // optimization of shared rotations.
// func jd(a, b *ps.Stack) []string {
// 	result := []string{}
// 	*a, _ = ps.NewStack(rank(a.Nums))
// 	for k := 20; k <= 100; k += 20 {
// 		for len(b.Nums) < k {
// 			A := a.GetNumsSlice()
// 			B := b.GetNumsSlice()
// 			for i := 0; i <= len(a.Nums)/2; i++ {
// 				rotsA := 0
// 				rotsB := 0
// 				upB := true
// 				if A[i] <= k {
// 					for j := i; j > 0; j-- {
// 						rotsA++
// 					}
// 					if len(b.Nums) > 0 {
// 						iMin, m := ps.MinInt(B)
// 						iMax, _ := ps.MaxInt(B)
// 						targetIndex := iMin
// 						if A[i] < m {
// 							targetIndex = iMax
// 						} else {
// 							for j := 0; j < len(B); j++ {
// 								if B[j] < A[i] && B[j] > B[targetIndex] {
// 									targetIndex = j
// 								}
// 							}
// 						}
// 						if targetIndex > len(B)/2 {
// 							upB = false
// 							for j := targetIndex; j < len(B); j++ {
// 								rotsB++
// 							}
// 						} else {
// 							for j := targetIndex; j > 0; j-- {
// 								rotsB++
// 							}
// 						}
// 					}
// 					if upB {
// 						shared := 0
// 						shared = min(rotsA, rotsB)
// 						rotsA -= shared
// 						rotsB -= shared
// 						for j := 0; j < shared; j++ {
// 							ps.Rr(a, b)
// 							result = append(result, "rr")
// 						}
// 					}
// 					for j := 0; j < rotsB; j++ {
// 						if upB {
// 							ps.Rx(b)
// 							result = append(result, "rb")
// 						} else {
// 							ps.Rrx(b)
// 							result = append(result, "rrb")
// 						}
// 					}
// 					for j := 0; j < rotsA; j++ {
// 						ps.Rx(a)
// 						result = append(result, "ra")
// 					}
// 					ps.Px(b, a)
// 					result = append(result, "pb")
// 					break
// 				}

// 				if A[len(A)-i-1] <= k {
// 					for j := len(A) - i - 1; j < len(A); j++ {
// 						rotsA++
// 					}
// 					if len(b.Nums) > 0 {
// 						iMin, m := ps.MinInt(B)
// 						iMax, _ := ps.MaxInt(B)
// 						targetIndex := iMin
// 						if A[len(A)-i-1] < m {
// 							targetIndex = iMax
// 						} else {
// 							for j := 0; j < len(B); j++ {
// 								if B[j] < A[len(A)-i-1] && B[j] > B[targetIndex] {
// 									targetIndex = j
// 								}
// 							}
// 						}
// 						if targetIndex > len(B)/2 {
// 							upB = false
// 							for j := targetIndex; j < len(B); j++ {
// 								rotsB++
// 							}
// 						} else {
// 							for j := targetIndex; j > 0; j-- {
// 								rotsB++
// 							}
// 						}
// 					}
// 					if !upB {
// 						shared := 0
// 						shared = min(rotsA, rotsB)
// 						rotsA -= shared
// 						rotsB -= shared
// 						for j := 0; j < shared; j++ {
// 							ps.Rrr(a, b)
// 							result = append(result, "rrr")
// 						}
// 					}
// 					for j := 0; j < rotsB; j++ {
// 						if upB {
// 							ps.Rx(b)
// 							result = append(result, "rb")
// 						} else {
// 							ps.Rrx(b)
// 							result = append(result, "rrb")
// 						}
// 					}
// 					for j := 0; j < rotsA; j++ {
// 						ps.Rrx(a)
// 						result = append(result, "rra")
// 					}
// 					ps.Px(b, a)
// 					result = append(result, "pb")
// 					break
// 				}
// 			}
// 		}
// 	}
// 	for len(b.Nums) > 0 {
// 		ps.Px(a, b)
// 		result = append(result, "pa")
// 	}
// 	finalRotations := justRotate(*a)
// 	ps.Run(a, b, finalRotations)
// 	result = append(result, finalRotations...)
// 	return result
// }
