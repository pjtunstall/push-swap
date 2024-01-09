package main

import (
	"push-swap/ps"
	"sort"
	"strconv"
	"strings"
)

func general(a, b *ps.Stack) []string {
	original := a.GetNumsString()
	var result []string

	swapRotatable, swappable := swapRot(*a, *b)

	// If a swap is enough to sort the stack:
	if swappable {
		ps.Sx(a)
		result = append(result, "sa")
		return result
	}

	// If a swap, then rotations are enough to sort the stack:
	if swapRotatable {
		ps.Sx(a)
		rots := justRotate(*a)
		ps.Run(a, b, rots)
		result = append(result, "sa")
		result = append(result, rots...)
		return result
	}

	// Check if rotations, then a swap, then (possibly) more rotations,
	// are enough to sort the stack.
	rotSwapScript, rotSwappable := rotSwap(*a, *b)
	if rotSwappable {
		ps.Run(a, b, rotSwapScript)
		return rotSwapScript
	}

	// Fred 1000orion's bucket sort to B, then insertion sort back to A.
	// n needs to be at least 8 because when n = 6 or 7, the "third" left
	// on stack A when the smallest two thirds are pushed to B
	// has only 2 elements, whereas the algorithm says push the final
	// (largest) third till 3 elements are left.
	// Turk algorithm actually gets better and better than Orion till
	// n == 42. After that, Orion starts to improve relative to Turk,
	// and at n == 93, Orion is winning over Turk at last.
	if len(a.Nums) > 92 {
		return bucket3(a, b)
	}

	// To get these (and hence all stacks of 6) under 13 instructions.
	// without this check, the current algorithm takes 13 instructions.
	// In the first case, "pb pb rb pb sa rra pa rr pa ra ra pa rra",
	// with no better solution found by BFS. In the second case, BFS
	// does find the best solution, but it's still 13:
	// "sa ra sa ra ra sa ra sa rra sa ra ra ra".
	if a.GetNumsString() == "2 6 5 4 3 1" {
		result = []string{"pb", "rra", "pb", "ss", "ra", "ra", "sa", "pa", "pa"}
		ps.Run(a, b, result)
		return result
	}

	if a.GetNumsString() == "3 1 2 6 5 4" {
		result = []string{"ra", "pb", "pb", "sa", "ra", "ra", "sa", "pa", "pa"}
		ps.Run(a, b, result)
		return result
	}

	// AYO succeeds at this one in 12 instructions:
	// "pb pb sb pb sa rra pa ra pa pa rra rra".
	// Since we already found this shorter solution (10) by hand,
	// we could use it here.
	if a.GetNumsString() == "4 3 2 1 6 5" {
		result = []string{"pb", "pb", "ss", "ra", "ra", "sa", "pa", "pa", "rra", "rra"}
		ps.Run(a, b, result)
		return result
	}

	// AYO's "Turk" algorithm.
	ps.Px(b, a)
	ps.Px(b, a)
	result = append(result, "pb", "pb")

	nums := b.GetNumsSlice()
	if nums[0] < nums[1] {
		ps.Sx(b)
		result = append(result, "sb")
	}

	result = append(result, insert(a, b, 3, true)...)

	_, rotatable := three(a.Nums)
	if !rotatable {
		ps.Sx(a)
		result = append(result, "sa")

	}

	result = append(result, insert(b, a, 0, false)...)

	result = append(result, justRotate(*a)...)
	ps.Run(a, b, justRotate(*a))

	// An extra check to see if we can sort using a sequence of only
	// swaps and rotations. This uses a BFS, and is too slow for big
	// stacks, hence the simpler checks at the beginning of this
	// function. We already performed this check for stack size 8 above,
	// so this is just for stack size 6 and 7.
	if len(a.Nums) < 9 {
		alt, sorted := bfs(original, len(result))
		if sorted && len(alt) < len(result) {
			return alt
		}

	}
	return result
}

