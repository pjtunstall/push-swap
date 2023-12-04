package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	ins "push-swap/pkg/instructions"
	parse "push-swap/pkg/parse"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var a ins.Stack
	var b ins.Stack

	if parse.InitializeStacks(&a, &b) != nil {
		fmt.Println("Error")
		return
	}

	instructionsToCheck, err := readInstructions()
	if err != nil {
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
	}
}

func readInstructions() ([]string, error) {
	instructionsToCheck := []string{}
	reader := bufio.NewReader(os.Stdin)

	inputChan := make(chan string)
	errChan := make(chan error)
	timeoutChan := make(chan []struct{})

	fi, _ := os.Stdin.Stat()

	go func() {
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				errChan <- err
			}
			inputChan <- strings.TrimSuffix(input, "\n")
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		timeoutChan <- []struct{}{}
	}()

	for {
		input := ""

		select {
		case input = <-inputChan:
			if input == "" {
				// If input is from the command line, move the cursor up one line
				// to remove the blank line that would otherwise be printed.
				// See 'Bitmasks: a Detour' in README.
				if (fi.Mode() & os.ModeCharDevice) != 0 {
					fmt.Print("\033[1A")
				}
				return instructionsToCheck, nil
			}
			instructionsToCheck = append(instructionsToCheck, input)
		case <-errChan:
			return nil, errors.New("failed to read input")
		case <-timeoutChan:
			return instructionsToCheck, nil
		}
	}
}
