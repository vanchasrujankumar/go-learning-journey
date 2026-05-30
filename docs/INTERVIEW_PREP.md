# 🎯 Go Interview Preparation Guide

## 📋 Overview

Comprehensive interview prep covering easy to medium Go questions with pinpoint learning resources. Each topic links directly to documentation so you can master fundamentals fast.

### Difficulty Levels
| Level | Description | Typical For |
|-------|-------------|-------------|
| ⭐ Easy | Syntax, basic concepts | Junior, internship |
| ⭐⭐ Medium | Patterns, concurrency, testing | Mid-level (2-5 years) |

---

## 📚 Table of Contents

1. [Go Basics](#1-go-basics-⭐)
2. [Data Types & Collections](#2-data-types--collections-⭐)
3. [Control Flow & Functions](#3-control-flow--functions-⭐)
4. [Error Handling](#4-error-handling-⭐)
5. [Structs & Interfaces](#5-structs--interfaces-⭐⭐)
6. [Concurrency](#6-concurrency-⭐⭐)
7. [Packages & Modules](#7-packages--modules-⭐)
8. [Testing](#8-testing-⭐⭐)
9. [Common Interview Questions](#9-common-interview-questions-⭐⭐)
10. [Coding Challenges](#10-coding-challenges)
11. [Mock Interview Tips](#11-mock-interview-tips)
12. [Quick Reference Cheatsheet](#12-quick-reference-cheatsheet)

---

## 1. Go Basics ⭐

### Q1: What is Go and what makes it special?
**Go** (Golang) is a statically typed, compiled language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.

**Key characteristics:**
- **Fast compilation** – compiles to single binary
- **Concurrent by design** – goroutines, channels
- **Garbage collected** – automatic memory management
- **Simple syntax** – only 25 keywords
- **Static typing** – type safety at compile time
- **Great tooling** – built-in fmt, test, vet

**Resources:**
- 📖 [Go FAQ](https://golang.org/doc/faq)
- 📖 [Why Go?](https://golang.org/doc/)
- 📖 [Go Spec](https://golang.org/ref/spec)
- 📖 [Effective Go](https://golang.org/doc/effective_go)

---

### Q2: What are the basic data types in Go?
```go
bool        // true/false
string      // immutable text
int/int8/int16/int32/int64
uint/uint8/uint16/uint32/uint64
float32/float64
byte        // alias for uint8
rune        // alias for int32 (Unicode)
complex64/complex128
```

**Resources:**
- 📖 [Go Types](https://golang.org/ref/spec#Types)
- 📖 [Tour of Go - Basics](https://tour.golang.org/basics/11)

---

### Q3: Difference between `var`, `:=`, and `const`?

```go
var name string = "Go"       // Explicit type
var version = 1.21            // Type inferred
language := "Go"              // Short declaration (inside functions only)
const pi = 3.14159            // Immutable constant

// Multiple declaration
var (
    x, y int = 1, 2
    name string
)
```

| Feature | var | := | const |
|---------|-----|----|-------|
| Scope | Global/Local | Local only | Global/Local |
| Mutable | ✅ | ✅ | ❌ |
| Type inference | ✅ | ✅ | ✅ |
| Zero value | ✅ | ❌ (must assign) | N/A |

**Resources:**
- 📖 [Go Variables](https://golang.org/ref/spec#Variable_declarations)
- 📖 [Go Constants](https://golang.org/ref/spec#Constant_declarations)
- 📖 [Short Variable Declarations](https://golang.org/ref/spec#Short_variable_declarations)

---

### Q4: What are zero values in Go?

Variables declared without explicit initialization get **zero values**:

| Type | Zero Value |
|------|-----------|
| int, float | `0`, `0.0` |
| string | `""` |
| bool | `false` |
| pointer, slice, map, chan, interface | `nil` |
| struct | Zero values for all fields |

```go
var i int          // 0
var s string       // ""
var b bool         // false
var p *int         // nil
var m map[int]int  // nil
```

**Resources:**
- 📖 [Zero Value](https://golang.org/ref/spec#The_zero_value)
- 📖 [Tour of Go - Zero](https://tour.golang.org/basics/12)

---

### Q5: What is the difference between `Print`, `Println`, and `Printf`?

```go
fmt.Print("hello")          // No newline, no formatting
fmt.Println("hello")        // Adds newline at end, adds spaces between args
fmt.Printf("%s %d", "hello", 42)  // Formatted (like C printf)

// Format specifiers
%v  // Default format
%+v // Struct with field names
%#v // Go-syntax representation
%T  // Type
%d  // Integer
%s  // String
%f  // Float
%.2f // Float with 2 decimal places
```

**Resources:**
- 📖 [fmt Package](https://golang.org/pkg/fmt/)
- 📖 [Go by Example - Formatting](https://gobyexample.com/string-formatting)

---

## 2. Data Types & Collections ⭐

### Q6: Array vs Slice – what's the difference?

```go
// Array: fixed size
var arr [5]int = [5]int{1, 2, 3, 4, 5}

// Slice: dynamic size
var slc []int = []int{1, 2, 3}
slc = append(slc, 4)  // Can grow
```

| Feature | Array | Slice |
|---------|-------|-------|
| Size | Fixed at compile time | Dynamic |
| Pass by value | ✅ (copies entire array) | ❌ (passes reference to underlying array) |
| Can grow | ❌ | ✅ (via `append`) |
| Literal | `[3]int{1,2,3}` | `[]int{1,2,3}` |
| Use case | Known fixed size | Most Go code |

**Interview tip:** Slices are **always preferred** over arrays in Go unless you have a specific reason.

**Resources:**
- 📖 [Go Slices](https://golang.org/ref/spec#Slice_types)
- 📖 [Go Slices Blog Post](https://golang.org/blog/slices-intro)
- 📖 [Go by Example - Slices](https://gobyexample.com/slices)
- 📖 [Arrays vs Slices](https://www.golangprograms.com/go-language/arrays-vs-slices.html)

---

### Q7: How does `append` work internally?

```go
// append returns a new slice, may allocate new backing array
s := []int{1, 2, 3}
s = append(s, 4)    // length becomes 4

// Append multiple elements
s = append(s, 5, 6, 7)

// Append another slice (must spread)
s = append(s, []int{8, 9}...)

// Capacity doubling strategy
// Small slices: approximate double
// >1024 elements: ~25% growth
```

**Growth algorithm:**
```go
// Simplified append behavior
func append(slice []int, data ...int) []int {
    m := len(slice)
    n := m + len(data)
    if n > cap(slice) {
        newSlice := make([]int, n, growCap(slice, n))
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[:n]
    copy(slice[m:n], data)
    return slice
}
```

**Resources:**
- 📖 [Append Internals](https://golang.org/ref/spec#Appending_and_copying_slices)
- 📖 [Go Slice Tricks](https://github.com/golang/go/wiki/SliceTricks)

---

### Q8: How do maps work in Go?

```go
// Declaration and initialization
m := make(map[string]int)
m["key"] = 42

// Literal initialization
m := map[string]int{"a": 1, "b": 2}

// Check if key exists
value, exists := m["key"]
if exists {
    // use value
}

// Delete
delete(m, "key")

// Iteration (order is RANDOM!)
for key, value := range m {
    fmt.Println(key, value)
}
```

**Key points:**
- Maps are **reference types** (passed by reference)
- Map iteration order is **not guaranteed** (intentionally random)
- Maps are **not safe for concurrent use** (use `sync.RWMutex` or `sync.Map`)
- Reading a nil map returns zero value; **writing to nil map panics**

**Resources:**
- 📖 [Go Maps Blog](https://golang.org/blog/maps)
- 📖 [Go by Example - Maps](https://gobyexample.com/maps)
- 📖 [Map Iteration Order](https://stackoverflow.com/questions/11853396/in-go-why-are-iterations-over-maps-random)

---

### Q9: Strings – mutable or immutable?

**Strings in Go are immutable.** Once created, they cannot be changed.

```go
s := "hello"
// s[0] = 'H'  // COMPILE ERROR: cannot assign to s[0]

// To "modify" a string, create a new one
bytes := []byte(s)
bytes[0] = 'H'
s = string(bytes)  // "Hello"

// Efficient string building
var sb strings.Builder
sb.WriteString("hello")
sb.WriteString(" world")
result := sb.String()
```

**Resources:**
- 📖 [Go Strings Blog](https://golang.org/blog/strings)
- 📖 [strings.Builder](https://golang.org/pkg/strings/#Builder)
- 📖 [Go by Example - Strings](https://gobyexample.com/strings-and-runes)

---

## 3. Control Flow & Functions ⭐

### Q10: What are the different types of for loops?

```go
// 1. Traditional (C-style)
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// 2. While-style
sum := 0
for sum < 100 {
    sum += sum
}

// 3. Infinite loop (needs break)
for {
    // until break
    break
}

// 4. Range-based (most common in Go)
nums := []int{1, 2, 3}
for index, value := range nums {
    fmt.Println(index, value)
}

// Skip index
for _, value := range nums {
    fmt.Println(value)
}

// Iterate map
for key, value := range myMap {
    fmt.Println(key, value)
}

// Iterate string (runewise)
for index, rune := range "hello" {
    fmt.Printf("%d: %c\n", index, rune)
}
```

**Resources:**
- 📖 [Go For Loops](https://golang.org/ref/spec#For_statements)
- 📖 [Tour of Go - For](https://tour.golang.org/flowcontrol/1)
- 📖 [Tour of Go - Range](https://tour.golang.org/moretypes/16)

---

### Q11: Switch statements – what makes Go's switch unique?

```go
// 1. No break needed (automatic)
switch x {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")  // No fallthrough by default
default:
    fmt.Println("other")
}

// 2. Multiple values per case
switch x {
case 1, 2, 3:
    fmt.Println("small")
case 4, 5, 6:
    fmt.Println("medium")
}

// 3. Expressionless switch (like if-else chain)
switch {
case x < 10:
    fmt.Println("small")
case x < 100:
    fmt.Println("medium")
default:
    fmt.Println("large")
}

// 4. Type switch
var v interface{} = 42
switch v.(type) {
case int:
    fmt.Println("int")
case string:
    fmt.Println("string")
}
```

**Resources:**
- 📖 [Go Switch](https://golang.org/ref/spec#Switch_statements)
- 📖 [Tour of Go - Switch](https://tour.golang.org/flowcontrol/9)

---

### Q12: Multiple return values and named returns?

```go
// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Named return values (initialize to zero values)
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // "naked" return
}

// Use multiple returns
result, err := divide(10, 2)
if err != nil {
    log.Fatal(err)
}
```

**Resources:**
- 📖 [Go Functions](https://golang.org/ref/spec#Function_types)
- 📖 [Tour of Go - Multiple Results](https://tour.golang.org/basics/6)

---

### Q13: Variadic functions?

```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// Call with any number of args
sum(1, 2, 3)     // 6
sum(1, 2, 3, 4)  // 10

// Spread slice
nums := []int{1, 2, 3}
sum(nums...)  // 6
```

**Resources:**
- 📖 [Variadic Functions](https://golang.org/ref/spec#Function_types)
- 📖 [Go by Example - Variadic](https://gobyexample.com/variadic-functions)

---

### Q14: What is `defer` and how does it work?

```go
func readFile() error {
    f, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer f.Close()  // Executes when function returns
    
    // Or multiple defers (LIFO order)
    defer fmt.Println("third")
    defer fmt.Println("second")
    defer fmt.Println("first")
    // Output: first second third
}

// Common use cases:
defer mutex.Unlock()     // Unlock mutex
defer f.Close()          // Close file
defer db.Close()         // Close connection
defer wg.Done()          // Signal WaitGroup done

// Defer evaluates arguments immediately, not when deferred function runs
x := 1
defer fmt.Println(x)  // Prints "1", not current value of x
x = 2
```

**Resources:**
- 📖 [Defer Blog](https://golang.org/blog/defer-panic-and-recover)
- 📖 [Tour of Go - Defer](https://tour.golang.org/flowcontrol/12)

---

## 4. Error Handling ⭐

### Q15: How does Go handle errors?

Go treats errors as **values** (not exceptions). The standard pattern:

```go
// Define custom errors
var ErrNotFound = errors.New("not found")
var ErrPermission = fmt.Errorf("permission denied for user %s", user)

// Return error as last return value
func getUser(id int) (User, error) {
    if id < 0 {
        return User{}, ErrNotFound
    }
    return users[id], nil
}

// Handle error
user, err := getUser(5)
if err != nil {
    log.Printf("Error: %v", err)
    return
}

// Error wrapping (Go 1.13+)
if err != nil {
    return fmt.Errorf("getUser failed: %w", err)
}

// Unwrap errors
if errors.Is(err, ErrNotFound) {
    // specific handling
}
```

**Resources:**
- 📖 [Error Handling Blog](https://golang.org/blog/error-handling-and-go)
- 📖 [Working with Errors (Go 1.13)](https://golang.org/blog/go1.13-errors)
- 📖 [errors Package](https://golang.org/pkg/errors/)

---

### Q16: Panic vs Recover?

```go
// Panic: unexpected, program should stop
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")  // Unwinds stack
    }
    return a / b
}

// Recover: catch panic (only useful in deferred functions)
func safeDivide() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    divide(5, 0)  // Would panic, but recovered
}
```

**Never use panic for normal error handling.** Use it for truly unrecoverable situations.

**Resources:**
- 📖 [Defer, Panic, Recover](https://golang.org/blog/defer-panic-and-recover)
- 📖 [Effective Go - Panic](https://golang.org/doc/effective_go#panic)

---

## 5. Structs & Interfaces ⭐⭐

### Q17: How do structs work?

```go
// Define struct
type Person struct {
    Name    string
    Age     int
    Email   string
    Address Address  // Nested struct
}

// Tagged struct (used for serialization)
type User struct {
    ID    int    `json:"id" bson:"_id"`
    Name  string `json:"name" bson:"name"`
    Email string `json:"email" bson:"email"`
}

// Method on struct
func (p Person) Greet() string {
    return fmt.Sprintf("Hi, I'm %s", p.Name)
}

// Pointer receiver (can modify)
func (p *Person) Birthday() {
    p.Age++
}

// Initialize
p := Person{Name: "Alice", Age: 30}
p.Greet()  // "Hi, I'm Alice"
p.Birthday()  // Age becomes 31
```

**Resources:**
- 📖 [Tour of Go - Structs](https://tour.golang.org/moretypes/2)
- 📖 [Go by Example - Structs](https://gobyexample.com/structs)
- 📖 [Struct Tags](https://golang.org/pkg/reflect/#StructTag)

---

### Q18: Value receiver vs Pointer receiver?

```go
// Value receiver (copy of struct - cannot modify original)
func (u User) FullName() string {
    return u.Name
}

// Pointer receiver (original - can modify)
func (u *User) UpdateName(name string) {
    u.Name = name  // Modifies original
}

// When to use pointer receiver:
// 1. Method modifies the receiver
// 2. Large struct (avoid copying)
// 3. Consistency (if one method uses pointer, use for all)
```

**Resources:**
- 📖 [Tour of Go - Methods](https://tour.golang.org/methods/4)
- 📖 [Go FAQ - Value vs Pointer](https://golang.org/doc/faq#methods_on_values_or_pointers)

---

### Q19: How do interfaces work in Go?

```go
// Define interface
type Writer interface {
    Write([]byte) (int, error)
}

// Implement implicitly (no "implements" keyword)
type FileWriter struct{}

func (f FileWriter) Write(data []byte) (int, error) {
    // implementation
    return len(data), nil
}

// Can implement multiple interfaces
type Reader interface {
    Read([]byte) (int, error)
}

type FileReadWriter struct{}
func (f FileReadWriter) Write(data []byte) (int, error) { ... }
func (f FileReadWriter) Read(data []byte) (int, error)  { ... }

// Interface as parameter (most powerful pattern)
func ProcessData(w Writer) error {
    _, err := w.Write([]byte("data"))
    return err
}

// Empty interface (any type)
var anything interface{} = "hello"
var anything any = 42  // Go 1.18+ alias

// Type assertion
value, ok := anything.(string)
if ok {
    fmt.Println(value)
}
```

**Key insight:** "Bigger the interface, weaker the abstraction" – keep interfaces small.

**Resources:**
- 📖 [Tour of Go - Interfaces](https://tour.golang.org/methods/9)
- 📖 [Effective Go - Interfaces](https://golang.org/doc/effective_go#interfaces)
- 📖 [Go Interfaces](https://gobyexample.com/interfaces)

---

### Q20: Composition vs Inheritance in Go?

Go uses **composition** (not inheritance):

```go
// Composition: embed types
type Address struct {
    City    string
    Country string
}

type Person struct {
    Name    string
    Address          // Promoted fields: can access Person.City directly
}

// Instead of inheritance hierarchy
// ❌ Traditional OOP: Animal → Mammal → Dog
// ✅ Go composition: Dog has Animal characteristics

p := Person{Name: "Alice"}
p.City = "NYC"  // Direct access thanks to promotion

// Embedded interface
type ReadWriter interface {
    Reader
    Writer
}
```

**Resources:**
- 📖 [Go Embedding](https://golang.org/ref/spec#Struct_types)
- 📖 [Composition over Inheritance](https://stackoverflow.com/questions/50385340/does-go-have-inheritance)

---

## 6. Concurrency ⭐⭐

### Q21: Goroutines vs Threads?

```go
// Start a goroutine (lightweight thread)
go func() {
    fmt.Println("Running concurrently")
}()

// Compare goroutines to OS threads:
```

| Feature | Goroutine | OS Thread |
|---------|-----------|-----------|
| Stack size | ~4KB (grows/shrinks) | ~1MB (fixed) |
| Creation cost | ~1µs | ~1ms |
| Scheduling | Go runtime (M:N) | OS kernel |
| Context switch | ~100ns | ~1µs |
| Count per app | Millions | Thousands |

**Resources:**
- 📖 [Goroutines vs Threads](https://golang.org/doc/effective_go#goroutines)
- 📖 [Go Concurrency Video](https://www.youtube.com/watch?v=f6kdp27TYZs)
- 📖 [Tour of Go - Goroutines](https://tour.golang.org/concurrency/1)

---

### Q22: How do channels work?

```go
// Unbuffered channel (synchronous)
ch := make(chan int)
ch <- 42  // Blocks until someone receives
val := <-ch

// Buffered channel (async until full)
ch := make(chan int, 3)
ch <- 1  // Non-blocking (buffer available)
ch <- 2  // Non-blocking
ch <- 3  // Non-blocking
ch <- 4  // BLOCKS (buffer full)

// Close channel
close(ch)

// Range over channel
for val := range ch {
    fmt.Println(val)
}

// Select (handle multiple channels)
select {
case v := <-ch1:
    fmt.Println("from ch1:", v)
case v := <-ch2:
    fmt.Println("from ch2:", v)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no channel ready")
}
```

**Resources:**
- 📖 [Tour of Go - Channels](https://tour.golang.org/concurrency/2)
- 📖 [Go by Example - Channels](https://gobyexample.com/channels)
- 📖 [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)

---

### Q23: How to prevent data races?

```go
// Method 1: Mutex
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

// Method 2: RWMutex (readers don't block each other)
type Cache struct {
    mu   sync.RWMutex
    data map[string]interface{}
}

func (c *Cache) Get(key string) interface{} {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.data[key]
}

func (c *Cache) Set(key string, val interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = val
}

// Method 3: Atomic operations
var counter atomic.Int64
counter.Add(1)
fmt.Println(counter.Load())

// Method 4: Channels (share by communicating)
ch := make(chan int)
go func() {
    ch <- 42
}()
result := <-ch

// **ALWAYS** detect with race detector
// go test -race ./...
// go run -race main.go
```

**Resources:**
- 📖 [Go Memory Model](https://golang.org/ref/mem)
- 📖 [Race Detector](https://golang.org/doc/articles/race_detector.html)
- 📖 [sync Package](https://golang.org/pkg/sync/)
- 📖 [sync/atomic Package](https://golang.org/pkg/sync/atomic/)

---

### Q24: Select statement use cases?

```go
// 1. Multiplex channel reads
select {
case msg1 := <-ch1:
    fmt.Println("Received:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received:", msg2)
}

// 2. Timeout pattern
select {
case result := <-ch:
    fmt.Println(result)
case <-time.After(5 * time.Second):
    fmt.Println("timeout")
}

// 3. Non-blocking send/receive
select {
case ch <- value:
    fmt.Println("sent")
default:
    fmt.Println("channel full, skipping")
}

// 4. Done channel pattern
for {
    select {
    case <-done:
        return
    case v := <-workCh:
        process(v)
    }
}

// 5. Random selection (when multiple cases are ready)
select {
case ch1 <- 1:  // ~50% chance
case ch2 <- 2:  // ~50% chance
}
```

**Resources:**
- 📖 [Tour of Go - Select](https://tour.golang.org/concurrency/5)
- 📖 [Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw)

---

### Q25: WaitGroup usage?

```go
func main() {
    var wg sync.WaitGroup

    tasks := []string{"task1", "task2", "task3"}

    for _, task := range tasks {
        wg.Add(1)  // Increment counter
        go func(t string) {
            defer wg.Done()  // Decrement counter
            process(t)
        }(task)
    }

    wg.Wait()  // Block until counter reaches 0
    fmt.Println("All done")
}
```

**Resources:**
- 📖 [sync.WaitGroup](https://golang.org/pkg/sync/#WaitGroup)
- 📖 [Go by Example - WaitGroups](https://gobyexample.com/waitgroups)

---

### Q26: Worker pool pattern?

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        results <- job * 2  // Do work
    }
}

func main() {
    const numJobs = 100
    const numWorkers = 5

    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Start workers
    for w := 1; w <= numWorkers; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for r := 1; r <= numJobs; r++ {
        <-results
    }
}
```

**Resources:**
- 📖 [Worker Pools](https://gobyexample.com/worker-pools)
- 📖 [Go Concurrency Patterns Article](https://blog.golang.org/pipelines)

---

## 7. Packages & Modules ⭐

### Q27: How do Go packages work?

```go
// file: calc/math.go
package calc  // Package declaration

func Add(a, b int) int { return a + b }  // Exported (capital letter)
func sub(a, b int) int { return a - b }  // Unexported (lowercase)

// file: main.go
package main

import (
    "fmt"
    "calc"                     // Module-local package
    "github.com/gin-gonic/gin" // External module
)

func main() {
    fmt.Println(calc.Add(5, 3))
}
```

**Key rules:**
- Package name = last element of import path (e.g., `"database/sql"` → `sql`)
- **Exported** names start with capital letter
- **Unexported** names start with lowercase (package-private)
- No circular imports allowed

**Resources:**
- 📖 [Go Packages Guide](https://golang.org/doc/code#Organization)
- 📖 [Package Naming](https://golang.org/doc/effective_go#package-names)

---

### Q28: Go Modules (go.mod)?

```bash
# Initialize new module
go mod init github.com/user/project

# go.mod file created:
module github.com/user/project

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/redis/go-redis/v9 v9.3.0
)

# Common commands
go mod tidy     # Clean up dependencies
go mod download # Download all dependencies
go mod vendor   # Create vendor directory
go get pkg@v1.0 # Add/update dependency
```

**Resources:**
- 📖 [Go Modules Ref](https://golang.org/ref/mod)
- 📖 [Managing Dependencies](https://golang.org/doc/modules/managing-dependencies)

---

## 8. Testing ⭐⭐

### Q29: How to write unit tests in Go?

```go
// file: math.go
func Add(a, b int) int { return a + b }

// file: math_test.go
package calc

import "testing"

// Basic test
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2,3) = %d; want %d", result, expected)
    }
}

// Table-driven test (preferred pattern)
func TestAddTableDriven(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
        {"zero", 0, 5, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}

// Run: go test -v ./...
```

**Resources:**
- 📖 [testing Package](https://golang.org/pkg/testing/)
- 📖 [Go Testing Blog](https://golang.org/blog/add-a-test)
- 📖 [Table-Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests)
- 📖 [Go by Example - Testing](https://gobyexample.com/testing)

---

### Q30: Benchmark testing?

```go
// file: math_test.go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(i, i)
    }
}

// Run benchmark
// go test -bench=. -benchmem

// Example output:
// BenchmarkAdd   100000000   12.3 ns/op   0 B/op   0 allocs/op
```

**Resources:**
- 📖 [testing Package - Benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)
- 📖 [Benchmarks Blog](https://golang.org/blog/benchmarks)
- 📖 [Go by Example - Benchmarking](https://gobyexample.com/testing-and-benchmarking)

---

### Q31: Test coverage?

```bash
# Generate coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out  # View in browser

# Command line coverage
go test -v -cover ./...
```

**Resources:**
- 📖 [Coverage Blog](https://blog.golang.org/cover)
- 📖 [go test -cover](https://golang.org/pkg/testing/#hdr-Test_coverage)

---

## 9. Common Interview Questions ⭐⭐

### Design & Architecture

**Q32: How would you design a URL shortener in Go?**
- **Endpoints:** `POST /shorten`, `GET /{shortCode}`
- **Storage:** Redis (fast), MongoDB (persistent)
- **Components:**
  - Handler → Service → Repository pattern
  - Base62 encoding for short codes
  - Cache layer with LRU eviction
- **Concurrency:** Read/write mutex for counter

**Q33: Build a rate limiter in Go**
```go
// Token bucket pattern
type RateLimiter struct {
    mu       sync.Mutex
    tokens   float64
    maxTokens float64
    rate     float64
    lastRefill time.Time
}

func (rl *RateLimiter) Allow() bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    now := time.Now()
    elapsed := now.Sub(rl.lastRefill).Seconds()
    rl.tokens = math.Min(rl.maxTokens, rl.tokens + elapsed * rl.rate)
    rl.lastRefill = now
    
    if rl.tokens >= 1 {
        rl.tokens--
        return true
    }
    return false
}
```

**Resources:**
- 📖 [Rate Limiting Patterns](https://cloud.google.com/architecture/rate-limiting-strategies-patterns)
- 📖 [System Design Primer](https://github.com/donnemartin/system-design-primer)

---

### Best Practices

**Q34: What are Go best practices?**

| Practice | Why |
|----------|-----|
| `gofmt` formatting | Standardized code style |
| Meaningful names | `user.Get()` not `user.FetchUserData()` |
| Keep interfaces small | 1-3 methods max |
| Handle errors explicitly | Never ignore errors |
| Use `go vet` | Catches subtle bugs |
| Write tests | Table-driven tests |
| No global state | Makes testing hard |
| Use context | Timeouts, cancellation |

**Resources:**
- 📖 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- 📖 [Effective Go](https://golang.org/doc/effective_go)
- 📖 [Uber Go Guide](https://github.com/uber-go/guide)

---

### Tricky Questions

**Q35: What does this print?**
```go
defer func() { fmt.Println(recover()) }()
defer panic("b")
defer panic("a")
// panic() from deferred functions only catches most recent
// Output: b (last panic before recover)
```

**Q36: Slice append surprise**
```go
s := make([]int, 3)
s2 := append(s, 1)  // s2 has new backing array if capacity exceeded
s[0] = 10
fmt.Println(s[0], s2[0])  // s2[0] is still 0 if append grew capacity
```

**Q37: Nil vs empty slice**
```go
var s []int           // nil slice, len=0, cap=0
s = make([]int, 0)     // empty slice, len=0, cap=0
s = []int{}            // empty slice, len=0, cap=0

// nil check:
if len(s) == 0 {  // Works for all three
}
```

**Resources:**
- 📖 [Go Slice Tricks](https://github.com/golang/go/wiki/SliceTricks)
- 📖 [Common Go Mistakes](https://go.dev/wiki/CommonMistakes)
- 📖 [100 Go Mistakes Book](https://100go.co/)

---

## 10. Coding Challenges

### Easy Level

| Problem | Concepts | Link |
|---------|----------|------|
| Two Sum | Maps, iteration | [LeetCode](https://leetcode.com/problems/two-sum/) |
| FizzBuzz | Loops, conditionals | [LeetCode](https://leetcode.com/problems/fizz-buzz/) |
| Reverse String | Slices, loops | [LeetCode](https://leetcode.com/problems/reverse-string/) |
| Palindrome Check | Strings | [LeetCode](https://leetcode.com/problems/valid-palindrome/) |
| Max Subarray | Sliding window | [LeetCode](https://leetcode.com/problems/maximum-subarray/) |

### Medium Level

| Problem | Concepts | Link |
|---------|----------|------|
| Reverse Linked List | Pointers | [LeetCode](https://leetcode.com/problems/reverse-linked-list/) |
| Word Frequency | Maps | [LeetCode](https://leetcode.com/problems/word-frequency/) |
| Task Scheduler | Priority queue | [LeetCode](https://leetcode.com/problems/task-scheduler/) |
| LRU Cache | Doubly linked list | [LeetCode](https://leetcode.com/problems/lru-cache/) |
| Binary Search | Recursion | [LeetCode](https://leetcode.com/problems/binary-search/) |

### Concurrency Challenges

| Problem | Concepts | Link |
|---------|----------|------|
| Print in Order | Goroutines | [LeetCode](https://leetcode.com/problems/print-in-order/) |
| FooBar | Channels | [LeetCode](https://leetcode.com/problems/print-foobar-alternately/) |
| Dining Philosophers | Mutex, channels | [LeetCode](https://leetcode.com/problems/the-dining-philosophers/) |

**Practice Platforms:**
- 📖 [Exercism Go Track](https://exercism.org/tracks/go) – guided learning
- 📖 [LeetCode Go Tag](https://leetcode.com/tag/go/) – coding problems
- 📖 [HackerRank Go](https://www.hackerrank.com/domains/go) – challenges
- 📖 [Codewars Go](https://www.codewars.com/?language=go) – kata

---

## 11. Mock Interview Tips

### Before the Interview
- [ ] Review Go syntax in [Tour of Go](https://tour.golang.org/) (1 hour)
- [ ] Practice [LeetCode](https://leetcode.com/) problems in Go
- [ ] Know your resume projects inside out
- [ ] Prepare questions to ask the interviewer

### During the Interview
- **Think out loud** – explain your reasoning
- **Start simple** – get working solution first, then optimize
- **Handle errors** – always check return errors
- **Write idiomatic Go** – use `range`, `:=`, proper formatting
- **Consider concurrency** – when appropriate, mention goroutines
- **Test your logic** – walk through edge cases

### After Writing Code
1. Check for off-by-one errors
2. Ensure nil checks
3. Verify error handling
4. Run through example
5. Ask about concurrency if applicable

### Common Evaluations
- **Problem solving** – can you break down complex problems?
- **Code quality** – is it idiomatic Go?
- **Communication** – do you explain clearly?
- **Edge cases** – do you test boundary conditions?
- **Optimization** – can you improve time/space complexity?

### Questions to Ask the Interviewer
- What does the team's Go tech stack look like?
- How does the team handle code reviews?
- What kind of projects would I work on?
- What is the team size and structure?

**Resources:**
- 📖 [Go Interview Questions (GitHub)](https://github.com/EdixonAlberto/awesome-go-interview)
- 📖 [System Design Interviews](https://www.youtube.com/watch?v=ZgdS0EUiiHY)
- 📖 [Technical Interview Guide](https://github.com/jwasham/coding-interview-university)

---

## 12. Quick Reference Cheatsheet

### Syntax Quick Reference

```go
// Variable declarations
var x int = 5
y := 10
const Pi = 3.14

// Basic types
var b bool = true
var s string = "hello"
var i int = 42
var f float64 = 3.14

// Collections
arr := [3]int{1, 2, 3}
slc := []int{1, 2, 3}
mp := map[string]int{"a": 1}
st := struct{Name string}{"Alice"}

// Control flow
for i := 0; i < 10; i++ {}
for i, v := range slice {}
if condition {}
switch value {}
switch {}  // expression-less

// Functions
func add(a, b int) int { return a + b }
func divide(a, b int) (int, error) {}

// Concurrency
go func() {}
ch := make(chan int)
ch <- 1
v := <-ch
select {}
var wg sync.WaitGroup
wg.Add(1); wg.Done(); wg.Wait()

// Error handling
var ErrFoo = errors.New("foo")
if err != nil { return }
defer f.Close()
panic("unexpected")
recover()
```

### Key Go Commands

```bash
go version           # Version info
go run main.go       # Run without building
go build             # Compile to binary
go test ./...        # Run all tests
go test -v -race     # Verbose + race detection
go fmt ./...         # Format all Go files
go vet ./...         # Static analysis
go mod init          # Create new module
go mod tidy          # Clean dependencies
go get pkg@v1.0      # Add dependency
go doc fmt.Println   # View documentation
go tool cover -html  # Coverage report
```

### Common Gotchas

| Gotcha | Explanation |
|--------|-------------|
| `:=` not available outside function | Use `var` in package scope |
| Map iteration order is random | Don't depend on order |
| Slice passed by reference | Modifications affect caller |
| Strings immutable | Must convert to `[]byte` to modify |
| `break` not needed in `switch` | Automatic break (unique to Go) |
| `defer` evaluates args immediately | Args captured at defer time |
| Channel close only from sender side | Closing receiver causes panic |
| Nil map write panics | Always `make` maps before writing |
| Interface value can be nil | Check both `nil` and underlying type |

---

## 📚 Final Resources

### Must-Read
- [Effective Go](https://golang.org/doc/effective_go) – official best practices
- [Go FAQ](https://golang.org/doc/faq) – common questions answered
- [Go Tour](https://tour.golang.org/) – interactive learning
- [Go by Example](https://gobyexample.com/) – code snippets

### Deep Dives
- [Go Memory Model](https://golang.org/ref/mem)
- [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Advanced Go Testing](https://go.dev/blog/subtests)
- [Profiling Go Programs](https://go.dev/blog/pprof)

### Practice Platforms
- [Exercism Go](https://exercism.org/tracks/go)
- [LeetCode Go Tag](https://leetcode.com/tag/go/)
- [HackerRank Go](https://www.hackerrank.com/domains/go)
- [Codewars Go](https://www.codewars.com/?language=go)
- [Project Euler](https://projecteuler.net/)

### Mock Interviews
- [Pramp](https://www.pramp.com/) – free peer interviews
- [Interviewing.io](https://interviewing.io/) – anonymous practice
- [LeetCode Mock](https://leetcode.com/mockinterview/)

### Books
- [The Go Programming Language](https://www.gopl.io/) – definitive
- [Go in Action](https://www.manning.com/books/go-in-action)
- [Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/)
- [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941295/)

---

**Good luck with your Go interviews! 🎯**

Remember:
- ✅ Practice coding daily
- ✅ Review fundamentals weekly
- ✅ Build projects to apply knowledge
- ✅ Review system design concepts
- ✅ Know your Go standard library
- ✅ Stay calm and think out loud

---

**Last Updated**: May 30, 2026  
**Status**: ✅ Ready for Interview Prep  
**Difficulty**: Easy to Medium

---

*Consistent preparation beats last-minute cramming. Code daily! 🚀*
