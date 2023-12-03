package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	top  int
	nums []int
}

var a stack
var b stack
var instructions []string

func main() {
	if len(os.Args) < 2 {
		return
	}
	err := initialize()
	if err != nil {
		log.Fatal(err)
	}
	runInstructions()
	checkStacks()
}

func checkStacks() {
	if len(b.nums) != 0 {
		fmt.Println("KO")
		return
	}
	for i := range a.nums {
		if i == len(a.nums)-1 {
			break
		}
		if a.nums[i] > a.nums[i+1] {
			fmt.Println("KO")
			return
		}
	}
	fmt.Println("OK")
}

func runInstructions() {
	// fmt.Println(instructions)
	// fmt.Println("--------------------------------------- Initial")
	// fmt.Print("a: ")
	// for i, v := range a.nums {
	// 	if i == a.top {
	// 		fmt.Printf("[%d] ", v)
	// 	} else {
	// 		fmt.Printf("%d ", v)
	// 	}
	// }
	// fmt.Print("\nb:")
	for _, instruction := range instructions {
		switch instruction {
		case "sa":
			sx(&a)
		case "sb":
			sx(&b)
		case "ss":
			ss()
		case "pa":
			px(&a, &b)
		case "pb":
			px(&b, &a)
		case "ra":
			rx(&a)
		case "rb":
			rx(&b)
		case "rr":
			rr()
		case "rra":
			rrx(&a)
		case "rrb":
			rrx(&b)
		case "rrr":
			rrr()
		}
		// fmt.Println("\n---------------------------------------", instruction)
		// fmt.Print("a: ")
		// for i, v := range a.nums {
		// 	if i == a.top {
		// 		fmt.Printf("[%d] ", v)
		// 	} else {
		// 		fmt.Printf("%d ", v)
		// 	}
		// }
		// fmt.Print("\nb: ")
		// for i, v := range b.nums {
		// 	if i == b.top {
		// 		fmt.Printf("[%d] ", v)
		// 	} else {
		// 		fmt.Printf("%d ", v)
		// 	}
		// }
	}
	// fmt.Println()
}

func initialize() error {
	source := strings.Split(os.Args[1], (" "))
	nums := []int{}

	// TODO: check for errors and edge cases such as empty input

	for _, v := range source {
		n, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		nums = append(nums, n)
	}

	a = stack{top: 0, nums: nums}
	b = stack{top: -1, nums: []int{}}

	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		input = strings.TrimSuffix(input, "\n")
		instructions = append(instructions, input)
	}
	instructions = instructions[:len(instructions)-1]
	return nil
}

// Instructions
func rx(x *stack) {
	if len(x.nums) < 2 {
		return
	}
	if x.top == len(x.nums)-1 {
		x.top = 0
		return
	}
	x.top += 1
}

func rrx(x *stack) {
	if len(x.nums) < 2 {
		return
	}
	if x.top == 0 {
		x.top = len(x.nums) - 1
		return
	}
	x.top -= 1
}

func rr() {
	rx(&a)
	rx(&b)
}

func rrr() {
	rrx(&a)
	rrx(&b)
}

func sx(x *stack) {
	if len(x.nums) < 2 {
		return
	}
	if x.top == len(x.nums)-1 {
		x.nums[x.top], x.nums[0] = x.nums[0], x.nums[x.top]
		return
	}
	x.nums[x.top], x.nums[x.top+1] = x.nums[x.top+1], x.nums[x.top]
}

func ss() {
	sx(&a)
	sx(&b)
}

func px(x *stack, y *stack) {
	if len(y.nums) == 0 {
		return
	}
	if len(x.nums) == 0 {
		x.nums = []int{y.nums[y.top]}
		x.top = 0
	} else {
		head := append([]int{y.nums[y.top]}, x.nums[x.top:]...)
		x.nums = append(head, x.nums[:x.top]...)
	}

	if len(y.nums) == 1 {
		y.nums = []int{}
		y.top = -1
		return
	}
	if y.top == 0 {
		y.nums = y.nums[1:]
		return
	}
	if y.top == len(y.nums)-1 {
		y.nums = y.nums[:y.top]
		y.top = 0
		return
	}
	y.nums = append(y.nums[:y.top], y.nums[y.top+1:]...)
}
