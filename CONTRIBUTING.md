# Contributing Guide

## Welcome to Go Learning Journey! 👋

This is a collaborative learning repository. Everyone is welcome to contribute improvements, fixes, and enhancements.

## How to Contribute

### 1. Fork & Clone

```bash
# Fork on GitHub first, then clone
git clone https://github.com/YOUR_USERNAME/go-learning-journey.git
cd go-learning-journey

# Add upstream
git remote add upstream https://github.com/SrujanKumar/go-learning-journey.git
```

### 2. Create a Branch

```bash
# Feature branch
git checkout -b feature/add-new-exercise

# Bug fix branch
git checkout -b fix/typo-in-readme

# Docs branch
git checkout -b docs/improve-module-03
```

### 3. Make Changes

Follow these guidelines:

#### For Code Examples
- Add comments explaining the concept
- Include error handling
- Keep it simple and focused
- Follow Go conventions

#### For Documentation
- Use clear, concise language
- Add hyperlinks to resources
- Include code examples where relevant
- Update table of contents

#### For Exercises
- Provide both the exercise template AND solution
- Include expected output
- Add hints for students
- Document difficulty level

### 4. Commit Message Conventions

Use semantic commit messages:

```bash
# Feature
git commit -m "feat: add module 03 examples for control flow"

# Fix
git commit -m "fix: correct typo in data types README"

# Documentation
git commit -m "docs: add GraphQL guide to architecture"

# Refactor
git commit -m "refactor: simplify error handling in examples"

# Chore
git commit -m "chore: update dependencies via renovate"

# Test
git commit -m "test: add unit tests for calculator exercise"
```

### 5. Push & Create PR

```bash
# Push to your fork
git push origin feature/add-new-exercise

# Create Pull Request on GitHub
# - Link related issues
# - Describe changes clearly
# - Add screenshots if relevant
```

## Code Style Guidelines

### Go Code

```go
// ✅ Good
package main

import (
    "fmt"
    "strings"
)

// ProcessText handles text processing
func ProcessText(input string) (string, error) {
    if input == "" {
        return "", fmt.Errorf("input cannot be empty")
    }
    
    return strings.ToUpper(input), nil
}

// ❌ Bad
package main
import "fmt"
func process(s string) string {
return strings.ToUpper(s)
}
```

### Documentation

```markdown
# ✅ Good Heading

Clear, concise explanation with examples.

## Subheading

- Bullet point 1
- Bullet point 2

### Code Example

\`\`\`go
// Go code example
\`\`\`

### Links

[Link Text](https://example.com)

# ❌ Bad

BAD HEADING!!!

Some explanation that is too long and unclear without proper structure.

Code example without formatting.
```

## Commit Types

| Type | Description |
|------|-------------|
| `feat` | New feature or module |
| `fix` | Bug fix or correction |
| `docs` | Documentation changes |
| `style` | Code formatting (no logic change) |
| `refactor` | Code reorganization |
| `perf` | Performance improvement |
| `test` | Test additions/changes |
| `chore` | Build, deps, config |
| `ci` | CI/CD configuration |

## PR Checklist

Before submitting:

- [ ] Code follows Go conventions
- [ ] Comments are clear and helpful
- [ ] Examples work and are tested
- [ ] Documentation is updated
- [ ] No merge conflicts
- [ ] Commit messages follow convention
- [ ] All changes are relevant to the PR

## Review Process

1. **Automated Checks**
   - Go fmt & vet pass
   - Tests pass
   - No conflicts

2. **Manual Review**
   - Code quality
   - Accuracy of content
   - Clarity of explanation

3. **Feedback & Iteration**
   - Address comments
   - Request re-review if needed
   - Iterate until approval

4. **Merge**
   - Squash commits (if needed)
   - Merge to main branch
   - Close related issues

## Common Contributions

### Adding a New Module

1. Create `NN-module-name/` directory
2. Add `README.md` with learning objectives
3. Create `examples/` with 3-5 working examples
4. Create `solutions/` with reference solutions
5. Update main `README.md` with module link
6. Commit: `feat: add module NN - module name`

### Fixing Typos

1. Edit the file
2. Commit: `fix: typo in module-name description`
3. PR title: `Fix typo in Module NN`

### Improving Documentation

1. Clarify content
2. Add better examples
3. Add missing links
4. Commit: `docs: improve module NN documentation`

### Adding Exercises

1. Create exercise template
2. Create solution
3. Add to module README
4. Commit: `feat: add exercises for module NN`

## Questions or Issues?

- Open an [Issue](https://github.com/SrujanKumar/go-learning-journey/issues)
- Use clear title and description
- Add relevant labels
- Include examples if applicable

## Recognition

Contributors are recognized in:
- `CONTRIBUTORS.md`
- GitHub insights
- Project documentation

## Code of Conduct

- Be respectful and inclusive
- Assume good intentions
- Welcome all experience levels
- Focus on constructive feedback
- Report violations to maintainers

## Resources

- [Git Guide](https://git-scm.com/book/en/v2)
- [GitHub Flow](https://guides.github.com/introduction/flow/)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Semantic Commits](https://www.conventionalcommits.org/)

---

**Thank you for contributing!** 🎉

Your improvements make this resource better for everyone learning Go.

Last Updated: May 30, 2026
