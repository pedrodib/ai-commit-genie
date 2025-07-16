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

	// Validate API key exists
	if geminiApiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is not set")
	}

	// Create the context
	ctx := context.Background()

	// Create the Gemini client
	llm, err := googleai.New(ctx, googleai.WithAPIKey(geminiApiKey), googleai.WithDefaultModel("gemini-2.5-flash"))
	if err != nil {
		log.Fatal("Failed to initialize Gemini client. Please check your API key configuration.")
	}
	// Generate the answer
	answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal("Failed to generate response from Gemini. Please check your API key and network connection.")
	}

	return answer
}
