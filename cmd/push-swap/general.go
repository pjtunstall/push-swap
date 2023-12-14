package main

import (
	"fmt"
	"push-swap/ps"
)

func general(a, b *ps.Stack) []string {
	n := len(a.Nums)
	runSize := n
	numberOfRuns := 1
	// if n > 5 {
	// 	numberOfRuns = n / runSize
	// }

	var result []string

	// Just in case any preliminary checks altered the stacks. May not be necessary.
	*a, _ = ps.NewStack(rank(a.Nums))

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

	if true {
		ps.Px(b, a)
		ps.Px(b, a)
		result = append(result, "pb", "pb")
		nums := b.GetNumsSlice()
		if nums[0] < nums[1] {
			ps.Sx(b)
			result = append(result, "sb")
		}
		result = append(result, sortToB(a, b)...)
		_, rotatable := three(a.Nums)
		if !rotatable {
			ps.Sx(a)
			result = append(result, "sa")
		}
		result = append(result, sortToA(a, b)...)
	} else {
		for runNumber := 0; runNumber < numberOfRuns; runNumber++ {
			result = append(result, sortRuns(a, b, runNumber, runSize)...)
		}
		// TODO: Rotate B shortest way to put the biggest number on top.
		// Adjust whatever value is returned by justRotate(*b) trimming any
		// trailing rrb. If the last value is rb, append another rb.
		// Run and append to result.
		B := b.GetNumsSlice()
		_, minB := ps.MinInt(B)
		if B[0] != minB {
			result = append(result, justRotate(*b)...)
			result = append(result, "rb")
		}
		for i := 0; i < n; i++ {
			ps.Px(a, b)
			result = append(result, "pa")
		}
	}

	result = append(result, justRotate(*a)...)
	ps.Run(a, b, justRotate(*a))

	fmt.Println("A:", a.GetNumsSlice())
	fmt.Println("B:", b.GetNumsSlice())
	fmt.Println()

	return result
}

func sortRuns(a, b *ps.Stack, runNumber int, runSize int) []string {
	result := []string{}

	for {
		A := a.GetNumsSlice()
		if len(A) == 0 {
			break
		}
		B := b.GetNumsSlice()
		if len(B) == (runNumber+1)*runSize {
			break
		}

		var cheapest ps.PushInfo

		// Search A fromm top and bottom smultaneously. The first element
		// found will be the one that takes the least steps to rotate to
		// the top.
		for i := range A {
			if A[i] > runNumber*runSize && A[i] <= (runNumber+1)*runSize {
				cheapest.Index = i
				cheapest.StepsA = i
				cheapest.Value = A[i]
				cheapest.Ra = true
				break
			}

			// Just -i not -1-i because it takes the ith element from the end
			// one more reverse rotation to reach the top than the number of
			// rotations that the ith element takes to reach the top.
			if i != 0 && A[i] > runNumber*runSize && A[i] <= (runNumber+1)*runSize {
				cheapest.Index = len(A) - i
				cheapest.StepsA = i
				cheapest.Value = A[len(A)-i]
				cheapest.Ra = false
				break
			}
		}

		if len(B) == 0 {
			if cheapest.Index != 0 {
				if cheapest.Ra {
					for i := 0; i < cheapest.StepsA; i++ {
						ps.Rx(a)
						result = append(result, "ra")
					}
				} else {
					for i := 0; i < cheapest.StepsA; i++ {
						ps.Rrx(a)
						result = append(result, "rra")
					}
				}
			}
			ps.Px(b, a)
			result = append(result, "pb")
			continue
		}

		foundOneLess := false
		for j, w := range B {
			if w < cheapest.Value {
				foundOneLess = true
			} else {
				if w > cheapest.TargetValue {
					cheapest.TargetValue = w
					cheapest.TargetIndex = j
				}
			}
		}
		if !foundOneLess {
			cheapest.TargetIndex, cheapest.TargetValue = ps.MaxInt(B)
		}

		if cheapest.TargetIndex > len(B)/2 {
			cheapest.StepsB = len(B) - cheapest.TargetIndex
		} else {
			cheapest.StepsB = cheapest.TargetIndex
			cheapest.Rb = true
		}

		cheapest.JointSteps = min(cheapest.StepsA, cheapest.StepsB)

		if (cheapest.Ra && cheapest.Rb) || (!cheapest.Ra && !cheapest.Rb) {
			cheapest.StepsA -= cheapest.JointSteps
			cheapest.StepsB -= cheapest.JointSteps
		}

		fmt.Println("A:", A)
		fmt.Println("B:", B)
		fmt.Println("cheapest:", cheapest)
		fmt.Println()

		if cheapest.Ra && cheapest.Rb {
			for j := 0; j < cheapest.JointSteps; j++ {
				ps.Rr(b, a)
				result = append(result, "rr")
			}
		} else if !cheapest.Ra && !cheapest.Rb {
			for j := 0; j < cheapest.JointSteps; j++ {
				ps.Rrr(a, b)
				result = append(result, "rrr")
			}
		}
		if cheapest.Ra {
			for j := 0; j < cheapest.StepsA; j++ {
				ps.Rx(a)
				result = append(result, "ra")
			}
		} else {
			for j := 0; j < cheapest.StepsA; j++ {
				ps.Rrx(a)
				result = append(result, "rra")
			}
		}
		if cheapest.Rb {
			for j := 0; j < cheapest.StepsB; j++ {
				ps.Rx(b)
				result = append(result, "rb")
			}
		} else {
			for j := 0; j < cheapest.StepsB; j++ {
				ps.Rrx(b)
				result = append(result, "rrb")
			}
		}
		ps.Px(b, a)
		result = append(result, "pb")
	}

	return result
}

