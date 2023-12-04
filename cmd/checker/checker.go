package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	ins "push-swap/pkg/instructions"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var a ins.Stack
	var b ins.Stack

	instructionsToCheck, err := initialize(&a, &b)
	if err != nil {
		fmt.Println("Error")
		return
	}
	runInstructions(&a, &b, instructionsToCheck)
	checkStacks(a, b)
}

func checkStacks(a, b ins.Stack) {
	if len(b.Nums) != 0 {
		fmt.Println("KO")
		return
	}
	for i := range a.Nums {
		if i == len(a.Nums)-1 {
			break
		}
		if a.Nums[i] > a.Nums[i+1] {
			fmt.Println("KO")
			return
		}
	}
	fmt.Println("OK")
}

func runInstructions(a, b *ins.Stack, instructionsToCheck []string) {
	// fmt.Println(instructions)
	// fmt.Println("--------------------------------------- Initial")
	// fmt.Print("a: ")
	// for i, v := range a.Nums {
	// 	if i == a.Top {
	// 		fmt.Printf("[%d] ", v)
	// 	} else {
	// 		fmt.Printf("%d ", v)
	// 	}
	// }
	// fmt.Print("\nb:")
	for _, instruction := range instructionsToCheck {
		switch instruction {
		case "sa":
			ins.Sx(a)
		case "sb":
			ins.Sx(b)
		case "ss":
			ins.Ss(a, b)
		case "pa":
			ins.Px(a, b)
		case "pb":
			ins.Px(b, a)
		case "ra":
			ins.Rx(a)
		case "rb":
			ins.Rx(b)
		case "rr":
			ins.Rr(a, b)
		case "rra":
			ins.Rrx(a)
		case "rrb":
			ins.Rrx(b)
		case "rrr":
			ins.Rrr(a, b)
		}
		// fmt.Println("\n---------------------------------------", instruction)
		// fmt.Print("a: ")
		// for i, v := range a.Nums {
		// 	if i == a.Top {
		// 		fmt.Printf("[%d] ", v)
		// 	} else {
		// 		fmt.Printf("%d ", v)
		// 	}
		// }
		// fmt.Print("\nb: ")
		// for i, v := range b.Nums {
		// 	if i == b.Top {
		// 		fmt.Printf("[%d] ", v)
		// 	} else {
		// 		fmt.Printf("%d ", v)
		// 	}
		// }
	}
	// fmt.Println()
}

func initialize(a, b *ins.Stack) ([]string, error) {
	source := strings.Split(os.Args[1], (" "))
	instructionsToCheck := []string{}
	numbers := []int{}

	// TODO: check for errors and edge cases such as empty input

	for _, v := range source {
		n, err := strconv.Atoi(v)
		if err != nil {
			return []string{}, err
		}
		numbers = append(numbers, n)
	}

	*a = ins.Stack{Top: 0, Nums: numbers}
	*b = ins.Stack{Top: -1, Nums: []int{}}

	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return []string{}, err
		}
		input = strings.TrimSuffix(input, "\n")
		instructionsToCheck = append(instructionsToCheck, input)
	}
	instructionsToCheck = instructionsToCheck[:len(instructionsToCheck)-1]
	return instructionsToCheck, nil
}
