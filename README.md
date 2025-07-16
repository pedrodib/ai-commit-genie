# AI Commit Genie 🧞

**Version:** 0.3 (beta)

AI Commit Genie is a command-line tool that uses the power of AI (Google Gemini, OpenAI, or Anthropic) to automatically generate clear, descriptive, and conventional commit messages from your staged git changes.

Stop writing `git commit -m "fix"` and let the genie do the work for you!

## Description

This tool integrates directly with your Git workflow. By running `git ai-commit-genie`, it:
1.  Analyzes the staged diff of your changes.
2.  Sends the diff to your chosen AI provider (Gemini, OpenAI, or Anthropic) to generate a high-quality commit message.
3.  Presents the suggested message for your approval.
4.  Commits the changes with the generated message if you approve.

The goal is to improve the quality of your commit history, speed up the versioning process, and standardize communication within your team.

## Requirements

Before you begin, ensure you have the following installed:
- **Go**: Version 1.20 or later.
- **Git**: The version control system itself.
- **An AI API Key**: Choose one of the supported providers:
  - **Gemini API Key**: Get one from [Google AI Studio](https://aistudio.google.com/app/apikey)
  - **OpenAI API Key**: Get one from [OpenAI Platform](https://platform.openai.com/api-keys)
  - **Anthropic API Key**: Get one from [Anthropic Console](https://console.anthropic.com/settings/keys)

## Installation

We've made installation simple with an interactive setup script.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/[YourGitHubUsername]/ai-commit-genie.git
    cd ai-commit-genie
    ```

2.  **Run the installation script:**
    ```bash
    ./install.sh
    ```
    The script will guide you through the process, ask you to choose your preferred AI provider, request the appropriate API key, and install the tool.

### ❗ Important: Setting up your PATH

The installation script will install the `git-ai-commit-genie` executable into Go's binary directory (usually `~/go/bin`). For your terminal to find and use this command from any directory, this location **must** be in your system's `PATH`.

The script will automatically detect if your `PATH` is configured correctly. If it's not, it will instruct you to add the following line to your shell's configuration file (e.g., `~/.bashrc`, `~/.zshrc`, or `~/.profile`):

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

After adding this line, remember to **restart your terminal** or run `source ~/.bashrc` (or the equivalent for your shell) for the changes to take effect.

## How to Use

The workflow is designed to be as simple as possible:

1.  **Stage your changes** as you normally would:
    ```bash
    git add .
    ```

2.  **Run the genie:**
    ```bash
    git ai-commit-genie
    ```

3.  **Approve the commit:**
    The tool will show you a suggested commit message and ask for your confirmation.
    ```
    --- Suggested Commit Message ---
    feat: implement interactive confirmation for commits

    This change introduces an interactive prompt that asks the user for
    confirmation before executing the git commit.

    ------------------------------------
    Do you want to commit with this message? (y/n): y
    ```
    Type `y` and press Enter to proceed. The commit will be made. If you type `n`, the commit will be canceled.

### AI Provider Support

AI Commit Genie supports multiple AI providers for generating commit messages:

#### Supported Providers:
- **Gemini (Google)**: Uses Gemini-2.5-flash, fast and efficient, great for most use cases
- **OpenAI (ChatGPT)**: Uses GPT-4.1-mini model for high-quality responses
- **Anthropic (Claude)**: Uses Claude-3.5-haiku-latest model for thoughtful and nuanced responses

#### Usage:

1. **Set your preferred provider during installation**
   The installation script will prompt you to select your preferred AI provider.

2. **Specify a provider for a single commit:**
   ```bash
   git ai-commit-genie --llm-provider openai     # Use OpenAI for this commit
   git ai-commit-genie --llm-provider gemini     # Use Gemini for this commit
   git ai-commit-genie --llm-provider anthropic  # Use Anthropic for this commit
   ```

3. **Manually edit your provider preference:**
   You can edit the `AI_COMMIT_PREFERRED_LLM_PROVIDER` variable in `~/.config/git-ai-commit-genie/.env`

### Language Support

AI Commit Genie supports multiple languages for generating commit messages. You can:

1. **Set your preferred language during installation**
   The installation script will prompt you to select your preferred language, which will be saved as `AI_COMMIT_LANG` in your configuration file.

2. **Specify a language for a single commit:**
   ```bash
   git ai-commit-genie --lang pt  # Generate commit message in Portuguese
   ```

3. **List all supported languages:**
   ```bash
   git ai-commit-genie --list-languages
   ```

4. **Manually edit your language preference:**
   You can edit the `AI_COMMIT_LANG` variable in `~/.config/git-ai-commit-genie/.env`

Currently supported languages:
- English (en)
- Português (pt)
- Español (es)
- Français (fr)
- Deutsch (de)
- 中文 (zh)
- 日本語 (ja)
- Русский (ru)

### Advanced Usage

You can combine multiple options for maximum flexibility:

```bash
# Generate commit message in Portuguese using OpenAI
git ai-commit-genie --lang pt --llm-provider openai

# Generate commit message in Spanish using Gemini
git ai-commit-genie --lang es --llm-provider gemini

# Generate commit message in French using Anthropic
git ai-commit-genie --lang fr --llm-provider anthropic
```

## Configuration

Your configuration is stored in `~/.config/git-ai-commit-genie/.env`. Here's what each setting does:

```bash
# API Keys - You only need the key for your preferred provider
GEMINI_API_KEY=your_gemini_key_here
OPENAI_API_KEY=your_openai_key_here
ANTHROPIC_API_KEY=your_anthropic_key_here

# Language for commit messages (en, pt, es, fr, de, zh, ja, ru)
AI_COMMIT_LANG=en

# Default AI provider (gemini, openai, anthropic)
AI_COMMIT_PREFERRED_LLM_PROVIDER=gemini
```

### Switching Providers

To switch between Gemini, OpenAI, and Anthropic:

1. **Get the appropriate API key** from the provider's website
2. **Edit your configuration file** at `~/.config/git-ai-commit-genie/.env`
3. **Add your new API key** and update the `AI_COMMIT_PREFERRED_LLM_PROVIDER` setting
4. **Test the new provider** with: `git ai-commit-genie --llm-provider [provider_name]`

Available providers: `gemini`, `openai`, `anthropic`

## Troubleshooting

### Common Issues

**"API Key not found" error:**
- Make sure you have the correct API key for your chosen provider
- Verify the key is properly set in your `.env` file
- Check that there are no extra spaces or quotes around the key

**"Unsupported provider" warning:**
- Ensure you're using `gemini`, `openai`, or `anthropic` as the provider name
- Check your configuration file for typos

**"No staged changes found":**
- Run `git add .` to stage your changes before using the tool
- Make sure you're in a Git repository

### Getting Help

If you encounter issues:
1. Check your configuration file: `cat ~/.config/git-ai-commit-genie/.env`
2. Test with explicit parameters: `git ai-commit-genie --lang en --llm-provider gemini`
3. Verify your API key is valid by testing it directly with the provider's API

## Author

-   **Pedro Dib**

[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/pedrodib)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/pedro-dib)

## License & Contributing

This project is open source and licensed under the MIT License.

We welcome contributions from the community! Feel free to open issues, submit pull requests, or fork the repository to create your own versions. Let's make commit messages better, together!