#!/bin/bash

# A script to install and configure the AI Commit Genie tool.

# Color codes for better output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}Starting AI Commit Genie setup...${NC}"

# Step 1: Check if Go is installed
echo "Checking for Go installation..."
if ! command -v go &> /dev/null
then
    echo -e "${RED}Error: Go is not installed or not in your PATH. Please install Go to continue.${NC}"
    exit 1
fi
echo "Go installation found."

# Step 2: Ask for the API Key first
echo -e "\n--- API Key Configuration ---"
read -p "Please enter your GEMINI_API_KEY: " API_KEY
if [ -z "$API_KEY" ]; then
    echo -e "${RED}API Key cannot be empty. Aborting installation.${NC}"
    exit 1
fi

# Step 3: Ask for language preference
echo -e "\n--- Language Configuration ---"
echo "Please select your preferred language for commit messages:"
echo "1) English"
echo "2) Português"
echo "3) Español"
echo "4) Français"
echo "5) Deutsch"
echo "6) 中文"
echo "7) 日本語"
echo "8) Русский"

# Default to English
LANG_CODE="en"
read -p "Enter your choice (1-8) [default: 1]: " LANG_CHOICE

case $LANG_CHOICE in
    2)
        LANG_CODE="pt"
        echo "Selected language: Português"
        ;;
    3)
        LANG_CODE="es"
        echo "Selected language: Español"
        ;;
    4)
        LANG_CODE="fr"
        echo "Selected language: Français"
        ;;
    5)
        LANG_CODE="de"
        echo "Selected language: Deutsch"
        ;;
    6)
        LANG_CODE="zh"
        echo "Selected language: 中文"
        ;;
    7)
        LANG_CODE="ja"
        echo "Selected language: 日本語"
        ;;
    8)
        LANG_CODE="ru"
        echo "Selected language: Русский"
        ;;
    *)
        echo "Selected language: English (default)"
        ;;
esac

# Step 4: Create the configuration directory and file
echo "Creating configuration file..."
CONFIG_DIR="$HOME/.config/git-ai-commit-genie"
mkdir -p "$CONFIG_DIR"
echo "GEMINI_API_KEY=${API_KEY}" > "$CONFIG_DIR/.env"
echo "AI_COMMIT_LANG=${LANG_CODE}" >> "$CONFIG_DIR/.env"
echo -e "Configuration saved to ${YELLOW}${CONFIG_DIR}/.env${NC}"

# Step 4: Install the Go program
echo -e "\nInstalling the 'git-ai-commit-genie' command..."
if ! go install ./cmd/git-ai-commit-genie; then
    echo -e "${RED}Installation failed. Please check the Go build errors above.${NC}"
    exit 1
fi
INSTALL_DIR="$(go env GOPATH)/bin"
echo -e "Command installed successfully to ${YELLOW}${INSTALL_DIR}${NC}"

# Step 5: Validate the PATH
echo -e "\nValidating PATH..."
if [[ ":$PATH:" != *":${INSTALL_DIR}:"* ]]; then
    echo -e "${YELLOW}IMPORTANT: Your PATH is not configured correctly.${NC}"
    echo "The installation directory ${YELLOW}${INSTALL_DIR}${NC} is not in your PATH."
    echo "To use 'git ai-commit-genie' from any directory, please add the following line to your shell profile (e.g., ~/.bashrc or ~/.zshrc):"
    echo -e "\n  ${GREEN}export PATH=\$PATH:${INSTALL_DIR}${NC}\n"
    echo "After adding it, please restart your terminal or run 'source ~/.bashrc'."
else
    echo -e "${GREEN}PATH validation successful! Your shell is ready.${NC}"
fi

echo -e "\n${GREEN}Setup complete!${NC}"
echo "You can now run 'git ai-commit-genie' inside any Git repository." 