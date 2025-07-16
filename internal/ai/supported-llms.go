package ai

const defaultProvider = "gemini"

type LLMStrategy func(string) string

var supportedLLMs = map[string]LLMStrategy{}

func RegisterLLM(name string, strategy LLMStrategy) {
	supportedLLMs[name] = strategy
}

// GetSupportedLLMs return a map of supported llmProviders
func GetSupportedLLMs() map[string]string {
	providers := make(map[string]string)

	for provider := range supportedLLMs {
		providers[provider] = provider
	}

	return providers
}

func GetDefaultProvider() string {
	return defaultProvider
}
