package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Load() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Warning: could not find home directory: %v", err)
		return
	}

	configPath := filepath.Join(home, ".config", "git-ai-commit-genie", ".env")

	// godotenv.Load will load the .env file from the specified path.
	// If the file does not exist, it returns an error, which we treat as a warning.
	if err := godotenv.Load(configPath); err != nil {
		log.Printf("Error: config file not found in %s. Using system environment variables, if available. Error: %v", configPath, err)
		os.Exit(1)
	}
}
