package main

import (
	"push-swap/ps"
)

func general(a, b *ps.Stack) []string {
	c, _ := ps.NewStack(a.GetNumsString())
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

	ps.Px(b, a)
	ps.Px(b, a)
	result = append(result, "pb", "pb")
	nums := b.GetNumsSlice()
	if nums[0] < nums[1] {
		ps.Sx(b)
		result = append(result, "sb")
	}
	// result = append(result, sortToB(a, b)...)
	result = append(result, merge(a, b, 3, true)...)
	_, rotatable := three(a.Nums)
	if !rotatable {
		ps.Sx(a)
		result = append(result, "sa")
	}
	// result = append(result, sortToA(a, b)...)
	result = append(result, merge(b, a, 0, false)...)

	result = append(result, justRotate(*a)...)
	ps.Run(a, b, justRotate(*a))

	// An extra check to see if we can sort using a sequence of only
	// swaps and rotations. This uses a BFS, and if too slow for big
	// stacks, hence the simpler checks at the beginning of this
	// function.
	if len(a.Nums) < 9 {
		alt := bfs(c, len(result))
		if len(alt) > 0 && len(alt) < len(result) {
			result = alt
		}
	}

	return result
}

// Sorts as it merges from one stack to another. Here a and
// b are just parameters, with a representing the source stack
// and b the destination stack. The stopAt parameter is the number
// of elements to leave in the source stack. If forward is true,
// the merge is from the real a to the real b, otherwise it is
// from the real b to the real a.
func merge(a, b *ps.Stack, stopAt int, forward bool) []string {
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
