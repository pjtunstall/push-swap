package main

import (
	"push-swap/ps"
	"sort"
	"strconv"
	"strings"
)

func four(a, b ps.Stack) []string {
	nums := make([]int, len(a.Nums))
	strs := make([]string, len(nums))
	copy(nums, a.Nums)
	sort.Ints(nums)

	rank := make(map[int]int)
	for i, v := range nums {
		rank[v] = i + 1
	}

	for i, v := range a.Nums {
		strs[i] = strconv.Itoa(rank[v])
	}
	s := strings.Join(strs, " ")

	switch s {
	case "1 2 3 4":
		return []string{}
	case "1 2 4 3":
		return []string{"pb", "sa", "ra", "pa"}
	case "1 3 2 4":
		return []string{"pb", "sa", "pa"}
	case "1 3 4 2":
		return []string{"rra", "sa"}
	case "1 4 2 3":
		return []string{"sa", "ra"}
	case "1 4 3 2":
		return []string{"pb", "sa", "rra", "pa"}

	case "2 1 3 4":
		return []string{"sa"}
	case "2 1 4 3":
		return []string{"pb", "pb", "ss", "pa", "pa"}
	case "2 3 1 4":
		return []string{"ra", "ra", "sa", "ra"}
	case "2 3 4 1":
		return []string{"rra"}
	case "2 4 1 3":
		return []string{"sa", "ra", "sa"}
	case "2 4 3 1":
		return []string{"sa", "rra", "sa", "ra"}

	case "3 1 2 4":
		return []string{"sa", "ra", "sa", "rra"}
	case "3 1 4 2":
		return []string{"sa", "rra", "sa"}
	case "3 2 1 4":
		return []string{"sa", "ra", "ra", "sa", "ra"}
	case "3 2 4 1":
		return []string{"sa", "rra"}
	case "3 4 1 2":
		return []string{"ra", "ra"}
	case "3 4 2 1":
		return []string{"ra", "ra", "sa"}

	case "4 1 2 3":
		return []string{"ra"}
	case "4 1 3 2":
		return []string{"ra", "ra", "sa", "rra"}
	case "4 2 1 3":
		return []string{"ra", "sa"}
	case "4 2 3 1":
		return []string{"rra", "sa", "ra"}
	case "4 3 1 2":
		return []string{"sa", "ra", "ra"}
	case "4 3 2 1":
		return []string{"sa", "ra", "ra", "sa"}
	}

	return []string{}
}

// func four(a, b ps.Stack) []string {
// 	result := []string{"pb"}        // Push top two to B.
// 	stayers := a.Nums[1:]           // The three that stay in A.
// 	_, stayersRot := three(stayers) // `stayersRot` is true if no swap is needed to sort A.

// 	if !stayersRot {
// 		stayers[0], stayers[1] = stayers[1], stayers[0]
// 		result = append(result, "sa")
// 	}

// 	// Consider maxB at top of B now, unless combineRotation is true.
// 	switch fitTheFourth(a.Nums[0], stayers) {
// 	case 1:
// 		result = append(result, "ra")
// 		stayers = append(stayers[1:], stayers[0])
// 	case 2:
// 		result = append(result, "rra")
// 		stayers = append(stayers[len(stayers)-1:], stayers[:len(stayers)-1]...)
// 	}
// 	result = append(result, "pa")
// 	a.Nums = append(a.Nums[:1], stayers...)
// 	result = append(result, justRotate(a)...)

// 	return result
// }