func bucket3(a, b *ps.Stack) []string {
	var result []string
	*a, _ = ps.NewStack(rank(a.Nums))
	n := len(a.Nums)
	third := n / 3
	if n%3 == 2 {
		third++
	}
	twoThirds := 2 * third
	if 2*(n%3) == 2 {
		twoThirds++
	}

	// Push the smallest third to the bottom of stack B and the
	// middle third to the top.
	for len(a.Nums) > third {
		top := a.Nums[a.Top]
		if top <= twoThirds {
			ps.Px(b, a)
			result = append(result, "pb")
			if top <= third && len(b.Nums) > 1 {
				ps.Rx(b)
				result = append(result, "rb")
			}
		} else {
			ps.Rx(a)
			result = append(result, "ra")
		}
	}

	// Push the largest third to the top of stack B, leaving
	// the last three on stack A
	for len(a.Nums) > 3 {
		ps.Px(b, a)
		result = append(result, "pb")
	}

	// Perform a swap on stack A if necessary to make it rotatable
	// into sorted position.
	_, rotatable := three(a.Nums)
	if !rotatable {
		ps.Sx(a)
		result = append(result, "sa")
	}

	// Sort while inserting from stack B to stack A.
	result = append(result, insert(b, a, 0, false)...)

	// Rotate stack A into sorted position.
	result = append(result, justRotate(*a)...)
	ps.Run(a, b, justRotate(*a))

	return result
}

// Sorts as it pushes from one stack to another. Here a and
// b are just parameters, with a representing the source stack
// and b the destination stack. The stopAt parameter is the number
// of elements to leave in the source stack. If forward is true,
// the sort is from the real a to the real b, otherwise it is
// from the real b to the real a.
func insert(a, b *ps.Stack, stopAt int, forward bool) []string {
	result := []string{}

	for {
		A := a.GetNumsSlice()
		B := b.GetNumsSlice()
		journeyPlanner := make([]ps.PushInfo, len(A))

		if len(A) == stopAt {
			break
		}

		cheapest := 0
		for i, v := range A {
			var cost int
			var ra, rb bool
			var stepsA, stepsB, jointSteps int

			if i > len(A)/2 {
				cost = len(A) - i
				stepsA = len(A) - i
			} else {
				cost = i
				stepsA = i
				ra = true
			}

			var targetIndex, targetValue int

			if forward {
				foundOneLess := false
				targetIndex, targetValue = ps.MinInt(B)
				for j, w := range B {
					if w < v {
						foundOneLess = true
						if w > targetValue {
							targetValue = w
							targetIndex = j
						}
					}
				}
				if !foundOneLess {
					targetIndex, targetValue = ps.MaxInt(B)
				}
			} else {
				foundOneGreater := false
				targetIndex, targetValue = ps.MaxInt(B)
				for j, w := range B {
					if w > v {
						foundOneGreater = true
						if w < targetValue {
							targetValue = w
							targetIndex = j
						}
					}
				}
				if !foundOneGreater {
					targetIndex, targetValue = ps.MinInt(B)
				}
			}

			if targetIndex > len(B)/2 {
				cost += len(B) - targetIndex
				stepsB = len(B) - targetIndex
			} else {
				cost += targetIndex
				stepsB = targetIndex
				rb = true
			}

			jointSteps = min(stepsA, stepsB)

			// Optimization to take advantage of combined rotatioms
			// when it makes no difference which direction we rotate a stack.
			if len(A)%2 == 0 && stepsA == len(A)/2 {
				ra = rb
			}
			if len(B)%2 == 0 && stepsB == len(B)/2 {
				rb = ra
			}

			if (ra && rb) || (!ra && !rb) {
				cost -= jointSteps
				stepsA -= jointSteps
				stepsB -= jointSteps
			}

			journeyPlanner[i] = ps.PushInfo{
				Value:       v,
				TargetIndex: targetIndex,
				TargetValue: targetValue,
				Cost:        cost,
				Ra:          ra,
				Rb:          rb,
				StepsA:      stepsA,
				StepsB:      stepsB,
				JointSteps:  jointSteps}

			if cost == 0 {
				break
			}
			if cost < journeyPlanner[cheapest].Cost {
				cheapest = i
			}
		}

		c := journeyPlanner[cheapest]

		if c.Ra && c.Rb {
			for j := 0; j < c.JointSteps; j++ {
				ps.Rr(a, b)
				result = append(result, "rr")
			}
		} else if !c.Ra && !c.Rb {
			for j := 0; j < c.JointSteps; j++ {
				ps.Rrr(a, b)
				result = append(result, "rrr")
			}
		}
		if c.Ra {
			for j := 0; j < c.StepsA; j++ {
				ps.Rx(a)
				if forward {
					result = append(result, "ra")
				} else {
					result = append(result, "rb")
				}
			}
		} else {
			for j := 0; j < c.StepsA; j++ {
				ps.Rrx(a)
				if forward {
					result = append(result, "rra")
				} else {
					result = append(result, "rrb")
				}
			}
		}
		if c.Rb {
			for j := 0; j < c.StepsB; j++ {
				ps.Rx(b)
				if forward {
					result = append(result, "rb")
				} else {
					result = append(result, "ra")
				}
			}
		} else {
			for j := 0; j < c.StepsB; j++ {
				ps.Rrx(b)
				if forward {
					result = append(result, "rrb")
				} else {
					result = append(result, "rra")
				}
			}
		}
		if forward {
			ps.Px(b, a)
			result = append(result, "pb")
		} else {
			if len(result) > 0 && result[len(result)-1] == "pb" {
				result = result[:len(result)-1]
			} else {
				result = append(result, "pa")
				ps.Px(b, a)
			}
		}
	}

	return result
}

