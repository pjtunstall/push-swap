package main

// // For use with hundredRun in general.go, an implementation of
// // Dan Sylvain's idea of leaving the longest increasing run
// // on stack A, and pushing the rest to B.
// func longestRun(nums []int) (int, int, int) {
// 	length := 1
// 	startIndex := 0
// 	startValue := nums[0]
// 	for i, v := range nums {
// 		count := 1
// 		for j := i; j+1 < len(nums) && nums[j+1] > nums[j]; j++ {
// 			count++
// 		}
// 		if count > length {
// 			length = count
// 			startIndex = i
// 			startValue = v
// 		}
// 	}
// 	return startIndex, startValue, length
// }
