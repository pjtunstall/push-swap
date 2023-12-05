package ps

func AreThereDuplicates(a Stack) bool {
	for i := range a.Nums {
		for j := range a.Nums {
			if i == j {
				continue
			}
			if a.Nums[i] == a.Nums[j] {
				return true
			}
		}
	}
	return false
}
