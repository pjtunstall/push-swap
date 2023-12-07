package main

import (
	"push-swap/ps"
)

func justRotate(a *ps.Stack) []string {
	var result []string
	var rx string
	var numberOfRotataionsNeeded int
	iMin, _ := ps.MinInt(a.Nums)
	up := len(a.Nums) / 2
	if iMin <= up {
		rx = "ra"
		numberOfRotataionsNeeded = iMin
	} else {
		rx = "rra"
		numberOfRotataionsNeeded = len(a.Nums) - iMin
	}

	for i := 0; i < numberOfRotataionsNeeded; i++ {
		result = append(result, rx)
		if rx == "ra" {
			a.Nums = append(a.Nums[1:], a.Nums[0])
		} else {
			a.Nums = append(a.Nums[len(a.Nums)-1:], a.Nums[:len(a.Nums)-1]...)
		}
	}

	return result
}
