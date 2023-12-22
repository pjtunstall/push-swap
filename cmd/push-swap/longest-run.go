package main

import "push-swap/ps"

// For use with hundredRun in general.go, an implementation of
// Dan Sylvain's idea of leaving the longest increasing run
// on stack A, and pushing the rest to B.
func longestRun(nums []int) (int, int, int) {
	length := 1
	startIndex := 0
	startValue := nums[0]
	maxI, _ := ps.MaxInt(nums)
	minI, _ := ps.MinInt(nums)

	for i, v := range nums {
		count := 1
		for j := i; ; j++ {
			if j+1 < len(nums) && ((nums[j+1] == nums[j]+1) || (j == maxI && j+1 == minI)) {
				count++
				continue
			}
			if j+1 == len(nums) && ((nums[0] == nums[j]+1) || (j == maxI && minI == 0)) {
				count++
				j = -1
				continue
			}
			break
		}
		if count > length {
			length = count
			startIndex = i
			startValue = v
		}
	}
	return startIndex, startValue, length
}
