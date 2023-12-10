package main

import (
	"push-swap/ps"
)

func justRotate(a ps.Stack) []string {
	nums := a.GetNumsSlice()
	var result []string
	var rx string
	distanceFromTop, _ := ps.MinInt(nums)
	midway := len(nums) / 2

	if distanceFromTop <= midway {
		rx = "ra"
	} else {
		rx = "rra"
		distanceFromTop = len(nums) - distanceFromTop
	}

	for i := 0; i < distanceFromTop; i++ {
		result = append(result, rx)
	}

	return result
}
