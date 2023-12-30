package main

import (
	"encoding/json"
	"fmt"
	"os"

	"push-swap/ps"
)

func main() {
	count := 0
	results := make(map[string][]string)

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			for k := 1; k <= 5; k++ {
				for l := 1; l <= 5; l++ {
					for m := 1; m <= 5; m++ {
						// for n := 1; n <= 7; n++ {
						// 	for o := 1; o <= 7; o++ {
						// if (i == j || i == k || i == l || i == m || i == n || i == o) || (j == k || j == l || j == m || j == n || j == o) || (k == l || k == m || k == n || k == o) || (l == m || l == n || l == o) || (m == n || m == o) || (n == o) {
						if (i == j || i == k || i == l || i == m) || (j == k || j == l || j == m) || (k == l || k == m) || (l == m) {
							continue
						}
						input := fmt.Sprintf("%d %d %d %d %d", i, j, k, l, m)
						a, _ := ps.NewStack(input)
						b, _ := ps.NewStack("")

						// turk modifies the stacks.
						turk := turk(&a, &b)
						// five := five(&a, &b)

						// bfs doesn't modify the stacks.
						beef, sorted := bfs(input, len(turk))

						if sorted && len(beef) < len(turk) {
							count++
							results[input] = beef
							fmt.Println("************************** bfs is shortest")
							fmt.Println("input: ", input)
							fmt.Println("length of `bfs`: ", len(beef))
							fmt.Println("length of `turk`: ", len(turk))
							fmt.Println("`bfs`: ", beef)
							fmt.Println("`turk`: ", turk)
							fmt.Println()
						}
					}
				}
			}
		}
	}
	// 	}
	// }
	fmt.Println("instances where bfs is strictly better than five: ", count)

	jsonData, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("results.json", jsonData, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func turk(a, b *ps.Stack) []string {
	result := []string{}

	if a.GetNumsString() == "4 3 2 1 6 5" {
		result = []string{"pb", "pb", "ss", "ra", "ra", "sa", "pa", "pa", "rra", "rra"}
		ps.Run(a, b, result)
		return result
	}

	if a.GetNumsString() == "2 6 5 4 3 1" {
		result = []string{"pb", "rra", "pb", "ss", "ra", "ra", "sa", "pa", "pa"}
		ps.Run(a, b, result)
		return result
	}

	if a.GetNumsString() == "3 1 2 6 5 4" {
		result = []string{"ra", "pb", "pb", "sa", "ra", "ra", "sa", "pa", "pa"}
		ps.Run(a, b, result)
		return result
	}

	swapRotatable, swappable := swapRot(*a, *b)

	// If a swap is enough to sort the stack:
	if swappable {
		ps.Sx(a)
		result = append(result, "sa")
		return result
	}

	// If a swap, then rotations are enough to sort the stack:
	if swapRotatable {
		ps.Sx(a)
		rots := justRotate(*a)
		ps.Run(a, b, rots)
		result = append(result, "sa")
		result = append(result, rots...)
		return result
	}

	// Check if rotations, then a swap, then (possibly) more rotations,
	// are enough to sort the stack.
	rotSwapScript, rotSwappable := rotSwap(*a, *b)
	if rotSwappable {
		ps.Run(a, b, rotSwapScript)
		return rotSwapScript
	}

	// AYO's "Turk" algorithm.
	ps.Px(b, a)
	ps.Px(b, a)
	result = append(result, "pb", "pb")

	nums := b.GetNumsSlice()
	if nums[0] < nums[1] {
		ps.Sx(b)
		result = append(result, "sb")
	}

	result = append(result, insert(a, b, 3, true)...)

	_, rotatable := three(a.Nums)
	if !rotatable {
		ps.Sx(a)
		result = append(result, "sa")

	}

	result = append(result, insert(b, a, 0, false)...)

	result = append(result, justRotate(*a)...)
	ps.Run(a, b, justRotate(*a))

	return result
}

func bfs(s string, n int) ([]string, bool) {
	var result []string

	q := [][]string{{"sa"}, {"ra"}, {"rra"}}
	for len(q) > 0 {
		v := q[0]
		if len(q) == 1 {
			q = [][]string{}
		} else {
			q = q[1:]
		}
		if len(v) >= n {
			break
		}
		a, _ := ps.NewStack(s)
		b, _ := ps.NewStack("")
		ps.Run(&a, &b, v)
		_, sorted := ps.Check(a, b)
		if sorted {
			return v, true
		}

		u := make([]string, len(v))
		copy(u, v)
		switch v[len(v)-1] {
		case "sa":
			q = append(q, append(u, "ra"))
			q = append(q, append(u, "rra"))
		case "ra":
			q = append(q, append(u, "sa"))
			q = append(q, append(u, "ra"))
		case "rra":
			q = append(q, append(u, "sa"))
			q = append(q, append(u, "rra"))
		}
	}

	return result, false
}

// func bfs(n int) {
// 	results := make(map[string][]string)

// 	q := [][]string{{"sa"}, {"ra"}, {"rra"}}
// 	for len(q) > 0 {
// 		v := q[0]
// 		if len(q) == 1 {
// 			q = [][]string{}
// 		} else {
// 			q = q[1:]
// 		}
// 		if len(v) > 10 {
// 			break
// 		}
// 		original := seed(5)
// 		a, _ := ps.NewStack(original)
// 		b, _ := ps.NewStack("")
// 		inv := inverse(v)
// 		ps.Run(&a, &b, inv)
// 		nums := make([]int, 5)
// 		copy(nums, a.GetNumsSlice())
// 		numsString := a.GetNumsString()
// 		five := five(&a, &b)
// 		a, _ = ps.NewStack(numsString)
// 		general := general(&a, &b)
// 		_, sorted := ps.Check(a, b)
// 		if !sorted {
// 			fmt.Println("Not sorted:", a.GetNumsString())
// 		} else {
// 			m := min(len(five), len(general))
// 			if len(v) < m {
// 				fmt.Println(numsString)
// 				fmt.Println("************************** bfs is shortest")
// 				fmt.Println("length of `five`: ", len(five))
// 				fmt.Println("length of `general`: ", len(general))
// 				fmt.Println("length of `bfs`: ", len(v))
// 				fmt.Println("`five`: ", five)
// 				fmt.Println("`general`: ", general)
// 				fmt.Println("`bfs`: ", v)
// 				fmt.Println()
// 				existing, ok := results[numsString]
// 				if !ok || len(v) < len(existing) {
// 					results[numsString] = v
// 				}
// 			}
// 		}

// 		u := make([]string, len(v))
// 		copy(u, v)
// 		switch v[len(v)-1] {
// 		case "sa":
// 			q = append(q, append(u, "ra"))
// 			q = append(q, append(u, "rra"))
// 		case "ra":
// 			q = append(q, append(u, "sa"))
// 			q = append(q, append(u, "ra"))
// 		case "rra":
// 			q = append(q, append(u, "sa"))
// 			q = append(q, append(u, "rra"))
// 		}
// 	}

// 	jsonData, err := json.Marshal(results)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	err = os.WriteFile("results.json", jsonData, 0644)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func inverse(instructions []string) []string {
// 	var result []string
// 	for i := len(instructions) - 1; i >= 0; i-- {
// 		switch instructions[i] {
// 		case "sa":
// 			result = append(result, "sa")
// 		case "sb":
// 			result = append(result, "sb")
// 		case "ss":
// 			result = append(result, "ss")
// 		case "pa":
// 			result = append(result, "pb")
// 		case "pb":
// 			result = append(result, "pa")
// 		case "ra":
// 			result = append(result, "rra")
// 		case "rb":
// 			result = append(result, "rrb")
// 		case "rr":
// 			result = append(result, "rrr")
// 		case "rra":
// 			result = append(result, "ra")
// 		case "rrb":
// 			result = append(result, "rb")
// 		case "rrr":
// 			result = append(result, "rr")
// 		}
// 	}
// 	return result
// }

// Check if a swap, then rotations are enough to sort the stack.
func swapRot(a, b ps.Stack) (bool, bool) {
	ps.Sx(&a)
	rotatable, sorted := ps.Check(a, b)
	ps.Sx(&a)
	return rotatable, sorted
}

// Check if rotations, then a swap (then possibly more rotations)
// are enough to sort the stack.
func rotSwap(a, b ps.Stack) ([]string, bool) {
	var result []string
	var rotSwappable bool

	for h := 1; h <= len(a.Nums)/2; h++ {
		for i := 1; i <= h; i++ {
			ps.Rx(&a)
		}
		ps.Sx(&a)
		rotatable, sorted := ps.Check(a, b)
		if sorted {
			rotSwappable = true
			for i := 1; i <= h; i++ {
				result = append(result, "ra")
			}
			result = append(result, "sa")
		} else if rotatable {
			rotSwappable = true
			for i := 1; i <= h; i++ {
				result = append(result, "ra")
			}
			result = append(result, "sa")
			result = append(result, justRotate(a)...)
		}
		ps.Sx(&a)
		for i := 1; i <= h; i++ {
			ps.Rrx(&a)
		}
		if rotSwappable {
			break
		}

		for i := 1; i <= h; i++ {
			ps.Rrx(&a)
		}
		ps.Sx(&a)
		rotatable, sorted = ps.Check(a, b)
		if sorted {
			rotSwappable = true
			for i := 1; i <= h; i++ {
				result = append(result, "rra")
			}
			result = append(result, "sa")
		} else if rotatable {
			rotSwappable = true
			for i := 1; i <= h; i++ {
				result = append(result, "rra")
			}
			result = append(result, "sa")
			result = append(result, justRotate(a)...)
		}
		ps.Sx(&a)
		for i := 1; i <= h; i++ {
			ps.Rx(&a)
		}
		if rotSwappable {
			break
		}
	}

	return result, rotSwappable
}

func justRotate(a ps.Stack) []string {
	nums := a.GetNumsSlice()
	var result []string
	var rx string
	distanceFromTop, _ := ps.MinInt(nums)
	midway := len(nums) / 2

	if distanceFromTop <= midway {
		rx = "ra"
	} else {
		rx = "rra"
		distanceFromTop = len(nums) - distanceFromTop
	}

	for i := 0; i < distanceFromTop; i++ {
		result = append(result, rx)
	}

	return result
}

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

// func rank(nums []int) string {
// 	rankMap := make(map[int]int)
// 	strs := make([]string, len(nums))
// 	numsCopy := make([]int, len(nums))
// 	copy(numsCopy, nums)

// 	sort.Ints(numsCopy)
// 	for i, v := range numsCopy {
// 		rankMap[v] = i + 1
// 	}

// 	for i, v := range nums {
// 		strs[i] = strconv.Itoa(rankMap[v])
// 	}
// 	s := strings.Join(strs, " ")

// 	return s
// }

// func five(a, b *ps.Stack) []string {
// 	var result []string
// 	*a, _ = ps.NewStack(rank(a.Nums))
// 	nums := a.GetNumsSlice()
// 	// numsString := a.GetNumsString()
// 	var maxB int
// 	var minB int
// 	var leftRot bool
// 	var combineRotation bool

// 	// jsonData, err := os.ReadFile("shortcuts-five.json")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return []string{}
// 	// }
// 	// shortcuts := make(map[string][]string)
// 	// err = json.Unmarshal(jsonData, &shortcuts)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return []string{}
// 	// }

// 	// v, ok := shortcuts[numsString]
// 	// if ok {
// 	// 	return v
// 	// }

// 	result = []string{"pb", "pb"} // Push top two to B.
// 	ps.Px(b, a)
// 	ps.Px(b, a)
// 	_, leftRot = three(a.GetNumsSlice()) // `stayersRot` is true if no swap is needed to sort them.
// 	maxB = max(nums[0], nums[1])
// 	minB = min(nums[0], nums[1])

// 	if leftRot {
// 		if nums[0] == maxB {
// 			combineRotation = true
// 		}
// 	} else {
// 		if nums[0] == maxB {
// 			ps.Ss(a, b)
// 			result = append(result, "ss")
// 		} else {
// 			ps.Sx(a)
// 			result = append(result, "sa")
// 		}
// 	}

// 	// Consider maxB at top of B now, unless combineRotation is true.
// 	switch fitTheFourth(maxB, a.GetNumsSlice()) {
// 	case 0:
// 		if combineRotation {
// 			result = append(result, "sb")
// 			ps.Sx(b)
// 		}
// 	case 1:
// 		if combineRotation {
// 			result = append(result, "rr")
// 			ps.Rr(a, b)
// 		} else {
// 			result = append(result, "ra")
// 			ps.Rx(a)
// 		}
// 	case 2:
// 		if combineRotation {
// 			result = append(result, "rrr")
// 			ps.Rrr(a, b)
// 		} else {
// 			result = append(result, "rra")
// 			ps.Rrx(a)
// 		}
// 	}
// 	if len(result) > 0 && result[len(result)-1] == "pb" {
// 		result = result[:len(result)-1]
// 	} else {
// 		result = append(result, "pa")
// 	}
// 	ps.Px(a, b)

// 	switch fitTheFifth(minB, a.GetNumsSlice()) {
// 	case 1:
// 		result = append(result, "ra")
// 		ps.Rx(a)
// 	case 2:
// 		result = append(result, "ra", "ra")
// 		ps.Rx(a)
// 		ps.Rx(a)
// 	case 3:
// 		result = append(result, "rra")
// 		ps.Rrx(a)
// 	}
// 	if len(result) > 0 && result[len(result)-1] == "pb" {
// 		result = result[:len(result)-1]
// 	} else {
// 		result = append(result, "pa")
// 		ps.Px(a, b)
// 	}

// 	rots := justRotate(*a)
// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, rots)

// 	return result
// }

// func fitTheFourth(x int, left []int) int {
// 	var position int
// 	iMax, maxStayer := ps.MaxInt(left)
// 	iMin, minStayer := ps.MinInt(left)

// 	if x < minStayer {
// 		position = iMin
// 	} else if x > maxStayer {
// 		position = iMax + 1
// 		if position > 2 {
// 			position = 0
// 		}
// 	} else {
// 		if x > left[0] && x < left[1] {
// 			position = 1
// 		} else if x > left[1] && x < left[2] {
// 			position = 2
// 		} else if x < left[0] && x > left[2] {
// 			position = 0
// 		}
// 	}
// 	return position
// }

// func fitTheFifth(x int, left []int) int {
// 	var position int
// 	iMin, minStayer := ps.MinInt(left)

// 	if x < minStayer {
// 		position = iMin
// 	} else {
// 		if x > left[0] && x < left[1] {
// 			position = 1
// 		} else if x > left[1] && x < left[2] {
// 			position = 2
// 		} else if x > left[2] && x < left[3] {
// 			position = 3
// 		} else {
// 			position = 0
// 		}
// 	}

// 	return position
// }

// // Works, just not as well as five().
// func fiveExperiment(a, b *ps.Stack) []string {
// 	var result []string
// 	*a, _ = ps.NewStack(rank(a.Nums))
// 	original := a.GetNumsString()
// 	nums := a.GetNumsSlice()

// 	rotatable, sorted := ps.Check(*a, *b)
// 	if sorted {
// 		return result
// 	}
// 	if rotatable {
// 		result = justRotate(*a)
// 		return result
// 	}

// 	for i := 0; i+1 < len(nums); i++ {
// 		if nums[i]+1 != nums[i+1] || (nums[i] == 5 && nums[i+1] != 1) {
// 			ps.Px(b, a)
// 			result = append(result, "pb")
// 			break
// 		}
// 	}

// 	var new []string

// 	s := rank(a.GetNumsSlice())
// 	c, _ := ps.NewStack(s)
// 	d, _ := ps.NewStack("")
// 	switch s {
// 	case "1 3 2 4":
// 		new = []string{"ra", "sa"}
// 	case "1 4 2 3":
// 		new = []string{"sa"}
// 	case "2 3 1 4":
// 		new = []string{"ra", "ra", "sa"}
// 	case "2 3 4 1":
// 		new = []string{}
// 	case "2 4 3 1":
// 		new = []string{"sa", "rra", "sa"}
// 	case "3 1 2 4":
// 		new = []string{"sa", "ra", "sa"}
// 	case "3 2 1 4":
// 		new = []string{"sa", "ra", "ra", "sa"}
// 	case "3 2 4 1":
// 		new = []string{"sa"}
// 	case "3 4 1 2":
// 		new = []string{}
// 	case "4 1 2 3":
// 		new = []string{}
// 	case "4 1 3 2":
// 		new = []string{"ra", "ra", "sa"}
// 	case "4 2 1 3":
// 		new = []string{"ra", "sa"}
// 	case "4 2 3 1":
// 		new = []string{"rra", "sa"}
// 	case "4 3 1 2":
// 		new = []string{"sa"}
// 	default:
// 		new = four(c, d)
// 	}

// 	result = append(result, new...)
// 	ps.Run(a, b, new)

// 	if b.Nums[b.Top] == 5 {
// 		switch a.Nums[a.Top] {
// 		case 2:
// 			result = append(result, "rra")
// 			ps.Rrx(a)
// 		case 3:
// 			result = append(result, "ra", "ra")
// 			ps.Rx(a)
// 			ps.Rx(a)
// 		case 4:
// 			result = append(result, "ra")
// 			ps.Rx(a)
// 		}
// 	} else {
// 		// No, we want to find the index of the number that is one more than the top of B.
// 		nums = a.GetNumsSlice()
// 		I := 0
// 		for i := range nums {
// 			if nums[i] == b.Nums[b.Top]+1 {
// 				I = i
// 			}
// 		}
// 		switch I {
// 		case 3:
// 			result = append(result, "rra")
// 			ps.Rrx(a)
// 		case 2:
// 			result = append(result, "ra", "ra")
// 			ps.Rx(a)
// 			ps.Rx(a)
// 		case 1:
// 			result = append(result, "ra")
// 			ps.Rx(a)
// 		}
// 	}

// 	ps.Px(a, b)
// 	result = append(result, "pa")

// 	_, sorted = ps.Check(*a, *b)
// 	if !sorted {
// 		result = append(result, justRotate(*a)...)
// 	}

// 	bfs, found := bfs(original, 8)
// 	if found && len(bfs) < len(result) {

// 		return bfs
// 	}

// 	return result
// }

// // This was in five() to check for a shorter push-free solution
// // before BFS was done at runtime:
// jsonData, err := os.ReadFile("shortcuts-five.json")
// if err != nil {
// 	fmt.Println(err)
// 	return []string{}
// }
// shortcuts := make(map[string][]string)
// err = json.Unmarshal(jsonData, &shortcuts)
// if err != nil {
// 	fmt.Println(err)
// 	return []string{}
// }
// v, ok := shortcuts[original]
// if ok {
// 	return v
// }

// *

// from push-swap
// package main

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

// // After Dan Sylvain: leave the longest increasing sequence on stack A,
// // and push the rest to B. Then insertion sort back with a cost check
// // to see which to push next.
// func hundredLIS(a, b *ps.Stack) []string {
// 	var result []string
// 	*a, _ = ps.NewStack(rank(a.Nums))
// 	LIS := longestIncreasingSequence(a.Nums)
// 	l := len(LIS)

// 	// Push the smallest half to the bottom of stack B and the
// 	// biggest half to the top, leaving any numbers in the longest
// 	// increasing sequence on stack A.
// 	for len(a.Nums) > l {
// 		found := false
// 		for i := range LIS {
// 			if a.Nums[a.Top] == LIS[i] {
// 				found = true
// 			}
// 		}
// 		if found {
// 			ps.Rx(a)
// 			result = append(result, "ra")
// 			continue
// 		}
// 		ps.Px(b, a)
// 		result = append(result, "pb")
// 		if b.Nums[b.Top] < 50 {
// 			ps.Rx(b)
// 			result = append(result, "rb")
// 		}
// 	}

// 	// Sort while inserting from stack B to stack A.
// 	result = append(result, insert(b, a, 0, false)...)

// 	// Rotate stack A into sorted position.
// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, justRotate(*a))

// 	return result
// }

// // JC's algorithm.
// func jc(a, b *ps.Stack) []string {
// 	var result []string

// 	for len(a.Nums) > 1 {
// 		ps.Px(b, a)
// 		result = append(result, "pb")
// 	}

// 	result = append(result, insert(b, a, 0, false)...)

// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, justRotate(*a))

// 	return result
// }

// package main

// import "push-swap/ps"

// // For use with hundredRun in general.go, which leaves the longest
// // increasing run on stack A, and pushes the rest to B.
// func longestRun(nums []int) (int, int, int) {
// 	length := 1
// 	startIndex := 0
// 	startValue := nums[0]
// 	maxI, _ := ps.MaxInt(nums)
// 	minI, _ := ps.MinInt(nums)

// 	for i, v := range nums {
// 		count := 1
// 		for j := i; ; j++ {
// 			if j+1 < len(nums) && ((nums[j+1] == nums[j]+1) || (j == maxI && j+1 == minI)) {
// 				count++
// 				continue
// 			}
// 			if j+1 == len(nums) && ((nums[0] == nums[j]+1) || (j == maxI && minI == 0)) {
// 				count++
// 				j = -1
// 				continue
// 			}
// 			break
// 		}
// 		if count > length {
// 			length = count
// 			startIndex = i
// 			startValue = v
// 		}
// 	}
// 	return startIndex, startValue, length
// }

// package main

// import "testing"

// func TestLongestRun(t *testing.T) {
// 	a := []int{1, 2, 4, 3, 5}
// 	startIndex, startValue, length := longestRun(a)
// 	if startIndex != 4 || startValue != 5 || length != 3 {
// 		t.Error("Wrong result")
// 	}
// }

// // Leave longest increasing run on stack A and push the smallest
// // half to the bottom of stack B and the biggest half to the top.
// func hundredRun(a, b *ps.Stack) []string {
// 	var result []string
// 	*a, _ = ps.NewStack(rank(a.Nums))
// 	A := a.GetNumsSlice()
// 	startIndex, startValue, length := longestRun(A)

// 	// Deal with the case where the longest run overlaps the top.
// 	diff := len(A) - startIndex
// 	if diff < length {
// 		for i := 0; i < length-diff; i++ {
// 			ps.Rx(a)
// 			result = append(result, "ra")
// 		}
// 	}

// 	// Push the smallest half to the bottom of stack B and the
// 	// biggest half to the top, leaving the longest increasing
// 	// run on stack A, leaving the longest increasing run on
// 	// stack A.
// 	for len(a.Nums) > length {
// 		if a.Nums[a.Top] == startValue {
// 			for i := 0; i < length; i++ {
// 				ps.Rx(a)
// 				result = append(result, "ra")
// 			}
// 		}
// 		ps.Px(b, a)
// 		result = append(result, "pb")
// 		if b.Nums[b.Top] < 50 {
// 			ps.Rx(b)
// 			result = append(result, "rb")
// 		}
// 	}

// 	// Sort while inserting from stack B to stack A.
// 	result = append(result, insert(b, a, 0, false)...)

// 	// Rotate stack A into sorted position.
// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, justRotate(*a))

// 	return result
// }

// These three stacks require more than 12 moves to sort by Fred's
// algorithm. The last two also require more than 12 moves to sort
// by AYO's algorithm. In all cases, they took 13. They're the only
// stacks of 6 that take more than 12.

// if a.GetNumsString() == "4 3 2 1 6 5" {
// 	// 	result = []string{"pb", "pb", "ss", "ra", "ra", "sa", "pa", "pa", "rra", "rra"}
// 	// 	ps.Run(a, b, result)
// 	// 	return result
// 	// }

// 	if a.GetNumsString() == "2 6 5 4 3 1" {
// 		result = []string{"pb", "rra", "pb", "ss", "ra", "ra", "sa", "pa", "pa"}
// 		ps.Run(a, b, result)
// 		return result
// 	}

// 	if a.GetNumsString() == "3 1 2 6 5 4" {
// 		result = []string{"ra", "pb", "pb", "sa", "ra", "ra", "sa", "pa", "pa"}
// 		ps.Run(a, b, result)
// 		return result
// 	}
