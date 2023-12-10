package ps

func Check(a, b Stack) (rotatable, sorted bool) {
	if len(b.Nums) != 0 {
		return false, false
	}

	iMax, _ := MaxInt(a.Nums)
	iMin, _ := MinInt(a.Nums)

	for i := range a.Nums {
		j := a.Top + i
		this := j % len(a.Nums)
		next := (j + 1) % len(a.Nums)
		if a.Nums[this] > a.Nums[next] && this != iMax {
			return false, false
		}
	}
	rotatable = true

	if iMin == a.Top {
		sorted = true
	}

	return rotatable, sorted
}