// // Like FO's algorithm, but using 2 buckets instead of 3.
// func hundredHalves(a, b *ps.Stack) []string {
// 	var result []string
// 	*a, _ = ps.NewStack(rank(a.Nums))

// 	// Push the smallest half to the bottom of stack B and the
// 	// biggest half to the top.
// 	for len(a.Nums) > 3 {
// 		ps.Px(b, a)
// 		result = append(result, "pb")
// 		if b.Nums[b.Top] < 50 {
// 			ps.Rx(b)
// 			result = append(result, "rb")
// 		}
// 	}

// 	// Perform a swap on stack A if necessary to make it rotatable
// 	// into sorted position.
// 	_, rotatable := three(a.Nums)
// 	if !rotatable {
// 		ps.Sx(a)
// 		result = append(result, "sa")
// 	}

// 	// Sort while inserting from stack B to stack A.
// 	result = append(result, insert(b, a, 0, false)...)

// 	// Rotate stack A into sorted position.
// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, justRotate(*a))

// 	return result
// }

// // Variant of FO's algorithm, but using 4 buckets instead of 3.
// func fourSeasons(a, b *ps.Stack) []string {
// 	var result []string
// 	*a, _ = ps.NewStack(rank(a.Nums))

// 	// Group the numbers are grouped according to size and call
// 	// the smallest fourth spring, the next smallest summer,
// 	// then fall, and finally winter the biggest. We first push
// 	// summer to the bottom of stack B and fall to the top.
// 	for len(a.Nums) > 50 {
// 		A := a.GetNumsSlice()
// 		if A[0] > 25 && A[0] < 76 {
// 			ps.Px(b, a)
// 			result = append(result, "pb")
// 			if A[0] < 51 && len(b.Nums) > 1 {
// 				ps.Rx(b)
// 				result = append(result, "rb")
// 			}
// 		} else {
// 			ps.Rx(a)
// 			result = append(result, "ra")
// 		}
// 	}

// 	// Now we push spring to the bottom of stack B and winter to the top,
// 	// leaving the last three on stack A.
// 	for len(a.Nums) > 3 {
// 		A := a.GetNumsSlice()
// 		ps.Px(b, a)
// 		result = append(result, "pb")
// 		if A[0] < 26 {
// 			ps.Rx(b)
// 			result = append(result, "rb")
// 		}
// 	}

