package main

import (
	"push-swap/ps"
)

func four(a, b ps.Stack) []string {
	result := []string{"pb"}        // Push top two to B.
	stayers := a.Nums[1:]           // The three that stay in A.
	_, stayersRot := three(stayers) // `stayersRot` is true if no swap is needed to sort A.

	if !stayersRot {
		stayers[0], stayers[1] = stayers[1], stayers[0]
		result = append(result, "sa")
	}

	// Consider maxB at top of B now, unless combineRotation is true.
	switch fitTheFourth(a.Nums[0], stayers) {
	case 1:
		result = append(result, "ra")
		stayers = append(stayers[1:], stayers[0])
	case 2:
		result = append(result, "rra")
		stayers = append(stayers[len(stayers)-1:], stayers[:len(stayers)-1]...)
	}
	result = append(result, "pa")
	a.Nums = append(a.Nums[:1], stayers...)
	result = append(result, justRotate(a)...)

	return result
}
