package main

import (
	"fmt"

	ins "push-swap/pkg/instructions"
	parse "push-swap/pkg/parse"
)

func main() {
	var a ins.Stack
	var b ins.Stack

	err := parse.InitializeStacks(&a, &b)
	if err != nil {
		if err.Error() == "not enough arguments" {
			return
		}
		fmt.Println("Error")
		return
	}
	if parse.CheckForDuplicates(a) {
		fmt.Println("Error")
		return
	}

	fmt.Println(a.Nums)
}