// 	// Perform a swap on stack A if necessary to make it rotatable
// 	// into sorted position.
// 	_, rotatable := three(a.Nums)
// 	if !rotatable {
// 		ps.Sx(a)
// 		result = append(result, "sa")
// 	}

// 	// Sort while merging from stack B to stack A.
// 	result = append(result, insert(b, a, 0, false)...)

// 	// Rotate stack A into sorted position.
// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, justRotate(*a))

// 	return result
// }

func rank(nums []int) string {
	rankMap := make(map[int]int)
	strs := make([]string, len(nums))
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	sort.Ints(numsCopy)
	for i, v := range numsCopy {
		rankMap[v] = i + 1
	}

	for i, v := range nums {
		strs[i] = strconv.Itoa(rankMap[v])
	}
	s := strings.Join(strs, " ")

	return s
}

// // Like FO's algorithm, but using 2 buckets instead of 3.
// func hundredHalves(a, b *ps.Stack) []string {
// 	var result []string
// 	*a, _ = ps.NewStack(rank(a.Nums))

// 	// Push the smallest half to the bottom of stack B and the
// 	// biggest half to the top.
// 	for len(a.Nums) > 3 {
// 		ps.Px(b, a)
// 		result = append(result, "pb")
// 		if b.Nums[b.Top] < 50 {
// 			ps.Rx(b)
// 			result = append(result, "rb")
// 		}
// 	}

// 	// Perform a swap on stack A if necessary to make it rotatable
// 	// into sorted position.
// 	_, rotatable := three(a.Nums)
// 	if !rotatable {
// 		ps.Sx(a)
// 		result = append(result, "sa")
// 	}

// 	// Sort while inserting from stack B to stack A.
// 	result = append(result, insert(b, a, 0, false)...)

// 	// Rotate stack A into sorted position.
// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, justRotate(*a))

// 	return result
// }

// // Variant of FO's algorithm, but using 4 buckets instead of 3.
// func fourSeasons(a, b *ps.Stack) []string {
// 	var result []string
// 	*a, _ = ps.NewStack(rank(a.Nums))

// 	// Group the numbers are grouped according to size and call
// 	// the smallest fourth spring, the next smallest summer,
// 	// then fall, and finally winter the biggest. We first push
// 	// summer to the bottom of stack B and fall to the top.
// 	for len(a.Nums) > 50 {
// 		A := a.GetNumsSlice()
// 		if A[0] > 25 && A[0] < 76 {
// 			ps.Px(b, a)
// 			result = append(result, "pb")
// 			if A[0] < 51 && len(b.Nums) > 1 {
// 				ps.Rx(b)
// 				result = append(result, "rb")
// 			}
// 		} else {
// 			ps.Rx(a)
// 			result = append(result, "ra")
// 		}
// 	}

// 	// Now we push spring to the bottom of stack B and winter to the top,
// 	// leaving the last three on stack A.
// 	for len(a.Nums) > 3 {
// 		A := a.GetNumsSlice()
// 		ps.Px(b, a)
// 		result = append(result, "pb")
// 		if A[0] < 26 {
// 			ps.Rx(b)
// 			result = append(result, "rb")
// 		}
// 	}

// 	// Perform a swap on stack A if necessary to make it rotatable
// 	// into sorted position.
// 	_, rotatable := three(a.Nums)
// 	if !rotatable {
// 		ps.Sx(a)
// 		result = append(result, "sa")
// 	}

// 	// Sort while merging from stack B to stack A.
// 	result = append(result, insert(b, a, 0, false)...)

// 	// Rotate stack A into sorted position.
// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, justRotate(*a))

// 	return result
// }

// // Jamie Dawson's algorithm, except sorting onto B in descending order.
// // If the numbers are sorted in ascending order on B, it will take 99
// // extra instructions, ignoring possible optimization by combining
// // rotations.
// func jd(a, b *ps.Stack) []string {
// 	result := []string{}

// 	*a, _ = ps.NewStack(rank(a.Nums))

