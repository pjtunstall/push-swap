package main

import (
	"push-swap/ps"
	"sort"
	"strconv"
	"strings"
)

func rank(nums []int) string {
	rankMap := make(map[int]int)
	strs := make([]string, len(nums))
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	sort.Ints(numsCopy)
	for i, v := range numsCopy {
		rankMap[v] = i + 1
	}

	for i, v := range nums {
		strs[i] = strconv.Itoa(rankMap[v])
	}
	s := strings.Join(strs, " ")

	return s
}

func four(a, b ps.Stack) []string {
	s := rank(a.Nums)

	switch s {
	case "1 2 3 4":
		return []string{}
	case "1 2 4 3":
		return []string{"pb", "sa", "ra", "pa"}
	case "1 3 2 4":
		return []string{"ra", "sa", "rra"} // or alternatively ...
		// return []string{"pb", "sa", "pa"}
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
