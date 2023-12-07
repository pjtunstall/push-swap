package ps

func MaxInt(nums []int) (int, int) {
	max := nums[0]
	index := 0
	for i, num := range nums {
		if num > max {
			max = num
			index = i
		}
	}
	return index, max
}

func MinInt(nums []int) (int, int) {
	min := nums[0]
	index := 0
	for i, num := range nums {
		if num < min {
			min = num
			index = i
		}
	}
	return index, min
}
