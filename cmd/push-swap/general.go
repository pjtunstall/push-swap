package main

import (
	"push-swap/ps"
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
	// n needs to be at least 8, at least in our implementation, because
	// when n = 6 or 7, the "third" left on stack A when the smallest two
	// thirds are pushed to B only has 2 elements, whereas the algorithm
	// says push the final third, consisting of the largest values, till
	// 3 elements are left. Turk algorithm actually gets better and better
	// than Orion till n == 42. After that, Orion starts to improve relative
	// to Turk, and at n == 93, Orion is winning over Turk at last.
	if len(a.Nums) > 92 {
		return orion(a, b)
	}

	// To get these (and hence all stacks of 6) under 13 instructions.
	// Without this check, the current algorithm takes 13 instructions.
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

func orion(a, b *ps.Stack) []string {
	var result []string
	*a, _ = ps.NewStack(rank(a.Nums))
	n := len(a.Nums)
	third := n / 3
	twoThirds := 2 * third

	switch n % 3 {
	case 1:
		twoThirds++
	case 2:
		third++
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
