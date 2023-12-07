package ps

func f(i, iMax, iMin int, full bool) bool {
	// Something is fishy here. Look at this tomorrow.
	if full || (i == iMax && i+1 == iMin) {
		return true
	}
	return false
}

func Check(a, b Stack, full bool) bool {
	if len(b.Nums) != 0 {
		return false
	}

	iMax, _ := MaxInt(a.Nums)
	iMin, _ := MinInt(a.Nums)

	for i := range a.Nums {
		if i == len(a.Nums)-1 {
			break
		}
		// fmt.Println(a.Nums[i], a.Nums[i+1])
		// fmt.Println(f(i, iMax, iMin, full))
		// This condition is not right.
		if a.Nums[i] > a.Nums[i+1] && f(i, iMax, iMin, full) {
			return false
		}
	}
	return true
}
