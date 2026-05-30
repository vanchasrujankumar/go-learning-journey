# Module 02: Data Types & Variables

## 🎯 Learning Objectives
By the end of this module, you will:
- [ ] Understand Go's type system
- [ ] Declare variables and constants
- [ ] Work with basic data types (int, float, string, bool)
- [ ] Manipulate slices and arrays
- [ ] Use maps for key-value storage
- [ ] Perform type conversions

**Estimated Time**: 3-4 hours  
**Difficulty**: ⭐ Beginner

---

## 📚 Key Concepts

### 1. Variables & Constants

```go
// Variable declaration - various ways
var name string = "Go"
var age int = 12
message := "Hello"  // Short declaration (type inferred)

// Constants
const version = 1.21
const (
    Red   = 0
    Green = 1
    Blue  = 2
)
```

### 2. Basic Data Types

| Type | Example | Description |
|------|---------|-------------|
| `int` | `42` | Integer number |
| `float64` | `3.14` | Floating point |
| `string` | `"hello"` | Text |
| `bool` | `true`, `false` | Boolean |
| `byte` | `65` | Single byte (0-255) |
| `rune` | `'A'` | Unicode character |

### 3. Collections

#### Arrays (fixed size)
```go
var numbers [5]int          // Array of 5 integers
var names [3]string = [3]string{"Alice", "Bob", "Charlie"}
```

#### Slices (dynamic)
```go
var nums []int                  // Slice (no size specified)
nums = append(nums, 1, 2, 3)   // Add elements
slice := []int{1, 2, 3, 4, 5}  // Initialize with values
```

#### Maps (key-value)
```go
person := make(map[string]string)
person["name"] = "Alice"
person["age"] = "30"

// Or initialize directly
scores := map[string]int{"Alice": 95, "Bob": 87}
```

### 4. Type Conversion

```go
var i int = 42
var f float64 = float64(i)      // int to float
var s string = strconv.Itoa(i)  // int to string
```

---

## 🔗 Essential Links

### Official Resources
- [Go Types](https://golang.org/ref/spec#Types)
- [Arrays, Slices and Maps](https://golang.org/doc/tutorial/create-module)
- [Go Spec - Variables](https://golang.org/ref/spec#Variable_declarations)

### Tutorials
- [Tour of Go - Basics](https://tour.golang.org/basics/1)
- [Go by Example - Variables](https://gobyexample.com/variables)
- [Go by Example - Arrays](https://gobyexample.com/arrays)
- [Go by Example - Slices](https://gobyexample.com/slices)
- [Go by Example - Maps](https://gobyexample.com/maps)

### Videos
- [Go Data Types Explained](https://www.youtube.com/watch?v=RnSdL1aXNl4)
- [Slices and Arrays in Go](https://www.youtube.com/watch?v=g0gN-Pz3SkY)

### Articles
- [Go Slices: usage and internals](https://golang.org/blog/slices-intro)
- [Go Maps](https://golang.org/blog/maps)

---

## 💻 Examples

### Example 1: Variables and Types
See `examples/01-variables.go`

### Example 2: Working with Strings
See `examples/02-strings.go`

### Example 3: Arrays and Slices
See `examples/03-arrays-slices.go`

### Example 4: Maps and Iteration
See `examples/04-maps.go`

### Example 5: Type Conversion
See `examples/05-type-conversion.go`

---

## ✋ Hands-On Exercises

### Exercise 1: Declare and Print Variables
**Objective**: Practice variable declaration and printing

**Instructions**:
1. Create a file `variables.go`
2. Declare variables for: name (string), age (int), height (float64)
3. Print them using Printf with appropriate format specifiers
4. Try both long and short declaration styles

**Solution**: See `solutions/ex1_variables.go`

---

### Exercise 2: Work with Arrays
**Objective**: Understand arrays and indexing

**Instructions**:
1. Create an array of 5 numbers
2. Assign values to each index
3. Print each element and calculate the sum
4. Print the length of the array

**Solution**: See `solutions/ex2_arrays.go`

---

### Exercise 3: Manipulate Slices
**Objective**: Learn slice operations

**Instructions**:
1. Create a slice from an array
2. Append new elements
3. Access elements and modify them
4. Print the slice length and capacity

**Solution**: See `solutions/ex3_slices.go`

---

### Exercise 4: Use Maps
**Objective**: Work with key-value storage

**Instructions**:
1. Create a map to store student grades (name -> score)
2. Add at least 3 students
3. Print the grades of each student
4. Update a grade
5. Check if a student exists in the map

**Solution**: See `solutions/ex4_maps.go`

---

### Exercise 5: Type Conversion Challenge
**Objective**: Practice converting between types

**Instructions**:
1. Create an integer variable
2. Convert it to float64 and print
3. Convert it to string and print
4. Parse a string to integer and print
5. Handle type assertion carefully

**Solution**: See `solutions/ex5_type_conversion.go`

---

## 🧪 Testing Your Knowledge

### Quiz
1. **What's the difference between arrays and slices?**
   - Arrays have fixed length, slices are dynamic

2. **How do you declare a variable in Go?**
   - Using `var`, `const`, or `:=` operator

3. **What's the zero value for an integer?**
   - `0`

4. **How do you add elements to a slice?**
   - Using `append()` function

5. **What's the syntax for maps?**
   - `map[keyType]valueType`

---

## 🚀 Practice Projects

### Mini Project: Simple To-Do System
Create a program that:
- Uses a slice to store todo items (strings)
- Allows adding new todos
- Prints all todos with numbers
- Tracks completion using a map

See `exercises/todo-mini-project.go`

---

## 📝 Common Mistakes to Avoid

```go
// ❌ WRONG - Can't use := with var keyword
var name := "Go"

// ✅ CORRECT
name := "Go"

// ❌ WRONG - Slices must be initialized before use (sometimes)
var s []int
s[0] = 5  // This will panic!

// ✅ CORRECT
s := make([]int, 5)  // Create with capacity
s[0] = 5

// ❌ WRONG - Type mismatch in operations
x := 5           // int
y := 3.14        // float64
z := x + y       // Compilation error!

// ✅ CORRECT
z := float64(x) + y
```

---

## 🎯 Summary

In this module you learned:
- Variable declaration (var, const, :=)
- Go's type system
- Collections (arrays, slices, maps)
- Type conversion
- Working with different data types

---

## 🔗 Additional Resources
- [Effective Go - Variables](https://golang.org/doc/effective_go#names)
- [Go Spec - Constants](https://golang.org/ref/spec#Constant_declarations)
- [Standard Library - strconv](https://golang.org/pkg/strconv/)

---

**Status**: Ready to Learn  
**Completion**: Mark as complete once you finish all exercises
