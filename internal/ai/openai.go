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

	// Create the context
	ctx := context.Background()

	// Create the OpenAi client
	llm, err := openai.New(openai.WithToken(openAiApiKey), openai.WithModel("gpt-4.1-mini"))
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
