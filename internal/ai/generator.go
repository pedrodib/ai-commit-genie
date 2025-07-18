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

func GetLLMProvider(llmProvider string) LLMStrategy {

	// Check if the requested LLM provider is supported
	if _, exists := supportedLLMs[llmProvider]; exists {
		return supportedLLMs[llmProvider]
	}

	// If not supported, default to OpenAI
	return supportedLLMs[GetDefaultProvider()]
}

// sanitizeCommitMessage removes markdown formatting and other unwanted characters from AI response
func sanitizeCommitMessage(message string) string {
	// Remove markdown code blocks (```)
	message = strings.ReplaceAll(message, "```", "")

	// Remove markdown inline code backticks
	message = strings.ReplaceAll(message, "`", "")

	// Remove extra quotes that might break git commit
	message = strings.ReplaceAll(message, "\"\"\"", "")
	message = strings.ReplaceAll(message, "'''", "")

	// Trim whitespace and newlines from start and end
	message = strings.TrimSpace(message)

	// Remove any leading/trailing quotes if they wrap the entire message
	if (strings.HasPrefix(message, "\"") && strings.HasSuffix(message, "\"")) ||
		(strings.HasPrefix(message, "'") && strings.HasSuffix(message, "'")) {
		message = message[1 : len(message)-1]
		message = strings.TrimSpace(message)
	}

	return message
}

func GenerateCommitMessage(diff string, langCode string, llmProvider string) string {
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

	IMPORTANT: 
	- You must write the entire commit message in ${langName}. Do not use English if the requested language is different.
	- Return ONLY the commit message text, without any markdown formatting, code blocks, or extra quotes.
	- Do not wrap the response in backticks, quotes, or any other formatting characters.

	Here is the code diff:

	{diff}

	Generate the commit message entirely in ${langName}:`

	// Getting provider Strategy based on llmProvider
	provider := GetLLMProvider(llmProvider)

	// Replace the diff and language name in the prompt
	prompt = strings.ReplaceAll(prompt, "{diff}", diff)
	prompt = strings.ReplaceAll(prompt, "${langName}", langName)

	// Generate the answer and sanitize it
	rawMessage := provider(prompt)
	return sanitizeCommitMessage(rawMessage)
}
