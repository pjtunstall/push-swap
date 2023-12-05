package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func getInstructions(r io.Reader, isTerminal bool) ([]string, error) {
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
