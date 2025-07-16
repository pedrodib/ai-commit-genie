package main

import (
	"ai-commit-genie/internal/ai"
	"ai-commit-genie/internal/config"
	"ai-commit-genie/internal/git"
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

const maxDiffLength = 10000 // Limit diff size to 10k characters to avoid large and expensive API requests

// determineLanguage centralizes the logic for determining the language to use
func determineLanguage(langFlag string) string {
	// Priority: command line flag > environment variable > default
	if langFlag != "" {
		if _, exists := ai.GetSupportedLanguages()[langFlag]; exists {
			return langFlag
		} else {
			fmt.Printf("Warning: Unsupported language code '%s'. Using configured language.\n", langFlag)
		}
	}
	
	// Get language from environment variable
	lang := os.Getenv("AI_COMMIT_LANG")
	if lang == "" {
		lang = "en" // Default to English
	}
	
	// Validate the language exists in supported languages
	if _, exists := ai.GetSupportedLanguages()[lang]; !exists {
		fmt.Printf("Warning: Configured language '%s' is not supported. Using English.\n", lang)
		return "en"
	}
	
	return lang
}

// askForConfirmation prompts the user with a question and waits for a y/n response.
// It reads from the provided io.Reader, making it testable.
func askForConfirmation(question string, reader io.Reader) bool {
	r := bufio.NewReader(reader)
	for {
		fmt.Print(question)
		response, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading your response: %v", err)
		}

		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		}
		if response == "n" || response == "no" {
			return false
		}
		// If the response is not y/n, the loop continues and asks again.
	}
}

func main() {
	// Parse command line flags
	var langFlag string
	var listLangs bool
	
	flag.StringVar(&langFlag, "lang", "", "Language for commit message (e.g., en, pt, es)")
	flag.BoolVar(&listLangs, "list-languages", false, "List all supported languages")
	flag.Parse()
	
	// If user requested to list languages
	if listLangs {
		fmt.Println("Supported languages:")
		for code, name := range ai.GetSupportedLanguages() {
			fmt.Printf("  %s: %s\n", code, name)
		}
		return
	}
	
	config.Load()

	// Determine the language to use (centralized logic)
	lang := determineLanguage(langFlag)
	
	// Get language name for display
	langName := ai.GetSupportedLanguages()[lang]

	diff, err := git.GetStagedDiff()
	if err != nil {
		log.Fatalf("Error getting diff: %v", err)
	}

	if diff == "" {
		fmt.Println("No staged changes found. Use 'git add' to prepare your files for the commit.")
		return
	}

	// Truncate the diff if it's too long
	if len(diff) > maxDiffLength {
		diff = diff[:maxDiffLength]
		fmt.Println("Warning: Diff is too long and has been truncated to 10000 characters.")
	}

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Generating commit message in %s...", langName)
	s.Start()

	commitMsg := ai.GenerateCommitMessage(diff, lang)

	s.Stop()

	fmt.Println("\n--- Suggested Commit Message ---")
	fmt.Println(commitMsg)
	fmt.Println("------------------------------------")

	if askForConfirmation("Do you want to commit with this message? (y/n): ", os.Stdin) {
		fmt.Println("Executing commit...")
		output, err := git.ExecuteCommit(commitMsg)
		if err != nil {
			log.Fatalf("Error executing git commit:\n%s", output)
		}
		fmt.Printf("Commit executed successfully!\n%s", output)
	} else {
		fmt.Println("Commit canceled.")
	}
}
