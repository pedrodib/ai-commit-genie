package ai

import (
	"context"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func init() {
	RegisterLLM("openai", GenerateWithOpenAi)
}

func GenerateWithOpenAi(prompt string) string {
	// Get the OpenAi API key
	openAiApiKey := os.Getenv("OPENAI_API_KEY")

	// Validate API key exists
	if openAiApiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}

	// Create the context
	ctx := context.Background()

	// Create the OpenAi client
	llm, err := openai.New(openai.WithToken(openAiApiKey), openai.WithModel("gpt-4.1-mini"))
	if err != nil {
		log.Fatal("Failed to initialize OpenAI client. Please check your API key configuration.")
	}

	// Generate the answer
	answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal("Failed to generate response from OpenAI. Please check your API key and network connection.")
	}

	return answer
}
