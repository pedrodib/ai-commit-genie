package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetStagedDiff() (string, error) {
	// Check if the current directory is a git repository
	checkCmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	var checkStderr bytes.Buffer
	checkCmd.Stderr = &checkStderr
	err := checkCmd.Run()
	if err != nil {
		return "", fmt.Errorf("not in a git repository: %v\n%s", err, checkStderr.String())
	}

	// Get the staged diff
	cmd := exec.Command("git", "diff", "--cached")
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Run the command
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error executing git diff: %v\n%s", err, stderr.String())
	}

	// Return the diff
	return out.String(), nil
}
