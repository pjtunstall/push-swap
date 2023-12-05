package ps

import (
	"errors"
	"strconv"
	"strings"
)

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
