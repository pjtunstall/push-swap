package main

import (
	"push-swap/ps"
)

func five(a, b *ps.Stack) []string {
	// Superfluous as this check is already done in main,
	// but here for the sake of TestFive.
	rotatable, sorted := ps.Check(*a, *b)
	if rotatable {
		if sorted {
			return []string{}
		} else {
			return justRotate(*a)
		}
	}

	var result []string
	*a, _ = ps.NewStack(rank(a.Nums))
	nums := a.GetNumsSlice()
	original := a.GetNumsString()
	var maxB int
	var minB int
	var leftRot bool
	var combineRotation bool

	result = []string{"pb", "pb"} // Push top two to B.
	ps.Px(b, a)
	ps.Px(b, a)
	_, leftRot = three(a.GetNumsSlice()) // `stayersRot` is true if no swap is needed to sort them.
	maxB = max(nums[0], nums[1])
	minB = min(nums[0], nums[1])
	if leftRot {
		if nums[0] == maxB {
			combineRotation = true
		}
	} else {
		if nums[0] == maxB {
			ps.Ss(a, b)
			result = append(result, "ss")
		} else {
			ps.Sx(a)
			result = append(result, "sa")
		}
	}
	// Consider maxB at top of B now, unless combineRotation is true.
	switch fitTheFourth(maxB, a.GetNumsSlice()) {
	case 0:
		if combineRotation {
			result = append(result, "sb")
			ps.Sx(b)
		}
	case 1:
		if combineRotation {
			result = append(result, "rr")
			ps.Rr(a, b)
		} else {
			result = append(result, "ra")
			ps.Rx(a)
		}
	case 2:
		if combineRotation {
			result = append(result, "rrr")
			ps.Rrr(a, b)
		} else {
			result = append(result, "rra")
			ps.Rrx(a)
		}
	}
	if len(result) > 0 && result[len(result)-1] == "pb" {
		result = result[:len(result)-1]
	} else {
		result = append(result, "pa")
	}
	ps.Px(a, b)
	switch fitTheFifth(minB, a.GetNumsSlice()) {
	case 1:
		result = append(result, "ra")
		ps.Rx(a)
	case 2:
		result = append(result, "ra", "ra")
		ps.Rx(a)
		ps.Rx(a)
	case 3:
		result = append(result, "rra")
		ps.Rrx(a)
	}
	if len(result) > 0 && result[len(result)-1] == "pb" {
		result = result[:len(result)-1]
	} else {
		result = append(result, "pa")
		ps.Px(a, b)
	}
	rots := justRotate(*a)
	result = append(result, justRotate(*a)...)
	ps.Run(a, b, rots)

	alt, found := bfs(original, len(result))
	if found && len(alt) < len(result) {
		return alt
	}

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
		if position > 2 {
			position = 0
		}
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
