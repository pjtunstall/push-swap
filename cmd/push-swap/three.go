package main

func three(nums []int) string {
	if nums[2] < nums[0] && nums[1] < nums[2] {
		return "ra\n"
	}
	if nums[0] < nums[1] && nums[2] < nums[0] {
		return "rra\n"
	}
	if nums[1] < nums[0] && nums[0] < nums[2] {
		return "sa\n"
	}
	if nums[0] < nums[1] && nums[2] < nums[1] {
		return "sa\nra\n"
	}
	if nums[1] < nums[0] && nums[2] < nums[1] {
		return "sa\nrra\n"
	}
	return ""
}
