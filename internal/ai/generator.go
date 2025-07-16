package ai

import (
	"strings"
)

// Map of language codes to their full names
var languageNames = map[string]string{
	"en": "English",
	"pt": "Português",
	"es": "Español",
	"fr": "Français",
	"de": "Deutsch",
	"zh": "中文",
	"ja": "日本語",
	"ru": "Русский",
}

// GetSupportedLanguages returns a map of supported language codes and their names
func GetSupportedLanguages() map[string]string {
	return languageNames
}

func GenerateCommitMessage(diff string, langCode string) string {
	// Get the language name, default to English if not found
	langName, exists := languageNames[langCode]
	if !exists {
		langName = "English"
	}

	// Create a single prompt in English with clear language instruction
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

	IMPORTANT: You must write the entire commit message in ${langName}. Do not use English if the requested language is different.

	Here is the code diff:

	{diff}

	Generate the commit message entirely in ${langName}:`

	// Replace the diff and language name in the prompt
	prompt = strings.ReplaceAll(prompt, "{diff}", diff)
	prompt = strings.ReplaceAll(prompt, "${langName}", langName)

	// Generate the answer
	return NewGemini(prompt)
}
