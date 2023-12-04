package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	ins "push-swap/pkg/instructions"
	"push-swap/pkg/parse"
	"push-swap/pkg/structs"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	// This variable will be true if the program is receiving input from the
	// command line, and false if it is receiving input from a pipe. We'll use
	// this to determine whether or not to move the cursor up one line after
	// reading input. See 'Bitmasks: a Detour' in README.
	isTerminal := (fi.Mode() & os.ModeCharDevice) != 0
	s, err := run(os.Stdin, os.Stdout, isTerminal, os.Args)
	if err != nil {
		if err.Error() == "error to print" {
			fmt.Println("Error")
		}
		return
	}
	fmt.Println(s)
}

func run(r io.Reader, w io.Writer, isTerminal bool, args []string) (string, error) {
	err := errors.New("error not to print")
	if len(args) < 2 {
		return "", err
	}

	var a, b structs.Stack

	if parse.InitializeStacks(&a, &b) != nil {
		return "", err
	}

	instructionsToCheck, err := readInstructions(r, isTerminal)
	if err != nil {
		return "", err
	}
	runInstructions(&a, &b, instructionsToCheck)
	return checkStacks(a, b), nil
}

func checkStacks(a, b structs.Stack) string {
	if len(b.Nums) != 0 {
		return "KO"
	}
	for i := range a.Nums {
		if i == len(a.Nums)-1 {
			break
		}
		if a.Nums[i] > a.Nums[i+1] {
			return "KO"
		}
	}
	return "OK"
}

func runInstructions(a, b *structs.Stack, instructionsToCheck []string) {
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

func readInstructions(r io.Reader, isTerminal bool) ([]string, error) {
	instructionsToCheck := []string{}
	reader := bufio.NewReader(r)

	inputChan := make(chan string)
	errChan := make(chan error)

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
		close(inputChan)
	}()

	for {
		input, ok := <-inputChan
		if !ok {
			return instructionsToCheck, nil
		}
		if input == "" {
			if isTerminal {
				fmt.Print("\033[1A")
			}
			return instructionsToCheck, nil
		}
		instructionsToCheck = append(instructionsToCheck, input)
	}
}
