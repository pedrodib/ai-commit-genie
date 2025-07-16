package ai

import (
	"context"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func init() {
	RegisterLLM("gemini", GenerateWithGemini)
}

func GenerateWithGemini(prompt string) string {

	// Get the Gemini API key
	geminiApiKey := os.Getenv("GEMINI_API_KEY")

	// Create the context
	ctx := context.Background()

	// Create the Gemini client
	llm, err := googleai.New(ctx, googleai.WithAPIKey(geminiApiKey), googleai.WithDefaultModel("gemini-2.5-flash"))
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
