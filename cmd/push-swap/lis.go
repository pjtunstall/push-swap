package main

// // For use with hundredLIS in general.go, an implementation of
// // Dan Sylvain's idea of leaving the longest increasing sequence
// // on stack A, and pushing the rest to B, then insertion sorting
// // B onto A with a cost check to see which to push next.
// func longestIncreasingSequence(nums []int) []int {
// 	if len(nums) == 0 {
// 		return []int{}
// 	}
// 	result := make([]int, 1, len(nums))
// 	for i := 0; i < len(nums); i++ {
// 		temp := make([]int, 1, len(nums))
// 		temp[0] = nums[i]
// 		for j := i; j < len(nums); j++ {
// 			if nums[j] > temp[len(temp)-1] {
// 				temp = append(temp, nums[j])
// 			}
// 		}
// 		if len(temp) > len(result) {
// 			result = make([]int, len(temp))
// 			copy(result, temp)
// 		}
// 	}
// 	return result
// }
