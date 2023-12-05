package main

import (
	"fmt"
	"os"

	"push-swap/ps"
)

func main() {
	var result string

	if len(os.Args) < 2 {
		return
	}

	// b, _ := ps.NewStack("")
	a, err := ps.NewStack(os.Args[1])
	if err != nil {
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