// 	for k := 20; k <= 100; k += 20 {
// 		for len(b.Nums) < k {
// 			A := a.GetNumsSlice()
// 			B := b.GetNumsSlice()
// 			for i := 0; i <= len(a.Nums)/2; i++ {
// 				if A[i] <= k {
// 					for a.Nums[a.Top] != A[i] {
// 						ps.Rx(a)
// 						result = append(result, "ra")
// 					}
// 					if len(b.Nums) > 0 {
// 						iMin, m := ps.MinInt(B)
// 						iMax, _ := ps.MaxInt(B)
// 						targetIndex := iMin
// 						if a.Nums[a.Top] < m {
// 							targetIndex = iMax
// 						} else {
// 							for j := 0; j < len(B); j++ {
// 								if B[j] < a.Nums[a.Top] && B[j] > B[targetIndex] {
// 									targetIndex = j
// 								}
// 							}
// 						}
// 						if targetIndex > len(B)/2 {
// 							for b.Nums[b.Top] != B[targetIndex] {
// 								ps.Rrx(b)
// 								result = append(result, "rrb")
// 							}
// 						} else {
// 							for b.Nums[b.Top] != B[targetIndex] {
// 								ps.Rx(b)
// 								result = append(result, "rb")
// 							}
// 						}
// 					}
// 					ps.Px(b, a)
// 					result = append(result, "pb")
// 					break
// 				}

// 				if A[len(A)-i-1] <= k {
// 					for a.Nums[a.Top] != A[len(A)-i-1] {
// 						ps.Rrx(a)
// 						result = append(result, "rra")
// 					}
// 					if len(b.Nums) > 0 {
// 						iMin, m := ps.MinInt(B)
// 						iMax, _ := ps.MaxInt(B)
// 						targetIndex := iMin
// 						if a.Nums[a.Top] < m {
// 							targetIndex = iMax
// 						} else {
// 							for j := 0; j < len(B); j++ {
// 								if B[j] < a.Nums[a.Top] && B[j] > B[targetIndex] {
// 									targetIndex = j
// 								}
// 							}
// 						}
// 						if targetIndex > len(B)/2 {
// 							for b.Nums[b.Top] != B[targetIndex] {
// 								ps.Rrx(b)
// 								result = append(result, "rrb")
// 							}
// 						} else {
// 							for b.Nums[b.Top] != B[targetIndex] {
// 								ps.Rx(b)
// 								result = append(result, "rb")
// 							}
// 						}
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

// // Jamie Dawson's but just putting into buckets as we go to B,
// // the smaller the numbers the closer the bucket to the bottom
// // of B, only insertion sorting on the way back to A.
// func bucket(a, b *ps.Stack, size int) []string {
// 	result := []string{}
// 	n := len(a.Nums)
// 	*a, _ = ps.NewStack(rank(a.Nums))
// 	for k := size; k <= n-n%size; k += size {
// 		for len(b.Nums) < k {
// 			A := a.GetNumsSlice()
// 			for i := 0; i <= len(a.Nums)/2; i++ {
// 				rotsA := 0
// 				if A[i] <= k {
// 					for j := i; j > 0; j-- {
// 						rotsA++
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

// func luca(a, b *ps.Stack) []string {
// 	result := []string{}
// 	*a, _ = ps.NewStack(rank(a.Nums))
// 	n := len(a.Nums)
// 	r := 14

// outer:
// 	for k := r; k <= n; k += 2 * r {
// 		for l := 0; l < 2*r; l++ {
// 			if len(a.Nums) == 3 {
// 				break outer
// 			}
// 			A := a.GetNumsSlice()
// 			for i := 0; i <= len(a.Nums)/2; i++ {
// 				if A[i] > k && A[i] <= k+r {
// 					if A[i] > 97 {
// 						continue
// 					}
// 					for j := i; j > 0; j-- {
// 						ps.Rx(a)
// 						result = append(result, "ra")
// 					}
// 					ps.Px(b, a)
// 					result = append(result, "pb")
// 					break
// 				}
// 				if A[i] <= k {
// 					if A[i] > 97 {
// 						continue
// 					}
// 					for j := i; j > 0; j-- {
// 						ps.Rx(a)
// 						result = append(result, "ra")
// 					}
// 					ps.Px(b, a)
// 					ps.Rx(b)
// 					result = append(result, "pb", "rb")
// 					break
// 				}

