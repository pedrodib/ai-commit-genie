# Contributing to AI Commit Genie 🧞

Thank you for your interest in contributing to AI Commit Genie! We welcome contributions from the community and are excited to see what you'll bring to the project.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [How to Contribute](#how-to-contribute)
- [Pull Request Process](#pull-request-process)
- [Coding Standards](#coding-standards)
- [Testing Guidelines](#testing-guidelines)
- [Issue Reporting](#issue-reporting)
- [Feature Requests](#feature-requests)

## 🤝 Code of Conduct

This project adheres to a code of conduct that we expect all contributors to follow:

- **Be respectful**: Treat everyone with respect and kindness
- **Be inclusive**: Welcome newcomers and help them get started
- **Be constructive**: Provide helpful feedback and suggestions
- **Be patient**: Remember that everyone has different skill levels

## 🚀 Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.20+**: [Download Go](https://golang.org/dl/)
- **Git**: [Install Git](https://git-scm.com/downloads)
- **An AI API Key**: Choose from [Gemini](https://aistudio.google.com/app/apikey), [OpenAI](https://platform.openai.com/api-keys), or [Anthropic](https://console.anthropic.com/settings/keys)

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/ai-commit-genie.git
   cd ai-commit-genie
   ```

## 🛠️ Development Setup

1. **Install dependencies**:
   ```bash
   go mod download
   ```

2. **Set up your environment**:
   ```bash
   cp .env-example ~/.config/git-ai-commit-genie/.env
   # Edit the .env file with your API keys
   ```

3. **Build the project**:
   ```bash
   go build -o git-ai-commit-genie ./cmd/git-ai-commit-genie
   ```

4. **Test the installation**:
   ```bash
   ./git-ai-commit-genie --list-languages
   ```

## 🤝 How to Contribute

### Types of Contributions

We welcome various types of contributions:

- 🐛 **Bug fixes**
- ✨ **New features**
- 📚 **Documentation improvements**
- 🧪 **Tests**
- 🌍 **Translations/Internationalization**
- 🎨 **UI/UX improvements**
- ⚡ **Performance optimizations**

### Areas for Contribution

- **New AI Providers**: Add support for additional LLM providers
- **Language Support**: Add new languages for commit messages
- **Security Enhancements**: Improve security measures
- **Performance**: Optimize API calls and response times
- **Testing**: Add unit tests and integration tests
- **Documentation**: Improve guides and examples

## 📝 Pull Request Process

### Before Submitting

1. **Create an issue** first to discuss major changes
2. **Check existing PRs** to avoid duplicates
3. **Test your changes** thoroughly
4. **Update documentation** if needed

### PR Checklist

- [ ] Code follows the project's coding standards
- [ ] All tests pass (when available)
- [ ] Documentation is updated
- [ ] Commit messages follow [Conventional Commits](https://conventionalcommits.org/)
- [ ] PR description clearly explains the changes

### PR Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Performance improvement
- [ ] Other (please describe)

## Testing
- [ ] Tested locally
- [ ] Added/updated tests
- [ ] All existing tests pass

## Screenshots (if applicable)
Add screenshots or GIFs demonstrating the changes
```

## 📏 Coding Standards

### Go Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Use `gofmt` to format your code
- Use meaningful variable and function names
- Add comments for exported functions and complex logic
- Keep functions small and focused

### Example:

```go
// GenerateCommitMessage creates an AI-generated commit message
// based on the provided diff and language preferences.
func GenerateCommitMessage(diff string, langCode string, llmProvider string) string {
    // Implementation here
}
```

### File Organization

- Keep related functionality in the same package
- Use clear package names that describe their purpose
- Separate concerns (AI logic, Git operations, configuration)

## 🧪 Testing Guidelines

### Writing Tests

- Write tests for new functionality
- Use table-driven tests when appropriate
- Mock external dependencies (AI APIs)
- Test both success and error cases

### Test Structure

```go
func TestGenerateCommitMessage(t *testing.T) {
    tests := []struct {
        name     string
        diff     string
        lang     string
        provider string
        want     string
    }{
        // Test cases here
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### Running Tests

```bash
go test ./...
go test -v ./internal/ai/
go test -cover ./...
```

## 🐛 Issue Reporting

### Bug Reports

When reporting bugs, please include:

- **Environment**: OS, Go version, AI provider
- **Steps to reproduce**: Clear, step-by-step instructions
- **Expected behavior**: What should happen
- **Actual behavior**: What actually happens
- **Error messages**: Full error output
- **Configuration**: Relevant config (without API keys)

### Bug Report Template

```markdown
**Environment:**
- OS: [e.g., macOS 13.0]
- Go version: [e.g., 1.21.0]
- AI Provider: [e.g., OpenAI]

**Steps to Reproduce:**
1. Run `git add .`
2. Run `git ai-commit-genie`
3. See error

**Expected Behavior:**
Should generate commit message

**Actual Behavior:**
Error: [paste error message]

**Additional Context:**
Any other relevant information
```

## 💡 Feature Requests

### Before Requesting

- Check if the feature already exists
- Search existing issues for similar requests
- Consider if it fits the project's scope

### Feature Request Template

```markdown
**Feature Description:**
Clear description of the proposed feature

**Use Case:**
Why is this feature needed? What problem does it solve?

**Proposed Solution:**
How do you envision this working?

**Alternatives Considered:**
What other approaches did you consider?

**Additional Context:**
Any other relevant information, mockups, or examples
```

## 🌍 Internationalization

### Adding New Languages

1. Add the language code and name to `languageNames` in `internal/ai/generator.go`
2. Update the installation script with the new language option
3. Test the new language with different AI providers
4. Update documentation

### Language Code Standards

Use ISO 639-1 codes:
- `en` for English
- `pt` for Portuguese
- `es` for Spanish
- etc.

## 📚 Documentation

### Documentation Standards

- Use clear, concise language
- Include code examples
- Keep README.md up to date
- Add inline comments for complex logic
- Update CHANGELOG.md for releases

### Documentation Structure

- **README.md**: Main project documentation
- **CONTRIBUTING.md**: This file
- **CHANGELOG.md**: Version history
- **docs/**: Additional documentation (if needed)

## 🏷️ Commit Message Guidelines

We follow [Conventional Commits](https://conventionalcommits.org/):

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

### Examples

```
feat(ai): add support for Claude 3.5 Sonnet model
fix(security): sanitize commit messages to prevent injection
docs(readme): add installation troubleshooting section
```

## 🎉 Recognition

Contributors will be recognized in:

- GitHub contributors list
- Release notes for significant contributions
- Special mentions for major features

## 📞 Getting Help

If you need help:

1. Check existing [issues](https://github.com/pedrodib/ai-commit-genie/issues)
2. Create a new issue with the "question" label
3. Be specific about what you're trying to achieve

## 📄 License

By contributing to AI Commit Genie, you agree that your contributions will be licensed under the [MIT License](LICENSE).

---

Thank you for contributing to AI Commit Genie! 🚀 Your contributions help make this tool better for developers worldwide.