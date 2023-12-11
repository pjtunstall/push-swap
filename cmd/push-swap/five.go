package main

import (
	"push-swap/ps"
)

func five(a, b *ps.Stack) []string {
	var result []string
	nums := a.GetNumsSlice()
	left := make([]int, 3, 5)
	var maxB int
	var minB int
	var leftRot bool
	var combineRotation bool

	swapRotatable, swappable := swapRot(*a, *b)
	if swappable {
		ps.Sx(a)
		return []string{"sa"}
	}
	if swapRotatable {
		ps.Sx(a)
		rots := justRotate(*a)
		ps.Run(a, b, rots)
		return append([]string{"sa"}, rots...)
	}

	rotSwapScript, rotSwappable := rotSwap(*a, *b)
	if rotSwappable {
		ps.Run(a, b, rotSwapScript)
		return rotSwapScript
	}

	result = []string{"pb", "pb"} // Push top two to B.
	copy(left, nums[2:])          // The three that stay in A.
	_, leftRot = three(left)      // `stayersRot` is true if no swap is needed to sort them.
	maxB = max(nums[0], nums[1])
	minB = min(nums[0], nums[1])

	if leftRot {
		if nums[0] == maxB {
			combineRotation = true
		}
	} else {
		left[0], left[1] = left[1], left[0]
		if nums[0] == maxB {
			result = append(result, "ss")
		} else {
			result = append(result, "sa")
		}
	}

	// Consider maxB at top of B now, unless combineRotation is true.
	switch fitTheFourth(maxB, left) {
	case 0:
		if combineRotation {
			result = append(result, "sb")
		}
	case 1:
		if combineRotation {
			result = append(result, "rr")
		} else {
			result = append(result, "ra")
		}
		left = append(left[1:], left[0])
	case 2:
		if combineRotation {
			result = append(result, "rrr")
		} else {
			result = append(result, "rra")
		}
		left = append(left[len(left)-1:], left[:len(left)-1]...)
	}
	result = append(result, "pa")
	left = append([]int{maxB}, left...)

	switch fitTheFifth(minB, left) {
	case 1:
		result = append(result, "ra")
		left = append(left[1:], left[0])
	case 2:
		result = append(result, "ra", "ra")
		left = append(left[2:], left[:2]...)
	case 3:
		result = append(result, "rra")
		left = append(left[len(left)-1:], left[:len(left)-1]...)
	}
	result = append(result, "pa")
	left = append([]int{minB}, left...)

	copy(a.Nums, left)
	a.Top = 0
	rots := justRotate(*a)
	result = append(result, justRotate(*a)...)
	ps.Run(a, b, rots)

	return result
}

func fitTheFourth(x int, left []int) int {
	var position int
	iMax, maxStayer := ps.MaxInt(left)
	iMin, minStayer := ps.MinInt(left)

	if x < minStayer {
		position = iMin
	} else if x > maxStayer {
		position = iMax + 1
	} else {
		if x > left[0] && x < left[1] {
			position = 1
		} else if x > left[1] && x < left[2] {
			position = 2
		} else if x < left[0] && x > left[2] {
			position = 0
		}
	}
	return position
}

func fitTheFifth(x int, left []int) int {
	var position int
	iMin, minStayer := ps.MinInt(left)

	if x < minStayer {
		position = iMin
	} else {
		if x > left[0] && x < left[1] {
			position = 1
		} else if x > left[1] && x < left[2] {
			position = 2
		} else if x > left[2] && x < left[3] {
			position = 3
		} else {
			position = 0
		}
	}

	return position
}

// Check if a swap, then rotations are enough to sort the stack.
func swapRot(a, b ps.Stack) (bool, bool) {
	ps.Sx(&a)
	rotatable, sorted := ps.Check(a, b)
	ps.Sx(&a)
	return rotatable, sorted
}

// Check if rotations, then a swap (then possibly more rotations)
// are enough to sort the stack.
func rotSwap(a, b ps.Stack) ([]string, bool) {
	var result []string
	var rotSwappable bool

	for h := 1; h <= len(a.Nums)/2; h++ {
		for i := 1; i <= h; i++ {
			ps.Rx(&a)
		}
		ps.Sx(&a)
		rotatable, sorted := ps.Check(a, b)
		if sorted {
			rotSwappable = true
			for i := 1; i <= h; i++ {
				result = append(result, "ra")
			}
			result = append(result, "sa")
		} else if rotatable {
			rotSwappable = true
			for i := 1; i <= h; i++ {
				result = append(result, "ra")
			}
			result = append(result, "sa")
			result = append(result, justRotate(a)...)
		}
		ps.Sx(&a)
		for i := 1; i <= h; i++ {
			ps.Rrx(&a)
		}
		if rotSwappable {
			break
		}

		for i := 1; i <= h; i++ {
			ps.Rrx(&a)
		}
		ps.Sx(&a)
		rotatable, sorted = ps.Check(a, b)
		if sorted {
			rotSwappable = true
			for i := 1; i <= h; i++ {
				result = append(result, "rra")
			}
			result = append(result, "sa")
		} else if rotatable {
			rotSwappable = true
			for i := 1; i <= h; i++ {
				result = append(result, "rra")
			}
			result = append(result, "sa")
			result = append(result, justRotate(a)...)
		}
		ps.Sx(&a)
		for i := 1; i <= h; i++ {
			ps.Rx(&a)
		}
		if rotSwappable {
			break
		}
	}

	return result, rotSwappable
}
