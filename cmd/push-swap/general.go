package main

import (
	"fmt"
	"push-swap/ps"
)

func general(a, b *ps.Stack) []string {
	var result []string

	// Check if a swap, or a swap then rotations, are enough to sort the stack.
	A, _ := ps.NewStack(a.GetNumsString())
	B, _ := ps.NewStack("")
	swapRotatable, swappable := swapRot(A, B)

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
	A, _ = ps.NewStack(a.GetNumsString())
	B, _ = ps.NewStack("")
	rotSwapScript, rotSwappable := rotSwap(A, B)
	if rotSwappable {
		fmt.Println("rotSwappable")
		ps.Run(a, b, rotSwapScript)
		return rotSwapScript
	}

	return result
}
