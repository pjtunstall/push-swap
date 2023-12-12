package main

import (
	"push-swap/ps"
)

func general(a, b *ps.Stack) []string {
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

	ps.Px(b, a)
	ps.Px(b, a)
	result = append(result, "pb", "pb")

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

			targetValue := 0
			var targetIndex int
			for j, w := range B {
				if w < v && w > targetValue {
					targetValue = w
					targetIndex = j
				}
			}
			if targetValue == 0 {
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

			if (ra && rb) || (!ra && !rb) {
				jointSteps = min(stepsA, stepsB)
				cost -= jointSteps
				stepsA -= jointSteps
				stepsB -= jointSteps
			}
			// TODO: Account for case where ra XOR rb, but jointSteps would be len(B)/2,
			// including where len(B) == 2.

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
			if cost < cheapest {
				cheapest = cost
			}
		}

		for _, v := range journeyPlanner {
			if v.Cost != journeyPlanner[cheapest].Cost {
				continue
			}
			if v.Ra && v.Rb {
				for j := 0; j < v.JointSteps; j++ {
					ps.Rr(b, a)
					result = append(result, "rr")
				}
			} else if !v.Ra && !v.Rb {
				for j := 0; j < v.JointSteps; j++ {
					ps.Rr(b, a)
					result = append(result, "rrr")
				}
			}
			if v.Ra {
				for j := 0; j < v.StepsA; j++ {
					ps.Rx(a)
					result = append(result, "ra")
				}
			}
			if v.Rb {
				for j := 0; j < v.StepsB; j++ {
					ps.Rx(b)
					result = append(result, "rb")
				}
			}
			ps.Px(b, a)
			result = append(result, "pb")
			break
		}

		// fmt.Println("A:", a.GetNumsSlice())
		// fmt.Println("B:", b.GetNumsSlice())
		// for _, v := range journeyPlanner {
		// 	fmt.Printf("%+v\n", v)
		// }
	}

	return result
}
