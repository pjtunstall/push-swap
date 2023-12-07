package main

func three(nums []int) ([]string, bool) {
	if nums[2] < nums[0] && nums[1] < nums[2] {
		return []string{"ra"}, true
	}
	if nums[0] < nums[1] && nums[2] < nums[0] {
		return []string{"rra"}, true
	}
	if nums[1] < nums[0] && nums[0] < nums[2] {
		return []string{"sa"}, false
	}
	if nums[0] < nums[1] && nums[2] < nums[1] {
		return []string{"sa", "ra"}, false
	}
	if nums[1] < nums[0] && nums[2] < nums[1] {
		return []string{"sa", "rra"}, false
	}
	return []string{}, true
}
