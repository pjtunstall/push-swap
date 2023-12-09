package main

import (
	"fmt"
	"log"
	"os"

	"push-swap/ps"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	b, _ := ps.NewStack("")
	a, err := ps.NewStack(os.Args[1])
	if err != nil {
		if err.Error() == "invalid argument" {
			fmt.Println("Error: invalid argument")
		}
		if err.Error() == "duplicate numbers" {
			fmt.Println("Error: duplicate numbers")
		}
		return
	}

	// The following variables are used to determine whether or not to move the
	// cursor up one line after reading input. `isTerminal` will be true if the
	// program is receiving input from the command line, and false if it's
	// receiving input from a pipe. See 'Bitmasks: a Detour' in README.
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	isTerminal := (fi.Mode() & os.ModeCharDevice) != 0

	instructions, err := getInstructions(os.Stdin, isTerminal)
	if err != nil {
		fmt.Println("Failed to get instructions.")
		return
	}

	if ps.Run(&a, &b, instructions) != nil {
		fmt.Println("Error: invalid instruction")
		return
	}
	_, sorted := ps.Check(a, b)

	if sorted {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}

	fmt.Printf("a: %v, a.Top: %v\n", a.Nums, a.Top)
	fmt.Printf("b: %v, b.Top %v\n", b.Nums, b.Top)

	fmt.Println("Interpret this as the numbers sorted to:", append(a.Nums[a.Top:], a.Nums[:a.Top]...))
}
