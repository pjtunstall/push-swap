package main

import (
	"fmt"
	"os"

	"push-swap/pkg/parse"
	"push-swap/pkg/structs"
)

func main() {
	var result string
	var a structs.Stack
	var b structs.Stack
	wrong := false

	err := parse.InitializeStacks(&a, &b)
	if err != nil {
		if err.Error() == "not enough arguments" {
			return
		}
		wrong = true
	}
	if parse.AreThereDuplicates(a) {
		wrong = true
	}
	if wrong {
		fmt.Fprintf(os.Stderr, "Error\n")
		return
	}

	switch len(a.Nums) {
	case 0:
		return
	case 1:
		return
	case 2:
		if a.Nums[0] < a.Nums[1] {
			return
		}
		result = "sa\n"
	case 3:
		result = three(a.Nums)
	}

	if len(result) == 0 {
		return
	}
	fmt.Print(result)
}

func three(nums []int) string {
	if nums[2] < nums[0] && nums[1] < nums[2] {
		return "ra\n"
	}
	if nums[0] < nums[1] && nums[2] < nums[0] {
		return "rra\n"
	}
	if nums[1] < nums[0] && nums[0] < nums[2] {
		return "sa\n"
	}
	if nums[0] < nums[1] && nums[2] < nums[1] {
		return "sa\nra\n"
	}
	if nums[1] < nums[0] && nums[2] < nums[1] {
		return "sa\nrra\n"
	}
	return ""
}
