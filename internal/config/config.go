package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func Load() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Warning: could not find home directory: %v", err)
		return
	}

	// Validate home directory path to prevent path traversal
	if strings.Contains(home, "..") || strings.Contains(home, "~") {
		log.Fatal("Invalid home directory path detected")
	}

	configPath := filepath.Join(home, ".config", "git-ai-commit-genie", ".env")

	// Additional security: ensure the config path is within expected directory
	expectedPrefix := filepath.Join(home, ".config", "git-ai-commit-genie")
	if !strings.HasPrefix(configPath, expectedPrefix) {
		log.Fatal("Invalid configuration path detected")
	}

	// Check if file exists and is readable
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Error: config file not found in %s. Please run the installation script first.", configPath)
	}

	// godotenv.Load will load the .env file from the specified path.
	if err := godotenv.Load(configPath); err != nil {
		log.Fatalf("Error loading config file from %s. Error: %v", configPath, err)
	}
}