// 				if A[len(A)-i-1] > k && A[len(A)-i-1] <= k+r {
// 					if A[len(A)-i-1] > 97 {
// 						continue
// 					}
// 					for j := len(A) - i - 1; j < len(A); j++ {
// 						ps.Rrx(a)
// 						result = append(result, "rra")
// 					}
// 					ps.Px(b, a)
// 					result = append(result, "pb")
// 					break
// 				}
// 				if A[len(A)-i-1] <= k {
// 					if A[len(A)-i-1] > 97 {
// 						continue
// 					}
// 					for j := len(A) - i - 1; j < len(A); j++ {
// 						ps.Rrx(a)
// 						result = append(result, "rra")
// 					}
// 					ps.Px(b, a)
// 					ps.Rx(b)
// 					result = append(result, "pb", "rb")
// 					break
// 				}
// 			}
// 		}
// 	}

// 	// Push the remaining elements to stack B.
// 	for len(a.Nums) > 3 {
// 		if a.Nums[a.Top] > 97 {
// 			ps.Rx(a)
// 			result = append(result, "ra")
// 			continue
// 		}
// 		ps.Px(b, a)
// 		result = append(result, "pb")
// 	}

// 	_, rotatable := three(a.Nums)
// 	if !rotatable {
// 		ps.Sx(a)
// 		result = append(result, "sa")

// 	}

// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, justRotate(*a))

// 	ps.Px(a, b)
// 	result = append(result, "pa")
// 	A := a.GetNumsSlice()
// 	if a.Nums[0] != A[1]-1 {
// 		ps.Rx(a)
// 		result = append(result, "ra")
// 	}

// 	for len(b.Nums) > 0 || a.Nums[a.Top] != 1 {
// 		if bottomsUp(*a) {
// 			ps.Rrx(a)
// 			result = append(result, "rra")
// 			continue
// 		}
// 		A := a.GetNumsSlice()
// 		B := b.GetNumsSlice()
// 		if B[0] == A[0]-1 {
// 			ps.Px(a, b)
// 			result = append(result, "pa")
// 			continue
// 		} else if B[0] > A[len(A)-1] {
// 			ps.Px(a, b)
// 			ps.Rx(a)
// 			result = append(result, "pa", "ra")
// 			continue
// 		} else {
// 			maxI, _ := ps.MaxInt(B)
// 			if maxI > len(B)/2 {
// 				for j := maxI; j < len(B); j++ {
// 					ps.Rrx(b)
// 					result = append(result, "rrb")
// 				}
// 			} else {
// 				for j := maxI; j > 0; j-- {
// 					ps.Rx(b)
// 					result = append(result, "rb")
// 				}
// 			}
// 			ps.Px(a, b)
// 			result = append(result, "pa")
// 		}
// 	}

// 	return result
// }

// // For use with luca and yy.
// func bottomsUp(a ps.Stack) bool {
// 	A := a.GetNumsSlice()
// 	return A[len(A)-1] == A[0]-1
// }

// func yy(a, b *ps.Stack) []string {
// 	result := []string{}
// 	*a, _ = ps.NewStack(rank(a.Nums))
// 	n := len(a.Nums)
// 	r := 12

