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

	// Create the context
	ctx := context.Background()

	// Create the Anthropic client
	llm, err := anthropic.New(anthropic.WithToken(anthropicApiKey), anthropic.WithModel("claude-3-5-haiku-latest"))
	if err != nil {
		log.Fatal(err)
	}

	// Generate the answer
	answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}

	return answer
}
