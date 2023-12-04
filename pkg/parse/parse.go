package parse

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"push-swap/pkg/structs"
)

func InitializeStacks(a, b *structs.Stack) error {
	if len(os.Args) < 2 {
		return errors.New("not enough arguments")
	}

	if os.Args[1] == "" {
		*a = structs.Stack{Top: -1, Nums: []int{}}
		*b = structs.Stack{Top: -1, Nums: []int{}}
		return nil
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

	*a = structs.Stack{Top: 0, Nums: numbers}
	*b = structs.Stack{Top: -1, Nums: []int{}}
	return nil
}

func AreThereDuplicates(a structs.Stack) bool {
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
