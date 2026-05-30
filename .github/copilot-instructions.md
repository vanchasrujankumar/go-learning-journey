# Copilot Instructions for Go Learning Journey

This file helps GitHub Copilot provide better suggestions and completions for this Go learning project.

## Project Context

- **Project Name**: Go Learning Journey
- **Language**: Go (golang)
- **Go Version**: 1.21+
- **Purpose**: Comprehensive Go learning repository from beginner to intermediate
- **Structure**: Organized modules with examples, exercises, and solutions

## Code Style Guidelines

### Naming Conventions
- **Functions**: camelCase (e.g., `calculateSum`, `getUserName`)
- **Constants**: UPPER_CASE (e.g., `MAX_RETRIES`, `DEFAULT_PORT`)
- **Packages**: lowercase, single word when possible (e.g., `helpers`, `models`)
- **Exported identifiers**: PascalCase (e.g., `UserService`, `GetUser`)
- **Unexported identifiers**: camelCase (e.g., `userData`, `getConfig`)

### Code Formatting
- Indentation: Tabs (Go standard)
- Line length: Max 100 characters
- Always run `gofmt` before committing
- Use `goimports` for import organization

### Error Handling
```go
// Always handle errors explicitly
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}

// Use error wrapping for better debugging
if err != nil {
    return fmt.Errorf("in function X: %w", err)
}
```

### Comments
- Package comment: Appears at the top of package
- Exported functions/types: Start with function/type name
- Unexported functions: Start with description
- Keep comments concise and meaningful

```go
// Package handlers provides HTTP request handlers
package handlers

// ProcessUser handles user data processing
func ProcessUser(id string) error {
    // implementation
}
```

### Best Practices
1. Keep functions small and focused
2. Use interfaces for abstraction
3. Prefer composition over inheritance
4. Write tests alongside code
5. Use meaningful variable names
6. Avoid nested deep conditionals
7. Return errors explicitly
8. Use goroutines for concurrent operations

## Module Structure Pattern

Each module should follow this structure:
```
module-name/
├── README.md              # Learning objectives and concepts
├── examples/              # Working code examples
│   ├── 01-concept.go
│   ├── 02-concept.go
│   └── ...
├── solutions/             # Reference solutions for exercises
│   ├── ex1_solution.go
│   ├── ex2_solution.go
│   └── ...
└── exercises/             # Practice exercise stubs
    ├── exercise1.go
    ├── exercise2.go
    └── ...
```

## Suggested Copilot Completions

When suggesting code, Copilot should:
1. Follow Go conventions
2. Include error handling
3. Add meaningful comments
4. Use appropriate abstractions
5. Consider performance implications
6. Include examples where helpful

## Common Patterns to Suggest

### Error Handling
```go
func SafeOperation() error {
    // operation code
    return nil
}
```

### Function Signatures
```go
// Suggest exported functions with clear signatures
func ProcessData(input []string) ([]string, error) {
    // implementation
}
```

### Testing
```go
// Suggest table-driven tests
func TestFunction(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        // test cases
    }
}
```

## Topics for Enhanced Suggestions

### When suggesting concurrency:
- Use goroutines for I/O-bound operations
- Use channels for goroutine communication
- Consider using sync.WaitGroup for coordination
- Think about race conditions

### When suggesting HTTP handling:
- Use net/http standard library when learning
- Implement proper error responses
- Add appropriate content-type headers
- Consider middleware patterns

### When suggesting testing:
- Table-driven tests for better coverage
- Benchmark functions for performance testing
- Use sub-tests with t.Run()
- Mock external dependencies

## Exclusions

Do NOT suggest:
- External packages for learning basics
- Complex frameworks for beginner exercises
- Non-idiomatic Go patterns
- Third-party replacements for std library basics

## Performance Considerations

- Suggest efficient data structures
- Consider memory allocation
- Optimize loops when relevant
- Think about goroutine leaks
- Profile before optimizing

## Documentation

- Add examples in comments for complex functions
- Use clear, concise variable names
- Document edge cases
- Explain "why" not just "what"

---

**Note**: These instructions help provide contextually relevant suggestions. They're especially useful for:
- Code completion
- Refactoring suggestions
- Test generation
- Documentation help
- Pattern recommendations

Last Updated: May 30, 2026
