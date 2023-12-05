package ps

func Rx(x *Stack) {
	if len(x.Nums) < 2 {
		return
	}
	if x.Top == len(x.Nums)-1 {
		x.Top = 0
		return
	}
	x.Top += 1
}

func Rrx(x *Stack) {
	if len(x.Nums) < 2 {
		return
	}
	if x.Top == 0 {
		x.Top = len(x.Nums) - 1
		return
	}
	x.Top -= 1
}

func Rr(a, b *Stack) {
	Rx(a)
	Rx(b)
}

func Rrr(a, b *Stack) {
	Rrx(a)
	Rrx(b)
}

func Sx(x *Stack) {
	if len(x.Nums) < 2 {
		return
	}
	if x.Top == len(x.Nums)-1 {
		x.Nums[x.Top], x.Nums[0] = x.Nums[0], x.Nums[x.Top]
		return
	}
	x.Nums[x.Top], x.Nums[x.Top+1] = x.Nums[x.Top+1], x.Nums[x.Top]
}

func Ss(a, b *Stack) {
	Sx(a)
	Sx(b)
}

func Px(x, y *Stack) {
	if len(y.Nums) == 0 {
		return
	}
	if len(x.Nums) == 0 {
		x.Nums = []int{y.Nums[y.Top]}
		x.Top = 0
	} else {
		head := append([]int{y.Nums[y.Top]}, x.Nums[x.Top:]...)
		x.Nums = append(head, x.Nums[:x.Top]...)
	}

	if len(y.Nums) == 1 {
		y.Nums = []int{}
		y.Top = -1
		return
	}
	if y.Top == 0 {
		y.Nums = y.Nums[1:]
		return
	}
	if y.Top == len(y.Nums)-1 {
		y.Nums = y.Nums[:y.Top]
		y.Top = 0
		return
	}
	y.Nums = append(y.Nums[:y.Top], y.Nums[y.Top+1:]...)
}
