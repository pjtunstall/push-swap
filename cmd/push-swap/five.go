package main

import (
	"push-swap/ps"
)

func five(a, b *ps.Stack) []string {
	var result []string
	var stayers []int
	var maxB int
	var minB int
	var stayersRot bool
	var combineRotation bool

	// c := ps.Stack{Nums: a.Nums[:1], Top: 0}
	// if ps.Check(c, *b, false) {
	// 	stayers = a.Nums[1:]
	// 	minB = a.Nums[0]
	// 	result = []string{"pb"}
	// 	goto skip
	// }

	result = []string{"pb", "pb"}  // Push top two to B.
	stayers = a.Nums[2:]           // The three that stay in A.
	_, stayersRot = three(stayers) // `stayersRot` is true if no swap is needed to sort A.
	maxB = max(a.Nums[0], a.Nums[1])
	minB = min(a.Nums[0], a.Nums[1])

	if stayersRot {
		if a.Nums[0] == maxB {
			combineRotation = true
		}
	} else {
		stayers[0], stayers[1] = stayers[1], stayers[0]
		if a.Nums[0] == maxB {
			result = append(result, "ss")
		} else {
			result = append(result, "sa")
		}
	}

	// Consider maxB at top of B now, unless combineRotation is true.
	switch fitTheFourth(maxB, stayers) {
	case 1:
		if combineRotation {
			result = append(result, "rr")
		} else {
			result = append(result, "ra")
		}
		stayers = append(stayers[1:], stayers[0])
	case 2:
		if combineRotation {
			result = append(result, "rrr")
		} else {
			result = append(result, "rra")
		}
		stayers = append(stayers[len(stayers)-1:], stayers[:len(stayers)-1]...)
	}
	result = append(result, "pa")
	stayers = append([]int{maxB}, stayers...)

	// skip:
	switch fitTheFifth(minB, stayers) {
	case 1:
		result = append(result, "ra")
	case 2:
		result = append(result, "ra", "ra")
	case 3:
		result = append(result, "rra", "rra")
	}
	result = append(result, "pa")

	return result
}

func fitTheFourth(x int, stayers []int) int {
	var position int
	iMax, maxStayer := ps.MaxInt(stayers)
	iMin, minStayer := ps.MinInt(stayers)

	if x < minStayer {
		position = iMin
	} else if x > maxStayer {
		position = iMax + 1
	} else {
		if x > stayers[0] && x < stayers[1] {
			position = 1
		} else if x > stayers[1] && x < stayers[2] {
			position = 2
		} else if x < stayers[0] {
			position = 0
		} else {
			position = 3
		}
	}
	return position
}

func fitTheFifth(x int, stayers []int) int {
	var position int
	iMin, minStayer := ps.MinInt(stayers)

	if x < minStayer {
		position = iMin
	} else {
		if x > stayers[0] && x < stayers[1] {
			position = 1
		} else if x > stayers[1] && x < stayers[2] {
			position = 2
		} else if x > stayers[2] && x < stayers[3] {
			position = 3
		} else {
			position = 0
		}
	}

	return position
}