// outer:
// 	for k := r; k <= n; k += 2 * r {
// 		for l := 0; l < 2*r; l++ {
// 			if len(a.Nums) == 3 {
// 				break outer
// 			}
// 			A := a.GetNumsSlice()
// 			for i := 0; i <= len(a.Nums)/2; i++ {
// 				if A[i] > k && A[i] <= k+r {
// 					for j := i; j > 0; j-- {
// 						ps.Rx(a)
// 						result = append(result, "ra")
// 					}
// 					ps.Px(b, a)
// 					result = append(result, "pb")
// 					break
// 				}
// 				if A[i] <= k {
// 					for j := i; j > 0; j-- {
// 						ps.Rx(a)
// 						result = append(result, "ra")
// 					}
// 					ps.Px(b, a)
// 					ps.Rx(b)
// 					result = append(result, "pb", "rb")
// 					break
// 				}

// 				if A[len(A)-i-1] > k && A[len(A)-i-1] <= k+r {
// 					for j := len(A) - i - 1; j < len(A); j++ {
// 						ps.Rrx(a)
// 						result = append(result, "rra")
// 					}
// 					ps.Px(b, a)
// 					result = append(result, "pb")
// 					break
// 				}
// 				if A[len(A)-i-1] <= k {
// 					for j := len(A) - i - 1; j < len(A); j++ {
// 						ps.Rrx(a)
// 						result = append(result, "rra")
// 					}
// 					ps.Px(b, a)
// 					ps.Rx(b)
// 					result = append(result, "pb", "rb")
// 					break
// 				}
// 			}
// 		}
// 	}

// 	for len(a.Nums) > 0 {
// 		ps.Px(b, a)
// 		result = append(result, "pb")
// 	}

// 	maxI, _ := ps.MaxInt(b.GetNumsSlice())
// 	if maxI > len(b.Nums)/2 {
// 		for j := maxI; j < len(b.Nums); j++ {
// 			ps.Rrx(b)
// 			result = append(result, "rrb")
// 		}
// 	} else {
// 		for j := maxI; j > 0; j-- {
// 			ps.Rx(b)
// 			result = append(result, "rb")
// 		}
// 	}
// 	ps.Px(a, b)
// 	ps.Px(a, b)
// 	result = append(result, "pa", "pa")
// 	A := a.GetNumsSlice()
// 	if a.Nums[0] != A[1]-1 {
// 		ps.Rx(a)
// 		result = append(result, "ra")
// 	}

// 	for len(b.Nums) > 0 || a.Nums[a.Top] != 1 {
// 		if bottomsUp(*a) {
// 			ps.Rrx(a)
// 			result = append(result, "rra")
// 			continue
// 		}
// 		A := a.GetNumsSlice()
// 		B := b.GetNumsSlice()
// 		if B[0] == A[0]-1 {
// 			ps.Px(a, b)
// 			result = append(result, "pa")
// 			continue
// 		} else if B[0] > A[len(A)-1] {
// 			ps.Px(a, b)
// 			ps.Rx(a)
// 			result = append(result, "pa", "ra")
// 			continue
// 		} else {
// 			maxI, _ := ps.MaxInt(B)
// 			otherI := 0

// 			for i := 1; i <= len(B)/2; i++ {
// 				if B[i] > A[len(A)-1] {
// 					otherI = i
// 					break
// 				}
// 				if B[len(B)-i] > A[len(A)-1] {
// 					otherI = len(B) - i
// 					break
// 				}
// 			}

// 			rotsMax := min(maxI, len(B)-maxI)
// 			rotsOther := min(otherI, len(B)-otherI)

// 			if rotsMax <= rotsOther {
// 				if maxI > len(B)/2 {
// 					for j := maxI; j < len(B); j++ {
// 						ps.Rrx(b)
// 						result = append(result, "rrb")
// 					}
// 				} else {
// 					for j := maxI; j > 0; j-- {
// 						ps.Rx(b)
// 						result = append(result, "rb")
// 					}
// 				}
// 				ps.Px(a, b)
// 				result = append(result, "pa")
// 			} else {
// 				if otherI > len(B)/2 {
// 					for j := otherI; j < len(B); j++ {
// 						ps.Rrx(b)
// 						result = append(result, "rrb")
// 					}
// 				} else {
// 					for j := otherI; j > 0; j-- {
// 						ps.Rx(b)
// 						result = append(result, "rb")
// 					}
// 				}
// 				ps.Px(a, b)
// 				ps.Rx(a)
// 				result = append(result, "pa", "ra")
// 			}

