package ps

import (
	"errors"
	"strconv"
	"strings"
)

type PushInfo struct {
	Index       int
	Value       int
	TargetIndex int
	TargetValue int
	Cost        int
	Ra          bool
	Rb          bool
	StepsA      int
	StepsB      int
	JointSteps  int
}

type Stack struct {
	Top  int
	Nums []int
}

func NewStack(str string) (Stack, error) {
	var x Stack

	if str == "" {
		x = Stack{Top: -1, Nums: []int{}}
		return x, nil
	}

	source := strings.Split(str, (" "))
	numbers := []int{}

	for _, v := range source {
		n, err := strconv.Atoi(v)
		if err != nil {
			return x, errors.New("invalid argument")
		}
		numbers = append(numbers, n)
	}

	x = Stack{Top: 0, Nums: numbers}
	if AreThereDuplicates(x) {
		return x, errors.New("duplicate numbers")
	}

	return x, nil
}

func (x Stack) Copy() Stack {
	var y Stack
	y.Top = x.Top
	y.Nums = make([]int, len(x.Nums))
	copy(y.Nums, x.Nums)
	return y
}

func (x Stack) GetNumsSlice() []int {
	if len(x.Nums) == 0 {
		return []int{}
	}
	result := make([]int, len(x.Nums))
	head := x.Nums[x.Top:]
	tail := x.Nums[:x.Top]
	copy(result, append(head, tail...))
	return result
}

func (x Stack) GetNumsString() string {
	if len(x.Nums) == 0 {
		return ""
	}
	nums := make([]int, len(x.Nums))
	head := x.Nums[x.Top:]
	tail := x.Nums[:x.Top]
	copy(nums, append(head, tail...))
	var sb strings.Builder
	for i, num := range nums {
		sb.WriteString(strconv.Itoa(num))
		if i == len(nums)-1 {
			break
		}
		sb.WriteString(" ")
	}
	return sb.String()
}

func (x Stack) Max() (int, int, error) {
	if len(x.Nums) == 0 {
		return 0, 0, errors.New("empty stack")
	}

	max := x.Nums[0]
	index := 0
	for i, v := range x.Nums {
		if v > max {
			max = v
			index = i
		}
	}
	return index, max, nil
}

func (x Stack) Min() (int, int, error) {
	if len(x.Nums) == 0 {
		return 0, 0, errors.New("empty stack")
	}

	min := x.Nums[0]
	index := 0
	for i, v := range x.Nums {
		if v < min {
			min = v
			index = i
		}
	}
	return index, min, nil
}
