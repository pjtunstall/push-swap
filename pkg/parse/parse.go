package parse

import (
	"errors"
	"os"
	"strconv"
	"strings"

	ins "push-swap/pkg/instructions"
)

func InitializeStacks(a, b *ins.Stack) error {
	if len(os.Args) < 2 {
		return errors.New("not enough arguments")
	}
	source := strings.Split(os.Args[1], (" "))
	numbers := []int{}

	for _, v := range source {
		n, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		numbers = append(numbers, n)
	}

	*a = ins.Stack{Top: 0, Nums: numbers}
	*b = ins.Stack{Top: -1, Nums: []int{}}
	return nil
}

func CheckForDuplicates(a ins.Stack) bool {
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
