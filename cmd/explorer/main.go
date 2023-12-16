package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"push-swap/ps"
)

func seed(n int) string {
	nums := make([]string, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = strconv.Itoa(i)
	}
	return strings.Join(nums, " ")
}

func bfs(n int) {
	results := make(map[string][]string)

	q := [][]string{{"sa"}, {"ra"}, {"rra"}}
	for len(q) > 0 {
		v := q[0]
		if len(q) == 1 {
			q = [][]string{}
		} else {
			q = q[1:]
		}
		if len(v) > 10 {
			break
		}
		original := seed(5)
		a, _ := ps.NewStack(original)
		b, _ := ps.NewStack("")
		inv := inverse(v)
		ps.Run(&a, &b, inv)
		nums := make([]int, 5)
		copy(nums, a.GetNumsSlice())
		numsString := a.GetNumsString()
		five := five(&a, &b)
		a, _ = ps.NewStack(numsString)
		general := general(&a, &b)
		sorted, _ := ps.Check(a, b)
		if !sorted {
			fmt.Println("Not sorted:", a.GetNumsString())
		} else {
			m := min(len(five), len(general))
			if len(v) < m {
				fmt.Println(numsString)
				fmt.Println("************************** bfs is shortest")
				fmt.Println("length of `five`: ", len(five))
				fmt.Println("length of `general`: ", len(general))
				fmt.Println("length of `bfs`: ", len(v))
				fmt.Println("`five`: ", five)
				fmt.Println("`general`: ", general)
				fmt.Println("`bfs`: ", v)
				fmt.Println()
				existing, ok := results[numsString]
				if !ok || len(v) < len(existing) {
					results[numsString] = v
				}
			}
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

func main() {
	bfs(5)
}

func inverse(instructions []string) []string {
	var result []string
	for i := len(instructions) - 1; i >= 0; i-- {
		switch instructions[i] {
		case "sa":
			result = append(result, "sa")
		case "sb":
			result = append(result, "sb")
		case "ss":
			result = append(result, "ss")
		case "pa":
			result = append(result, "pb")
		case "pb":
			result = append(result, "pa")
		case "ra":
			result = append(result, "rra")
		case "rb":
			result = append(result, "rrb")
		case "rr":
			result = append(result, "rrr")
		case "rra":
			result = append(result, "ra")
		case "rrb":
			result = append(result, "rb")
		case "rrr":
			result = append(result, "rr")
		}
	}
	return result
}

// func listFive() {
// 	for i := 1; i <= 5; i++ {
// 		for j := 1; j <= 5; j++ {
// 			for k := 1; k <= 5; k++ {
// 				for l := 1; l <= 5; l++ {
// 					for m := 1; m <= 5; m++ {
// 						if (i == j || i == k || i == l || i == m) || (j == k || j == l || j == m) || k == l || k == m || (l == m) {
// 							continue
// 						}
// 						input := fmt.Sprintf("%d %d %d %d %d", i, j, k, l, m)
// 						a, _ := ps.NewStack(input)
// 						b, _ := ps.NewStack("")
// 						ins := five(&a, &b)
// 						fmt.Printf("%v: %v", input, ins)

// 						_, sorted := ps.Check(a, b)
// 						if !sorted {
// 							fmt.Println(". Not sorted: ", a.GetNumsString())
// 						} else {
// 							fmt.Println()
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

func five(a, b *ps.Stack) []string {
	var result []string
	nums := a.GetNumsSlice()
	left := make([]int, 3, 5)
	var maxB int
	var minB int
	var leftRot bool
	var combineRotation bool

	// swapRotatable, swappable := swapRot(*a, *b)
	// if swappable {
	// 	ps.Sx(a)
	// 	return []string{"sa"}
	// }
	// if swapRotatable {
	// 	ps.Sx(a)
	// 	rots := justRotate(*a)
	// 	ps.Run(a, b, rots)
	// 	return append([]string{"sa"}, rots...)
	// }

	// rotSwapScript, rotSwappable := rotSwap(*a, *b)
	// if rotSwappable {
	// 	ps.Run(a, b, rotSwapScript)
	// 	return rotSwapScript
	// }

	result = []string{"pb", "pb"} // Push top two to B.
	copy(left, nums[2:])          // The three that stay in A.
	_, leftRot = three(left)      // `stayersRot` is true if no swap is needed to sort them.
	maxB = max(nums[0], nums[1])
	minB = min(nums[0], nums[1])

	if leftRot {
		if nums[0] == maxB {
			combineRotation = true
		}
	} else {
		left[0], left[1] = left[1], left[0]
		if nums[0] == maxB {
			result = append(result, "ss")
		} else {
			result = append(result, "sa")
		}
	}

	// Consider maxB at top of B now, unless combineRotation is true.
	switch fitTheFourth(maxB, left) {
	case 0:
		if combineRotation {
			result = append(result, "sb")
		}
	case 1:
		if combineRotation {
			result = append(result, "rr")
		} else {
			result = append(result, "ra")
		}
		left = append(left[1:], left[0])
	case 2:
		if combineRotation {
			result = append(result, "rrr")
		} else {
			result = append(result, "rra")
		}
		left = append(left[len(left)-1:], left[:len(left)-1]...)
	}
	if result[len(result)-1] == "pb" {
		result = result[:len(result)-1]
	} else {
		result = append(result, "pa")
	}
	left = append([]int{maxB}, left...)

	switch fitTheFifth(minB, left) {
	case 1:
		result = append(result, "ra")
		left = append(left[1:], left[0])
	case 2:
		result = append(result, "ra", "ra")
		left = append(left[2:], left[:2]...)
	case 3:
		result = append(result, "rra")
		left = append(left[len(left)-1:], left[:len(left)-1]...)
	}
	result = append(result, "pa")
	left = append([]int{minB}, left...)

	copy(a.Nums, left)
	a.Top = 0
	rots := justRotate(*a)
	result = append(result, justRotate(*a)...)
	ps.Run(a, b, rots)

	return result
}

func fitTheFourth(x int, left []int) int {
	var position int
	iMax, maxStayer := ps.MaxInt(left)
	iMin, minStayer := ps.MinInt(left)

	if x < minStayer {
		position = iMin
	} else if x > maxStayer {
		position = iMax + 1
	} else {
		if x > left[0] && x < left[1] {
			position = 1
		} else if x > left[1] && x < left[2] {
			position = 2
		} else if x < left[0] && x > left[2] {
			position = 0
		}
	}
	return position
}

func fitTheFifth(x int, left []int) int {
	var position int
	iMin, minStayer := ps.MinInt(left)

	if x < minStayer {
		position = iMin
	} else {
		if x > left[0] && x < left[1] {
			position = 1
		} else if x > left[1] && x < left[2] {
			position = 2
		} else if x > left[2] && x < left[3] {
			position = 3
		} else {
			position = 0
		}
	}

	return position
}

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

func general(a, b *ps.Stack) []string {
	var result []string

	// Just in case any preliminary checks altered the stacks. May not be necessary.
	*a, _ = ps.NewStack(rank(a.Nums))

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

	ps.Px(b, a)
	ps.Px(b, a)
	result = append(result, "pb", "pb")
	nums := b.GetNumsSlice()
	if nums[0] < nums[1] {
		ps.Sx(b)
		result = append(result, "sb")
	}
	result = append(result, sortToB(a, b)...)
	_, rotatable := three(a.Nums)
	if !rotatable {
		ps.Sx(a)
		result = append(result, "sa")
	}
	result = append(result, sortToA(a, b)...)

	result = append(result, justRotate(*a)...)
	ps.Run(a, b, justRotate(*a))

	// fmt.Println("A:", a.GetNumsSlice())
	// fmt.Println("B:", b.GetNumsSlice())
	// fmt.Println()

	return result
}

func sortToB(a, b *ps.Stack) []string {
	result := []string{}
	for {
		A := a.GetNumsSlice()
		B := b.GetNumsSlice()
		journeyPlanner := make([]ps.PushInfo, len(A))

		if len(A) == 3 {
			break
		}

		cheapest := 0
		for i, v := range A {
			var cost int
			var ra, rb bool
			var stepsA, stepsB, jointSteps int

			if i > len(A)/2 {
				cost = len(A) - i
				stepsA = len(A) - i
			} else {
				cost = i
				stepsA = i
				ra = true
			}

			foundOneLess := false
			targetIndex, targetValue := ps.MinInt(B)
			for j, w := range B {
				if w < v {
					foundOneLess = true
					if w > targetValue {
						targetValue = w
						targetIndex = j
					}
				}
			}
			if !foundOneLess {
				targetIndex, targetValue = ps.MaxInt(B)
			}

			if targetIndex > len(B)/2 {
				cost += len(B) - targetIndex
				stepsB = len(B) - targetIndex
			} else {
				cost += targetIndex
				stepsB = targetIndex
				rb = true
			}

			jointSteps = min(stepsA, stepsB)

			// // Optimization: Account for case where ra XOR rb, but
			// // either of the stacks would be rotated len(X)/2 times.
			// // In that case, it doesn't matter which way we rotate the
			// // stack that needs rotating len(X)/2 times, so we can
			// // set its direction of rotation to the same as the other
			// // to take advantage of any possible joint steps.
			// // But understand why the same code DOESN'T work in
			// // sortToA, and indeed fails to sort.
			// if stepsA == len(A)/2 {
			// 	ra = rb
			// }
			// if stepsB == len(B)/2 {
			// 	rb = ra
			// }

			if (ra && rb) || (!ra && !rb) {
				cost -= jointSteps
				stepsA -= jointSteps
				stepsB -= jointSteps
			}

			journeyPlanner[i] = ps.PushInfo{
				Value:       v,
				TargetIndex: targetIndex,
				TargetValue: targetValue,
				Cost:        cost,
				Ra:          ra,
				Rb:          rb,
				StepsA:      stepsA,
				StepsB:      stepsB,
				JointSteps:  jointSteps}

			if cost == 0 {
				break
			}
			if cost < journeyPlanner[cheapest].Cost {
				cheapest = i
			}
		}

		c := journeyPlanner[cheapest]

		// fmt.Println("A:", A)
		// fmt.Println("B:", B)
		// fmt.Println()

		// fmt.Println("cheapest:", c.Value)
		// fmt.Println("targetIndex:", c.TargetIndex)
		// fmt.Println("targetValue:", c.TargetValue)
		// fmt.Println("cost:", c.Cost)
		// fmt.Println("ra:", c.Ra)
		// fmt.Println("rb:", c.Rb)
		// fmt.Println("stepsA:", c.StepsA)
		// fmt.Println("stepsB:", c.StepsB)
		// fmt.Println("jointSteps:", c.JointSteps)
		// fmt.Println()

		if c.Ra && c.Rb {
			for j := 0; j < c.JointSteps; j++ {
				ps.Rr(b, a)
				result = append(result, "rr")
			}
		} else if !c.Ra && !c.Rb {
			for j := 0; j < c.JointSteps; j++ {
				ps.Rrr(a, b)
				result = append(result, "rrr")
			}
		}
		if c.Ra {
			for j := 0; j < c.StepsA; j++ {
				ps.Rx(a)
				result = append(result, "ra")
			}
		} else {
			for j := 0; j < c.StepsA; j++ {
				ps.Rrx(a)
				result = append(result, "rra")
			}
		}
		if c.Rb {
			for j := 0; j < c.StepsB; j++ {
				ps.Rx(b)
				result = append(result, "rb")
			}
		} else {
			for j := 0; j < c.StepsB; j++ {
				ps.Rrx(b)
				result = append(result, "rrb")
			}
		}
		ps.Px(b, a)
		result = append(result, "pb")

		// fmt.Println("A:", a.GetNumsSlice())
		// fmt.Println("B:", b.GetNumsSlice())
		// fmt.Println()

		// for _, v := range journeyPlanner {
		// 	fmt.Printf("%+v\n", v)
		// }
	}

	return result
}

func sortToA(a, b *ps.Stack) []string {
	result := []string{}
	for {
		A := a.GetNumsSlice()
		B := b.GetNumsSlice()
		journeyPlanner := make([]ps.PushInfo, len(B))

		if len(B) == 0 {
			break
		}

		cheapest := 0
		for i, v := range B {
			var cost int
			var ra, rb bool
			var stepsA, stepsB, jointSteps int

			if i > len(B)/2 {
				cost = len(B) - i
			} else {
				cost = i
				rb = true
			}
			stepsB = cost

			foundOneGreater := false
			targetIndex, targetValue := ps.MaxInt(A)
			for j, w := range A {
				if w > v {
					foundOneGreater = true
					if w < targetValue {
						targetValue = w
						targetIndex = j
					}
				}
			}
			if !foundOneGreater {
				targetIndex, targetValue = ps.MinInt(A)
			}

			if targetIndex > len(A)/2 {
				cost += len(A) - targetIndex
				stepsA = len(A) - targetIndex
			} else {
				cost += targetIndex
				stepsA = targetIndex
				ra = true
			}

			jointSteps = min(stepsA, stepsB)

			// // The following doesn't work here, although it does in
			// // sortToB. I'm not sure why. I guessed it was to do with
			// // a case where the initial pushes to B are not sorted,
			// // but that would suggest the problem would happen in
			// // sortToB too.

			// // Optimization: Account for case where ra XOR rb, but
			// // either of the stacks would be rotated len(X)/2 times.
			// // In that case, it doesn't matter which way we rotate the
			// // stack that needs rotating len(X)/2 times, so we can
			// // set its direction of rotation to the same as the other
			// // to take advantage of any possible joint steps.
			// if stepsA == len(A)/2 {
			// 	ra = rb
			// }
			// if stepsB == len(B)/2 {
			// 	rb = ra
			// }

			if (ra && rb) || (!ra && !rb) {
				cost -= jointSteps
				stepsA -= jointSteps
				stepsB -= jointSteps
			}

			journeyPlanner[i] = ps.PushInfo{
				Value:       v,
				TargetIndex: targetIndex,
				TargetValue: targetValue,
				Cost:        cost,
				Ra:          ra,
				Rb:          rb,
				StepsA:      stepsA,
				StepsB:      stepsB,
				JointSteps:  jointSteps}

			if cost == 0 {
				break
			}
			if cost < journeyPlanner[cheapest].Cost {
				cheapest = i
			}
		}

		c := journeyPlanner[cheapest]

		// fmt.Println("cheapest:", c.Value)
		// fmt.Println("targetIndex:", c.TargetIndex)
		// fmt.Println("targetValue:", c.TargetValue)
		// fmt.Println("cost:", c.Cost)
		// fmt.Println("ra:", c.Ra)
		// fmt.Println("rb:", c.Rb)
		// fmt.Println("stepsA:", c.StepsA)
		// fmt.Println("stepsB:", c.StepsB)
		// fmt.Println("jointSteps:", c.JointSteps)
		// fmt.Println()

		if c.Ra && c.Rb {
			for j := 0; j < c.JointSteps; j++ {
				ps.Rr(a, b)
				result = append(result, "rr")
			}
		} else if !c.Ra && !c.Rb {
			for j := 0; j < c.JointSteps; j++ {
				ps.Rrr(a, b)
				result = append(result, "rrr")
			}
		}
		if c.Ra {
			for j := 0; j < c.StepsA; j++ {
				ps.Rx(a)
				result = append(result, "ra")
			}
		} else {
			for j := 0; j < c.StepsA; j++ {
				ps.Rrx(a)
				result = append(result, "rra")
			}
		}
		if c.Rb {
			for j := 0; j < c.StepsB; j++ {
				ps.Rx(b)
				result = append(result, "rb")
			}
		} else {
			for j := 0; j < c.StepsB; j++ {
				ps.Rrx(b)
				result = append(result, "rrb")
			}
		}
		ps.Px(a, b)
		result = append(result, "pa")

		// for _, v := range journeyPlanner {
		// 	fmt.Printf("%+v\n", v)
		// }
	}

	return result
}

func rank(nums []int) string {
	rankMap := make(map[int]int)
	strs := make([]string, len(nums))
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	sort.Ints(numsCopy)
	for i, v := range numsCopy {
		rankMap[v] = i + 1
	}

	for i, v := range nums {
		strs[i] = strconv.Itoa(rankMap[v])
	}
	s := strings.Join(strs, " ")

	return s
}
