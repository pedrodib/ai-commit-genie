package ai

import (
	"context"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
)

func init() {
	RegisterLLM("anthropic", GenerateWithAnthropic)
}

func GenerateWithAnthropic(prompt string) string {
	// Get the Anthropic API key
	anthropicApiKey := os.Getenv("ANTHROPIC_API_KEY")

	// Validate API key exists
	if anthropicApiKey == "" {
		log.Fatal("ANTHROPIC_API_KEY environment variable is not set")
	}

	// Create the context
	ctx := context.Background()

	// Create the Anthropic client
	llm, err := anthropic.New(anthropic.WithToken(anthropicApiKey), anthropic.WithModel("claude-3-5-haiku-latest"))
	if err != nil {
		log.Fatal("Failed to initialize Anthropic client. Please check your API key configuration.")
	}

	// Generate the answer
	answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal("Failed to generate response from Anthropic. Please check your API key and network connection.")
	}

	return answer
}
