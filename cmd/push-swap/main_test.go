package main_test

import (
	"io"
	"os/exec"
	"strings"
	"testing"
)

// TestChecker tests the push-swap program against the checker program.
// It assumes there is a push-swap executable in the current directory
// and a checker executable in the checker directory.
func TestChecker(t *testing.T) {
	tests := []string{
		"240 108 360 420",
		"0 -1 666 333",
		"2 3 1 4",
		"2 3 4 1",
		"2 4 1 3",
		"2 4 3 1",
		"3 1 2 4",
		"3 2 1 4 5",
		"2 1 3 4 5",
	}

	for _, tc := range tests {
		mainCmd := exec.Command("./push-swap", tc)
		mainOutput, err := mainCmd.Output()
		if err != nil {
			t.Fatalf("Failed to execute main command: %v", err)
		}

		// Convert mainOutput to a string
		mainOutputStr := string(mainOutput)

		// Run the checker command and pass the output of the main function as an argument
		checkerCmd := exec.Command("../checker/checker", tc)
		stdin, err := checkerCmd.StdinPipe()
		if err != nil {
			t.Fatalf("Failed to get stdin pipe: %v", err)
		}

		go func() {
			defer stdin.Close()
			io.WriteString(stdin, mainOutputStr)
		}()

		output, err := checkerCmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to execute checker command: %v", err)
		}

		// Check the output of the checker command
		outputStr := string(output)
		if !strings.HasPrefix(outputStr, "OK") {
			t.Fatalf("Checker output on %v was not OK: %s", tc, outputStr)
		}
	}
}
