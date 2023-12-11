package main

import "push-swap/ps"

func general(a, b *ps.Stack) []string {
	var result []string

	// Check if a swap is enough to sort the stack.
	swapRotatable, swappable := swapRot(*a, *b)
	if swappable {
		ps.Sx(a)
		result = append(result, "sa")
		return result
	}

	// Check if a swap, then rotations are enough to sort the stack.
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

	return result
}
