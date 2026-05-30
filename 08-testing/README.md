# Module 08: Testing

## 🎯 Learning Objectives
By the end of this module, you will:
- [ ] Write unit tests with the `testing` package
- [ ] Use table-driven tests for comprehensive coverage
- [ ] Write and run benchmark functions
- [ ] Measure test coverage with `go test -cover`
- [ ] Use subtests for hierarchical test organization
- [ ] Write example functions that double as documentation
- [ ] Use test helpers and setup/teardown patterns

**Estimated Time**: 3-4 hours  
**Difficulty**: ⭐⭐⭐ Advanced

---

## 📚 Key Concepts

### 1. Testing Basics

```go
// File: math_test.go (must end with _test.go)
package math

import "testing"

// Test function: must start with Test, takes *testing.T
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

### 2. Running Tests

```bash
# Run all tests in current package
go test

# Run with verbose output
go test -v

# Run specific test
go test -run TestAdd

# Run all tests in all subdirectories
go test ./...

# Run tests with coverage
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out  # Visual coverage report
```

### 3. Table-Driven Tests

```go
func TestDivide(t *testing.T) {
    tests := []struct {
        name     string
        a, b     float64
        expected float64
        wantErr  bool
    }{
        {"positive", 10, 2, 5, false},
        {"by zero", 10, 0, 0, true},
        {"negative", -10, 2, -5, false},
        {"fraction", 7, 2, 3.5, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := Divide(tt.a, tt.b)
            if tt.wantErr && err == nil {
                t.Error("expected error but got none")
            }
            if !tt.wantErr && result != tt.expected {
                t.Errorf("got %f, want %f", result, tt.expected)
            }
        })
    }
}
```

### 4. Subtests

```go
func TestMath(t *testing.T) {
    // Group related tests
    t.Run("addition", func(t *testing.T) {
        if Add(2, 3) != 5 {
            t.Error("addition failed")
        }
    })
    t.Run("subtraction", func(t *testing.T) {
        if Subtract(5, 3) != 2 {
            t.Error("subtraction failed")
        }
    })
}
```

### 5. Benchmarks

```go
func BenchmarkAdd(b *testing.B) {
    // b.N is adjusted by the framework for stable timing
    for i := 0; i < b.N; i++ {
        Add(100, 200)
    }
}

// Run benchmarks
// go test -bench=.
// go test -bench=Add -benchmem  # Includes memory allocation stats
```

### 6. Example Functions (Documentation + Testing)

```go
// Example function — serves as both documentation and a test
func ExampleAdd() {
    sum := Add(2, 3)
    fmt.Println(sum)
    // Output: 5
}
```

---

## 🔗 Essential Links

### Official Resources
- [testing Package Documentation](https://golang.org/pkg/testing/)
- [Go Blog: The Testing Package](https://blog.golang.org/subtests)
- [Go Testing: Subtests and Sub-benchmarks](https://golang.org/blog/subtests)

### Tutorials
- [Table-Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Go by Example: Testing](https://gobyexample.com/testing)
- [Go by Example: Benchmarks](https://gobyexample.com/testing-and-benchmarking)

### Videos
- [Testing in Go (YouTube)](https://www.youtube.com/watch?v=ndmB0bj7Z1A)
- [Go Testing Best Practices](https://www.youtube.com/watch?v=ys5G3hZHbzE)

---

## 💻 Examples

### Example 1: Basic Unit Tests
See `examples/01-basic-tests/`

### Example 2: Table-Driven Tests
See `examples/02-table-driven-tests/`

### Example 3: Benchmarks and Subtests
See `examples/03-benchmarks/`

---

## ✋ Hands-On Exercises

### Exercise 1: Write Unit Tests for Calculator
**Objective**: Practice writing basic unit tests

**Instructions**:
1. Create `exercises/ex1_calculator_test.go`
2. Test a `calculator` package with functions: `Add`, `Subtract`, `Multiply`, `Divide`
3. Write tests for each function (include edge cases like dividing by zero)
4. Run tests with `go test -v`
5. Check coverage with `go test -cover`

**Solution**: See `solutions/ex1_calculator_test.go`

---

### Exercise 2: Table-Driven Test
**Objective**: Master table-driven test patterns

**Instructions**:
1. Create `exercises/ex2_table_driven_test.go`
2. Test a function `IsPalindrome(s string) bool`
3. Use a table-driven approach with at least 6 test cases:
   - Empty string, single character, simple palindrome, case-sensitive, non-palindrome, palindrome with spaces
4. Use subtests (`t.Run`) for each case
5. Run with verbose output

**Solution**: See `solutions/ex2_table_driven_test.go`

---

### Exercise 3: Benchmark a Function
**Objective**: Write and run benchmarks

**Instructions**:
1. Create `exercises/ex3_benchmark_test.go`
2. Benchmark two implementations of a function (e.g., string concatenation with `+` vs `strings.Builder`)
3. For each benchmark:
   - Concatenate 1000 strings
   - Use `b.ResetTimer()` and `b.N`
4. Run benchmarks with `go test -bench=. -benchmem`
5. Compare results — which is faster? Which allocates less?

**Solution**: See `solutions/ex3_benchmark_test.go`

---

## 🧪 Testing Your Knowledge

### Quiz
1. **What must a test file name end with?**
   - `_test.go`

2. **What's the signature of a test function?**
   - `func TestXxx(t *testing.T)`

3. **How do you run only a specific test?**
   - `go test -run TestFunctionName`

4. **What's the difference between `t.Error()` and `t.Fatal()`?**
   - `Error` reports failure but continues the test; `Fatal` stops the test immediately

5. **What does `go test -bench=.` do?**
   - Runs all benchmark functions in the package — each gets `b.N` iterations for stable timing

---

## 🚀 Next Steps

1. ✅ Complete all exercises in this module
2. ✅ Run all examples and understand testing patterns
3. ✅ Practice table-driven tests and benchmarks
4. 🎉 Move on to the Project Phase: Build real applications!

---

## 📝 Notes

Write your learning notes here:

```
- Testing is a first-class citizen in Go toolchain
- Table-driven tests are the idiomatic Go pattern
- go test -cover shows coverage; go tool cover -html visualizes it
- Example functions double as executable documentation
- Benchmark with -benchmem to see allocation counts
```

---

## 🔗 Additional Resources
- [Testing Techniques (Go Wiki)](https://github.com/golang/go/wiki/TableDrivenTests)
- [Advanced Testing with Go](https://www.youtube.com/watch?v=8hQG7QlcLBk)
- [Go Test Architecture](https://golang.org/src/cmd/go/test.go)
- [Coverage in Go](https://blog.golang.org/cover)

---

**Status**: Ready to Learn  
**Completion**: Mark as complete once you finish all exercises
