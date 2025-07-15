package git

import (
	"fmt"
	"os/exec"
)

// ExecuteCommit runs the 'git commit -m' command with the provided message.
// It returns the combined output of the command and an error if it fails.
func ExecuteCommit(message string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", message)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// The error from CombinedOutput already includes stderr, so we return it directly.
		return string(output), fmt.Errorf("error executing git commit: %w", err)
	}

	return string(output), nil
}
