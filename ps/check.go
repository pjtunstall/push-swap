package ps

func Check(a, b Stack) bool {
	if len(b.Nums) != 0 {
		return false
	}
	for i := range a.Nums {
		if i == len(a.Nums)-1 {
			break
		}
		if a.Nums[i] > a.Nums[i+1] {
			return false
		}
	}
	return true
}
