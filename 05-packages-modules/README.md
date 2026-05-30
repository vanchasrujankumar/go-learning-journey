# Module 05: Packages & Modules

## 🎯 Learning Objectives
By the end of this module, you will:
- [ ] Understand Go's package system
- [ ] Create and import custom packages
- [ ] Understand exported vs unexported (public vs private) identifiers
- [ ] Initialize a Go module with `go mod init`
- [ ] Manage dependencies with `go mod tidy`, `go get`, `go.sum`
- [ ] Understand versioning and semantic import versioning
- [ ] Use third-party packages from different modules

**Estimated Time**: 3-4 hours  
**Difficulty**: ⭐⭐ Intermediate

---

## 📚 Key Concepts

### 1. Package Basics

```go
// Every Go file belongs to a package
package mathutils

// Exported (starts with capital letter) — visible outside the package
func Add(a, b int) int {
    return a + b
}

// Unexported (starts with lowercase) — private to the package
func helper() int {
    return 42
}
```

### 2. Imports

```go
import (
    "fmt"
    "math/rand"

    "github.com/someuser/somepackage"
    "example.com/user/project/mypackage"
)

// Named import (avoids collision)
import myfmt "mylib/fmt"

// Blank import (triggers init() side effects)
import _ "image/png"
```

### 3. Module Initialization

```bash
# Create a new module
go mod init github.com/username/projectname

# This creates go.mod:
module github.com/username/projectname

go 1.21
```

### 4. Dependency Management

```bash
# Add a dependency
go get github.com/google/uuid

# Tidy up (removes unused, adds missing)
go mod tidy

# Upgrade
go get -u github.com/google/uuid

# Use a specific version
go get github.com/google/uuid@v1.4.0
```

### 5. go.mod and go.sum

```go
// go.mod — module definition and dependencies
module example.com/myapp

go 1.21

require (
    github.com/google/uuid v1.4.0
    github.com/gorilla/mux v1.8.1
)

// go.sum — checksums for dependency verification (auto-managed)
```

### 6. Internal Packages

```go
// Files in an "internal" package are only importable by
// the parent tree — enforces encapsulation
// project/internal/auth/login.go
package auth
// Can only be imported by code inside project/
```

---

## 🔗 Essential Links

### Official Resources
- [Go Code Organization](https://golang.org/doc/code#Organization)
- [Go Modules Reference](https://golang.org/doc/modules/managing-dependencies)
- [Go Module Mode](https://golang.org/ref/mod)

### Tutorials
- [Go by Example: Packages](https://gobyexample.com/packages)
- [How to Write Go Code](https://golang.org/doc/code)
- [Developing Go Modules](https://golang.org/doc/modules/developing)

### Videos
- [Go Modules Explained](https://www.youtube.com/watch?v=vaQJW2lNmQI)
- [Understanding Go Packages](https://www.youtube.com/watch?v=8V0SgT3o4tY)

---

## 💻 Examples

### Example 1: Custom Package
See `examples/01-custom-package/`

### Example 2: Using External Modules
See `examples/02-external-modules/`

### Example 3: Export vs Unexported
See `examples/03-export-visibility/`

---

## ✋ Hands-On Exercises

### Exercise 1: Create a Custom Package
**Objective**: Learn to structure and import custom packages

**Instructions**:
1. Create `exercises/ex1-custom-package/` directory
2. Inside, create a `calculator/` subdirectory with `calculator.go`
3. The `calculator` package exports: `Add`, `Subtract`, `Multiply`, `Divide`
4. `Divide` returns an error for division by zero
5. The `main.go` in `exercises/ex1-custom-package/` imports and uses the package
6. Initialize a go module

**Solution**: See `solutions/ex1-custom-package/`

---

### Exercise 2: Use an External Module
**Objective**: Practice adding and using third-party dependencies

**Instructions**:
1. Create `exercises/ex2-use-external.go`
2. Initialize a Go module
3. Add the `github.com/google/uuid` dependency
4. Generate and print 5 unique UUIDs
5. Run `go mod tidy` to update go.sum

**Expected Output**:
```
UUID 1: 6ba7b810-9dad-11d1-80b4-00c04fd430c8
UUID 2: ...
```

**Solution**: See `solutions/ex2-use-external.go`

---

### Exercise 3: Export vs Unexported
**Objective**: Understand Go's visibility rules

**Instructions**:
1. Create `exercises/ex3-visibility/` directory
2. Inside, create a `greetings/` package with:
   - Exported function `Hello(name string) string`
   - Exported function `Hi(name string) string` (calls `hello` internally)
   - Unexported function `hello(name string) string`
   - Unexported variable `defaultGreeting`
3. In `main.go`, call `Hello()` and `Hi()` successfully
4. Try calling `hello()` directly — verify it causes a compilation error
5. Comment about why the error occurs

**Solution**: See `solutions/ex3-visibility/`

---

## 🧪 Testing Your Knowledge

### Quiz
1. **What makes a Go identifier exported (public)?**
   - It starts with a capital letter

2. **What does `go mod tidy` do?**
   - Adds missing module requirements and removes unused ones

3. **What's the purpose of `go.sum`?**
   - It contains expected cryptographic checksums for each dependency to verify integrity

4. **What is the `internal` package convention?**
   - Packages inside an `internal` directory can only be imported by code in the parent module tree

5. **How do you use a specific version of a module?**
   - `go get example.com/module@v1.2.3`

---

## 🚀 Next Steps

1. ✅ Complete all exercises in this module
2. ✅ Run all examples and understand the module system
3. ✅ Practice adding, upgrading, and removing dependencies
4. ✅ Move to Module 06: Structs & Interfaces

---

## 📝 Notes

Write your learning notes here:

```
- Package name = last element of import path (e.g., "math/rand" → rand)
- Capital letter = exported, lowercase = unexported
- go.mod is the module manifest; go.sum is the lock file
- Always run go mod tidy after changing dependencies
- internal/ prevents external consumers from importing your internal code
```

---

## 🔗 Additional Resources
- [Go Modules Wiki](https://github.com/golang/go/wiki/Modules)
- [Semantic Import Versioning](https://research.swtch.com/vgo-import)
- [Go Module Mirrors](https://proxy.golang.org/)

---

**Status**: Ready to Learn  
**Completion**: Mark as complete once you finish all exercises
