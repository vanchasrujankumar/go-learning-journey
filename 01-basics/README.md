# Module 01: Basics & Setup

## 🎯 Learning Objectives
By the end of this module, you will:
- [ ] Understand Go's workspace structure
- [ ] Write and run your first Go program
- [ ] Understand basic syntax and comments
- [ ] Know how to use `go run` and `go build`
- [ ] Set up a Go project with modules

**Estimated Time**: 2-3 hours  
**Difficulty**: ⭐ Beginner

---

## 📚 Key Concepts

### 1. Go Workspace Structure
Go uses a simple workspace model:
```
~/go/
├── bin/        # Compiled binaries
├── pkg/        # Package objects
└── src/        # Source code
```

### 2. Your First Program
Every Go program needs a `main` package with a `main()` function.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### 3. Go Modules
Go 1.11+ uses modules for dependency management (replaces GOPATH):
```bash
go mod init github.com/username/projectname
```

---

## 🔗 Essential Links

### Official Resources
- [Go Install Guide](https://golang.org/doc/install)
- [Go Getting Started](https://golang.org/doc/code)
- [Workspace Guide](https://go.dev/doc/tutorial/getting-started)

### Tutorials
- [Tour of Go - Welcome](https://tour.golang.org/welcome/1)
- [Go by Example - Hello World](https://gobyexample.com/hello-world)

### Videos
- [Go in 100 Seconds](https://www.youtube.com/watch?v=446E-r0rZY8)
- [Golang Tutorial - Getting Started](https://www.youtube.com/watch?v=yyUHQIec83I&t=0s)

---

## 💻 Examples

### Example 1: Hello World
See `examples/01-hello-world.go`

### Example 2: Hello with Multiple Lines
See `examples/02-hello-formatted.go`

### Example 3: Using Variables
See `examples/03-variables.go`

---

## ✋ Hands-On Exercises

### Exercise 1: Your First Program
**Objective**: Create and run a simple Go program

**Instructions**:
1. Create a file named `my_first_program.go`
2. Write a program that prints "Welcome to Go!"
3. Run it with `go run my_first_program.go`

**Expected Output**:
```
Welcome to Go!
```

**Solution**: See `solutions/ex1_first_program.go`

---

### Exercise 2: Using Fmt Package
**Objective**: Learn different formatting functions

**Instructions**:
1. Create a file named `fmt_practice.go`
2. Use `fmt.Print()`, `fmt.Println()`, and `fmt.Printf()`
3. Print the following:
   - Your name on one line
   - Your age on the next line
   - A formatted message using Printf

**Solution**: See `solutions/ex2_fmt_practice.go`

---

### Exercise 3: Comments and Documentation
**Objective**: Learn Go commenting style

**Instructions**:
1. Create a file named `comments.go`
2. Add single-line comments (`//`)
3. Add a multi-line comment block (`/* */`)
4. Add a package comment above the package declaration

**Solution**: See `solutions/ex3_comments.go`

---

## 🧪 Testing Your Knowledge

### Quiz
1. What must every Go program have?
   - Answer: A `main` package with a `main()` function

2. What command runs Go code without creating a binary?
   - Answer: `go run`

3. What is the file that defines Go modules dependencies?
   - Answer: `go.mod`

4. How do you format strings in Go?
   - Answer: Using `fmt.Printf()` or string formatting

---

## 🚀 Next Steps

1. ✅ Complete all exercises in this module
2. ✅ Run all examples and modify them
3. ✅ Create your own variations
4. ✅ Move to Module 02: Data Types

---

## 📝 Notes

Write your learning notes here:

```
- Go is fast and simple
- Compilation is quick
- Single binary output
- Great for DevOps and system tools
```

---

## 🔗 Additional Resources
- [Go Code Organization](https://golang.org/doc/code)
- [Go Formatter](https://golang.org/cmd/gofmt/)
- [Go Playground](https://play.golang.org/)

---

**Status**: Ready to Learn  
**Completion**: Mark as complete once you finish all exercises
