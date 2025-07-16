package git

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// sanitizeCommitMessage removes potentially dangerous characters from commit message
func sanitizeCommitMessage(message string) string {
	// Remove null bytes and control characters
	message = strings.ReplaceAll(message, "\x00", "")

	// Remove or escape shell metacharacters that could be dangerous
	dangerousChars := []string{";", "&", "|", "`", "$", "(", ")", "<", ">", "\\"}
	for _, char := range dangerousChars {
		message = strings.ReplaceAll(message, char, "")
	}

	// Remove excessive newlines (keep max 5 consecutive)
	re := regexp.MustCompile(`\n{6,}`)
	message = re.ReplaceAllString(message, "\n\n")

	// Trim and limit length for safety
	message = strings.TrimSpace(message)
	if len(message) > 2000 { // Reasonable limit for commit messages
		message = message[:2000]
	}

	return message
}

// ExecuteCommit runs the 'git commit -m' command with the provided message.
// It returns the combined output of the command and an error if it fails.
func ExecuteCommit(message string) (string, error) {
	// Sanitize the commit message to prevent command injection
	sanitizedMessage := sanitizeCommitMessage(message)

	// Validate that the message is not empty after sanitization
	if strings.TrimSpace(sanitizedMessage) == "" {
		return "", fmt.Errorf("commit message is empty after sanitization")
	}

	cmd := exec.Command("git", "commit", "-m", sanitizedMessage)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// The error from CombinedOutput already includes stderr, so we return it directly.
		return string(output), fmt.Errorf("error executing git commit: %w", err)
	}

	return string(output), nil
}
