package main

import (
	"fmt"
	"os"
	"strings"

	"push-swap/ps"
)

func main() {
	var result []string

	if len(os.Args) < 2 {
		return
	}

	if os.Args[1] == "" {
		fmt.Println("Be reasonable! Any commands will sort an empty stack.")
		return
	}

	b, _ := ps.NewStack("")
	a, err := ps.NewStack(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error\n")
		return
	}

	rotatable, sorted := ps.Check(a, b)

	if rotatable {
		if sorted {
			result = []string{}
		} else {
			result = justRotate(a)
		}
	} else {
		switch len(a.Nums) {
		case 0:
			return
		case 1:
			return
		case 2:
			if a.Nums[0] < a.Nums[1] {
				return
			}
			result = []string{"sa"}
		case 3:
			result, _ = three(a.Nums)
		case 4:
			result = four(a, b)
		case 5:
			result = five(&a, &b)
		default:
			result = general(&a, &b)
		}
	}

	if len(result) == 0 {
		return
	}

	var sb strings.Builder
	for _, str := range result {
		sb.WriteString(str)
		sb.WriteRune('\n')
	}
	fmt.Print(sb.String())

	// fmt.Print("\n", len(result), "\n")
}