func sortToB(a, b *ps.Stack) []string {
	result := []string{}
	for {
		A := a.GetNumsSlice()
		B := b.GetNumsSlice()
		journeyPlanner := make([]ps.PushInfo, len(A))

		if len(A) == 3 {
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

			foundOneLess := false
			targetIndex, targetValue := ps.MinInt(B)
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

			if targetIndex > len(B)/2 {
				cost += len(B) - targetIndex
				stepsB = len(B) - targetIndex
			} else {
				cost += targetIndex
				stepsB = targetIndex
				rb = true
			}

			jointSteps = min(stepsA, stepsB)

			// // Optimization: Account for case where ra XOR rb, but
			// // either of the stacks would be rotated len(X)/2 times.
			// // In that case, it doesn't matter which way we rotate the
			// // stack that needs rotating len(X)/2 times, so we can
			// // set its direction of rotation to the same as the other
			// // to take advantage of any possible joint steps.
			// // But understand why the same code DOESN'T work in
			// // sortToA, and indeed fails to sort.
			// if stepsA == len(A)/2 {
			// 	ra = rb
			// }
			// if stepsB == len(B)/2 {
			// 	rb = ra
			// }

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

		// fmt.Println("A:", A)
		// fmt.Println("B:", B)
		// fmt.Println()

		// fmt.Println("cheapest:", c.Value)
		// fmt.Println("targetIndex:", c.TargetIndex)
		// fmt.Println("targetValue:", c.TargetValue)
		// fmt.Println("cost:", c.Cost)
		// fmt.Println("ra:", c.Ra)
		// fmt.Println("rb:", c.Rb)
		// fmt.Println("stepsA:", c.StepsA)
		// fmt.Println("stepsB:", c.StepsB)
		// fmt.Println("jointSteps:", c.JointSteps)
		// fmt.Println()

		if c.Ra && c.Rb {
			for j := 0; j < c.JointSteps; j++ {
				ps.Rr(b, a)
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
				result = append(result, "ra")
			}
		} else {
			for j := 0; j < c.StepsA; j++ {
				ps.Rrx(a)
				result = append(result, "rra")
			}
		}
		if c.Rb {
			for j := 0; j < c.StepsB; j++ {
				ps.Rx(b)
				result = append(result, "rb")
			}
		} else {
			for j := 0; j < c.StepsB; j++ {
				ps.Rrx(b)
				result = append(result, "rrb")
			}
		}
		ps.Px(b, a)
		result = append(result, "pb")

		// fmt.Println("A:", a.GetNumsSlice())
		// fmt.Println("B:", b.GetNumsSlice())
		// fmt.Println()

		// for _, v := range journeyPlanner {
		// 	fmt.Printf("%+v\n", v)
		// }
	}

	return result
}

func sortToA(a, b *ps.Stack) []string {
	result := []string{}
	for {
		A := a.GetNumsSlice()
		B := b.GetNumsSlice()
		journeyPlanner := make([]ps.PushInfo, len(B))

		if len(B) == 0 {
			break
		}

		cheapest := 0
		for i, v := range B {
			var cost int
			var ra, rb bool
			var stepsA, stepsB, jointSteps int

			if i > len(B)/2 {
				cost = len(B) - i
			} else {
				cost = i
				rb = true
			}
			stepsB = cost

			foundOneGreater := false
			targetIndex, targetValue := ps.MaxInt(A)
			for j, w := range A {
				if w > v {
					foundOneGreater = true
					if w < targetValue {
						targetValue = w
						targetIndex = j
					}
				}
			}
			if !foundOneGreater {
				targetIndex, targetValue = ps.MinInt(A)
			}

			if targetIndex > len(A)/2 {
				cost += len(A) - targetIndex
				stepsA = len(A) - targetIndex
			} else {
				cost += targetIndex
				stepsA = targetIndex
				ra = true
			}

			jointSteps = min(stepsA, stepsB)

			// // The following doesn't work here, although it does in
			// // sortToB. I'm not sure why. I guessed it was to do with
			// // a case where the initial pushes to B are not sorted,
			// // but that would suggest the problem would happen in
			// // sortToB too.

			// // Optimization: Account for case where ra XOR rb, but
			// // either of the stacks would be rotated len(X)/2 times.
			// // In that case, it doesn't matter which way we rotate the
			// // stack that needs rotating len(X)/2 times, so we can
			// // set its direction of rotation to the same as the other
			// // to take advantage of any possible joint steps.
			// if stepsA == len(A)/2 {
			// 	ra = rb
			// }
			// if stepsB == len(B)/2 {
			// 	rb = ra
			// }

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

		// fmt.Println("cheapest:", c.Value)
		// fmt.Println("targetIndex:", c.TargetIndex)
		// fmt.Println("targetValue:", c.TargetValue)
		// fmt.Println("cost:", c.Cost)
		// fmt.Println("ra:", c.Ra)
		// fmt.Println("rb:", c.Rb)
		// fmt.Println("stepsA:", c.StepsA)
		// fmt.Println("stepsB:", c.StepsB)
		// fmt.Println("jointSteps:", c.JointSteps)
		// fmt.Println()

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
				result = append(result, "ra")
			}
		} else {
			for j := 0; j < c.StepsA; j++ {
				ps.Rrx(a)
				result = append(result, "rra")
			}
		}
		if c.Rb {
			for j := 0; j < c.StepsB; j++ {
				ps.Rx(b)
				result = append(result, "rb")
			}
		} else {
			for j := 0; j < c.StepsB; j++ {
				ps.Rrx(b)
				result = append(result, "rrb")
			}
		}
		ps.Px(a, b)
		result = append(result, "pa")

		// for _, v := range journeyPlanner {
		// 	fmt.Printf("%+v\n", v)
		// }
	}

	return result
}
