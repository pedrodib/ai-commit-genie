package ai

import (
	"strings"
)

func GenerateCommitMessage(diff string) string {
	// Generate the prompt
	prompt := `You are an AI assistant that generates professional and concise git commit messages following the Conventional Commits specification.

		Each commit message should start with one of the following types, with their meanings:

		- feat: A new feature
		- fix: A bug fix
		- docs: Documentation only changes
		- style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
		- refactor: A code change that neither fixes a bug nor adds a feature
		- perf: A code change that improves performance
		- test: Adding missing tests or correcting existing tests
		- chore: Changes to the build process or auxiliary tools and libraries such as documentation generation

		Given the code changes in the diff below, write a commit message that:

		- Starts with the appropriate type from above
		- Uses the present tense
		- Is no longer than 72 characters in the subject line
		- Includes a short summary in the first line
		- Optionally adds a more detailed description separated by a blank line
		- Clearly describes what has been changed or fixed

		Here is the code diff:

		{diff}

		Please generate the appropriate commit message.`

	// Replace the diff in the prompt
	prompt = strings.ReplaceAll(prompt, "{diff}", diff)

	// Generate the answer
	return NewGemini(prompt)
}
