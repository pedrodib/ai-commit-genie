package main

import (
	"os"
	"strings"
	"testing"
)

func TestAskForConfirmation(t *testing.T) {
	// Table-driven tests
	testCases := []struct {
		name     string // Name of the test case
		input    string // Simulated user input
		expected bool   // Expected result
	}{
		{"Yes response (y)", "y\n", true},
		{"Yes response (yes)", "yes\n", true},
		{"Yes response (uppercase)", "Y\n", true},
		{"No response (n)", "n\n", false},
		{"No response (no)", "no\n", false},
		{"No response (uppercase)", "N\n", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new reader from the simulated input string
			reader := strings.NewReader(tc.input)

			// The 'question' part doesn't matter for the test, as we discard the output
			// We pass a 'black-hole' writer (io.Discard) for the prompt output
			question := "Test question? "

			// We need to temporarily redirect os.Stdout to avoid printing the question during tests
			// This is a common pattern for testing functions that print to stdout
			oldStdout := os.Stdout
			_, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function with our mocked reader
			result := askForConfirmation(question, reader)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			if result != tc.expected {
				t.Errorf("For input '%s', expected %v, but got %v", tc.input, tc.expected, result)
			}
		})
	}
}
