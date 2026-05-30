# Module 04: Functions & Methods

## 🎯 Learning Objectives
By the end of this module, you will:
- [ ] Define and call functions with parameters
- [ ] Return multiple values from functions
- [ ] Use named return values
- [ ] Write variadic functions
- [ ] Create and use closures
- [ ] Define methods on types
- [ ] Use functions as first-class values

**Estimated Time**: 3-4 hours  
**Difficulty**: ⭐⭐ Intermediate

---

## 📚 Key Concepts

### 1. Function Basics

```go
// Simple function
func greet(name string) string {
    return "Hello, " + name
}

// Multiple parameters (same type shorthand)
func add(x, y int) int {
    return x + y
}
```

### 2. Multiple & Named Return Values

```go
// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // naked return
}
```

### 3. Variadic Functions

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// Usage
fmt.Println(sum(1, 2, 3))    // 6
fmt.Println(sum(1, 2, 3, 4)) // 10
```

### 4. Closures

```go
// A closure is a function that captures surrounding variables
func counter() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    c := counter()
    fmt.Println(c()) // 1
    fmt.Println(c()) // 2
    fmt.Println(c()) // 3
}
```

### 5. Methods

```go
type Rectangle struct {
    Width, Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver (can modify)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

### 6. First-Class Functions

```go
// Assign function to variable
fn := func(s string) {
    fmt.Println(s)
}
fn("hello")

// Pass function as argument
func apply(f func(int) int, x int) int {
    return f(x)
}

double := func(x int) int { return x * 2 }
fmt.Println(apply(double, 5)) // 10
```

---

## 🔗 Essential Links

### Official Resources
- [Tour of Go: Functions](https://tour.golang.org/basics/4)
- [Tour of Go: Methods](https://tour.golang.org/methods/1)
- [Go Spec: Function declarations](https://golang.org/ref/spec#Function_declarations)

### Tutorials
- [Go by Example: Functions](https://gobyexample.com/functions)
- [Go by Example: Closures](https://gobyexample.com/closures)
- [Go by Example: Methods](https://gobyexample.com/methods)
- [Go by Example: Variadic Functions](https://gobyexample.com/variadic-functions)

### Videos
- [Go Functions Explained](https://www.youtube.com/watch?v=ku0vY10FX7k)
- [Go Closures in Depth](https://www.youtube.com/watch?v=JsXlD02Q8Gg)

---

## 💻 Examples

### Example 1: Functions with Multiple Returns
See `examples/01-functions.go`

### Example 2: Closures and Variadic Functions
See `examples/02-closures.go`

### Example 3: Methods and First-Class Functions
See `examples/03-methods.go`

---

## ✋ Hands-On Exercises

### Exercise 1: Calculator with Functions
**Objective**: Practice functions with multiple return values

**Instructions**:
1. Create `exercises/ex1_calculator.go`
2. Define functions for add, subtract, multiply, divide
3. Each function takes two floats and returns (float64, error)
4. Division by zero should return an error
5. In main(), call each function and print results

**Expected Output**:
```
10 + 5 = 15
10 - 5 = 5
10 * 5 = 50
10 / 5 = 2
10 / 0 = error: division by zero
```

**Solution**: See `solutions/ex1_calculator.go`

---

### Exercise 2: Person Struct with Methods
**Objective**: Practice methods on custom types

**Instructions**:
1. Create `exercises/ex2_person.go`
2. Define a `Person` struct with Name, Age, City fields
3. Add a method `Greet() string` that returns "Hi, I'm [Name] from [City]"
4. Add a method `HaveBirthday()` that increments Age (pointer receiver!)
5. In main(), create a Person, call Greet, call HaveBirthday, and print the new age

**Expected Output**:
```
Hi, I'm Alice from New York
Happy Birthday! Alice is now 31 years old
```

**Solution**: See `solutions/ex2_person.go`

---

### Exercise 3: Closure Counter
**Objective**: Understand closure state capture

**Instructions**:
1. Create `exercises/ex3_closure.go`
2. Write a function `newCounter()` that returns two closures:
   - `increment()` — increments and returns the count
   - `getValue()` — returns the current count
3. The counter value must be private (captured in closure)
4. In main(), create a counter, increment 5 times, print the value

**Expected Output**:
```
Count: 1
Count: 2
Count: 3
Count: 4
Count: 5
Final count: 5
```

**Solution**: See `solutions/ex3_closure.go`

---

## 🧪 Testing Your Knowledge

### Quiz
1. **What's the difference between a function and a method in Go?**
   - A method is a function with a receiver argument (attached to a type)

2. **What are named return values and when might you use them?**
   - Named returns are pre-declared in the signature; they're initialized to zero values and a naked `return` returns them. Useful for documenting intent.

3. **How do you create a variadic function?**
   - Use `...` before the type of the last parameter, e.g., `func sum(nums ...int) int`

4. **What does a closure capture from its surrounding scope?**
   - Variables — captured by reference, not by value

5. **When should you use a pointer receiver vs a value receiver?**
   - Pointer: when you need to modify the receiver or avoid copying large structs. Value: when the struct is small and immutable.

---

## 🚀 Next Steps

1. ✅ Complete all exercises in this module
2. ✅ Run all examples and experiment with variations
3. ✅ Understand closures and pointer vs value receivers
4. ✅ Move to Module 05: Packages & Modules

---

## 📝 Notes

Write your learning notes here:

```
- Functions are first-class citizens in Go
- Multiple return values replace exceptions/tuples
- Closures are great for stateful functions
- Method sets: value receiver works for both value and pointer; pointer receiver only works for pointer
```

---

## 🔗 Additional Resources
- [Effective Go: Functions](https://golang.org/doc/effective_go#functions)
- [Go by Example: All Functions](https://gobyexample.com/)
- [Practical Go: Real world advice](https://dave.cheney.net/practical-go)

---

**Status**: Ready to Learn  
**Completion**: Mark as complete once you finish all exercises
