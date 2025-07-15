# AI Commit Genie 🧞

**Version:** 0.1 (beta)

AI Commit Genie is a command-line tool that uses the power of Google's Gemini to automatically generate clear, descriptive, and conventional commit messages from your staged git changes.

Stop writing `git commit -m "fix"` and let the genie do the work for you!

## Description

This tool integrates directly with your Git workflow. By running `git ai-commit-genie`, it:
1.  Analyzes the staged diff of your changes.
2.  Sends the diff to the Gemini API to generate a high-quality commit message.
3.  Presents the suggested message for your approval.
4.  Commits the changes with the generated message if you approve.

The goal is to improve the quality of your commit history, speed up the versioning process, and standardize communication within your team.

## Requirements

Before you begin, ensure you have the following installed:
- **Go**: Version 1.20 or later.
- **Git**: The version control system itself.
- **A Gemini API Key**: You can get one from [Google AI Studio](https://aistudio.google.com/app/apikey).

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
    The script will guide you through the process, ask for your Gemini API key, and install the tool.

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

## Author

-   **Pedro Dib**

[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/pedrodib)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/pedro-dib)

## License & Contributing

This project is open source and licensed under the MIT License.

We welcome contributions from the community! Feel free to open issues, submit pull requests, or fork the repository to create your own versions. Let's make commit messages better, together!