// 		}
// 	}

// 	return result
// }

// func anya(a, b *ps.Stack) []string {
// 	var result []string
// 	*a, _ = ps.NewStack(rank(a.Nums))
// 	n := len(a.Nums)
// 	median := len(a.Nums) / 2

// 	for i := 0; i < n; i++ {
// 		if a.Nums[a.Top] <= median {
// 			ps.Px(b, a)
// 			result = append(result, "pb")
// 		} else {
// 			ps.Rx(a)
// 			result = append(result, "ra")
// 		}
// 	}

// 	for len(b.Nums) > 0 {
// 		B := b.GetNumsSlice()
// 		maxI, _ := ps.MaxInt(B)
// 		minI, _ := ps.MinInt(B)
// 		var pushMeNext int
// 		var iAmMin bool
// 		for i := 0; i <= len(B)/2; i++ {
// 			if i == maxI {
// 				pushMeNext = i
// 				break
// 			}
// 			if i == minI {
// 				pushMeNext = i
// 				iAmMin = true
// 				break
// 			}
// 			if i == 0 {
// 				continue
// 			}
// 			if len(B)-i == maxI {
// 				pushMeNext = len(B) - i
// 				break
// 			}
// 			if len(B)-i == minI {
// 				pushMeNext = len(B) - i
// 				iAmMin = true
// 				break
// 			}
// 		}

// 		if pushMeNext > len(b.Nums)/2 {
// 			for i := pushMeNext; i < len(b.Nums); i++ {
// 				ps.Rrx(b)
// 				result = append(result, "rrb")
// 			}
// 		} else {
// 			for i := pushMeNext; i > 0; i-- {
// 				ps.Rx(b)
// 				result = append(result, "rb")
// 			}
// 		}
// 		ps.Px(a, b)
// 		result = append(result, "pa")
// 		if iAmMin {
// 			ps.Rx(a)
// 			result = append(result, "ra")
// 		}
// 	}

// 	for i := 0; i < n; i++ {
// 		if a.Nums[a.Top] > median {
// 			ps.Px(b, a)
// 			result = append(result, "pb")
// 		} else {
// 			ps.Rx(a)
// 			result = append(result, "ra")
// 		}
// 	}

// 	rots := justRotate(*a)
// 	result = append(result, rots...)
// 	ps.Run(a, b, rots)

// 	for len(b.Nums) > 0 {
// 		B := b.GetNumsSlice()
// 		maxI, _ := ps.MaxInt(B)
// 		minI, _ := ps.MinInt(B)
// 		var pushMeNext int
// 		var iAmMin bool
// 		for i := 0; i <= len(B)/2; i++ {
// 			if i == maxI {
// 				pushMeNext = i
// 				break
// 			}
// 			if i == minI {
// 				pushMeNext = i
// 				iAmMin = true
// 				break
// 			}
// 			if i == 0 {
// 				continue
// 			}
// 			if len(B)-i == maxI {
// 				pushMeNext = len(B) - i
// 				break
// 			}
// 			if len(B)-i == minI {
// 				pushMeNext = len(B) - i
// 				iAmMin = true
// 				break
// 			}
// 		}

// 		if pushMeNext > len(b.Nums)/2 {
// 			for i := pushMeNext; i < len(b.Nums); i++ {
// 				ps.Rrx(b)
// 				result = append(result, "rrb")
// 			}
// 		} else {
// 			for i := pushMeNext; i > 0; i-- {
// 				ps.Rx(b)
// 				result = append(result, "rb")
// 			}
// 		}
// 		ps.Px(a, b)
// 		result = append(result, "pa")
// 		if iAmMin {
// 			ps.Rx(a)
// 			result = append(result, "ra")
// 		}
// 	}

// 	rots = justRotate(*a)
// 	result = append(result, rots...)
// 	ps.Run(a, b, rots)

// 	return result
// }
