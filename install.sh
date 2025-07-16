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

# Step 2: Ask for LLM provider preference
echo -e "\n--- LLM Provider Configuration ---"
echo "Please select your preferred AI provider:"
echo "1) Gemini (Google)"
echo "2) OpenAI (ChatGPT)"
echo "3) Anthropic (Claude)"

# Default to Gemini
LLM_PROVIDER="gemini"
read -p "Enter your choice (1-3) [default: 1]: " LLM_CHOICE

case $LLM_CHOICE in
    2)
        LLM_PROVIDER="openai"
        echo "Selected provider: OpenAI"
        ;;
    3)
        LLM_PROVIDER="anthropic"
        echo "Selected provider: Anthropic"
        ;;
    *)
        echo "Selected provider: Gemini (default)"
        ;;
esac

# Function to validate API key format
validate_api_key() {
    local key="$1"
    local provider="$2"
    
    # Remove any whitespace
    key=$(echo "$key" | tr -d '[:space:]')
    
    # Check if key is empty
    if [ -z "$key" ]; then
        return 1
    fi
    
    # Check for dangerous characters that could be used for injection
    if [[ "$key" =~ [;\&\|\`\$\(\)\<\>\\] ]]; then
        echo -e "${RED}Error: API key contains invalid characters.${NC}"
        return 1
    fi
    
    # Basic length validation (API keys are typically long)
    if [ ${#key} -lt 10 ]; then
        echo -e "${RED}Error: API key seems too short. Please check your key.${NC}"
        return 1
    fi
    
    # Provider-specific validation
    case "$provider" in
        "openai")
            if [[ ! "$key" =~ ^sk- ]]; then
                echo -e "${YELLOW}Warning: OpenAI API keys typically start with 'sk-'. Please verify your key.${NC}"
            fi
            ;;
        "anthropic")
            if [[ ! "$key" =~ ^sk- ]]; then
                echo -e "${YELLOW}Warning: Anthropic API keys typically start with 'sk-'. Please verify your key.${NC}"
            fi
            ;;
        "gemini")
            # Gemini keys have different format, just check basic structure
            if [ ${#key} -lt 20 ]; then
                echo -e "${YELLOW}Warning: Gemini API keys are typically longer. Please verify your key.${NC}"
            fi
            ;;
    esac
    
    return 0
}

# Step 3: Ask for API Key based on selected provider
echo -e "\n--- API Key Configuration ---"
if [ "$LLM_PROVIDER" = "openai" ]; then
    echo "You selected OpenAI. You'll need an OpenAI API key."
    echo "Get your API key from: https://platform.openai.com/api-keys"
    read -s -p "Please enter your OPENAI_API_KEY: " API_KEY
    echo  # New line after hidden input
    API_KEY_VAR="OPENAI_API_KEY"
elif [ "$LLM_PROVIDER" = "anthropic" ]; then
    echo "You selected Anthropic. You'll need an Anthropic API key."
    echo "Get your API key from: https://console.anthropic.com/settings/keys"
    read -s -p "Please enter your ANTHROPIC_API_KEY: " API_KEY
    echo  # New line after hidden input
    API_KEY_VAR="ANTHROPIC_API_KEY"
else
    echo "You selected Gemini. You'll need a Gemini API key."
    echo "Get your API key from: https://aistudio.google.com/app/apikey"
    read -s -p "Please enter your GEMINI_API_KEY: " API_KEY
    echo  # New line after hidden input
    API_KEY_VAR="GEMINI_API_KEY"
fi

# Validate the API key
if ! validate_api_key "$API_KEY" "$LLM_PROVIDER"; then
    echo -e "${RED}API Key validation failed. Aborting installation.${NC}"
    exit 1
fi

# Step 4: Ask for language preference
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

# Step 5: Create the configuration directory and file
echo "Creating configuration file..."
CONFIG_DIR="$HOME/.config/git-ai-commit-genie"
mkdir -p "$CONFIG_DIR"

# Create .env file with appropriate API key and configuration
if [ "$LLM_PROVIDER" = "openai" ]; then
    echo "OPENAI_API_KEY=${API_KEY}" > "$CONFIG_DIR/.env"
    echo "GEMINI_API_KEY=" >> "$CONFIG_DIR/.env"
    echo "ANTHROPIC_API_KEY=" >> "$CONFIG_DIR/.env"
elif [ "$LLM_PROVIDER" = "anthropic" ]; then
    echo "ANTHROPIC_API_KEY=${API_KEY}" > "$CONFIG_DIR/.env"
    echo "GEMINI_API_KEY=" >> "$CONFIG_DIR/.env"
    echo "OPENAI_API_KEY=" >> "$CONFIG_DIR/.env"
else
    echo "GEMINI_API_KEY=${API_KEY}" > "$CONFIG_DIR/.env"
    echo "OPENAI_API_KEY=" >> "$CONFIG_DIR/.env"
    echo "ANTHROPIC_API_KEY=" >> "$CONFIG_DIR/.env"
fi

echo "AI_COMMIT_LANG=${LANG_CODE}" >> "$CONFIG_DIR/.env"
echo "AI_COMMIT_PREFERRED_LLM_PROVIDER=${LLM_PROVIDER}" >> "$CONFIG_DIR/.env"
echo -e "Configuration saved to ${YELLOW}${CONFIG_DIR}/.env${NC}"

# Step 6: Install the Go program
echo -e "\nInstalling the 'git-ai-commit-genie' command..."
if ! go install ./cmd/git-ai-commit-genie; then
    echo -e "${RED}Installation failed. Please check the Go build errors above.${NC}"
    exit 1
fi
INSTALL_DIR="$(go env GOPATH)/bin"
echo -e "Command installed successfully to ${YELLOW}${INSTALL_DIR}${NC}"

# Step 7: Validate the PATH
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