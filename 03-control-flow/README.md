# Module 03: Control Flow

## 🎯 Learning Objectives
By the end of this module, you will:
- [ ] Use `if/else` conditional statements
- [ ] Write all three forms of `for` loops (classic, condition-only, range)
- [ ] Use `switch` statements with and without expressions
- [ ] Understand `defer` and its stacking behavior
- [ ] Handle errors with `panic` and `recover`
- [ ] Control program flow with `break`, `continue`, and `goto`

**Estimated Time**: 2-3 hours  
**Difficulty**: ⭐ Beginner

---

## 📚 Key Concepts

### 1. If/Else Statements

```go
if x > 10 {
    fmt.Println("x is greater than 10")
} else if x > 5 {
    fmt.Println("x is greater than 5")
} else {
    fmt.Println("x is 5 or less")
}

// If with short statement
if result := calculate(); result > 0 {
    fmt.Println("Positive:", result)
}
```

### 2. For Loops (3 Types)

Go has only `for` — no `while` or `do-while`.

```go
// Type 1: Classic for (init; condition; post)
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Type 2: Condition-only (acts like while)
sum := 1
for sum < 100 {
    sum += sum
}

// Type 3: Range loop (iterate over collections)
nums := []int{10, 20, 30}
for index, value := range nums {
    fmt.Printf("nums[%d] = %d\n", index, value)
}

// Range with maps
for key, value := range map[string]int{"a": 1, "b": 2} {
    fmt.Printf("%s -> %d\n", key, value)
}

// Infinite loop
for {
    // break to exit
}
```

### 3. Switch Statements

```go
// Switch with expression
switch day {
case "Monday", "Tuesday":
    fmt.Println("Weekday")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Unknown")
}

// Switch without expression (clean if/else chains)
score := 85
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
case score >= 70:
    grade = "C"
default:
    grade = "F"
}
```

### 4. Defer

`defer` schedules a function call to run after the surrounding function returns.

```go
func readFile(filename string) {
    f, _ := os.Open(filename)
    defer f.Close()  // Runs when readFile returns
    // ... work with f ...
}

// Deferred calls are stacked (LIFO)
defer fmt.Println("third")
defer fmt.Println("second")
defer fmt.Println("first")
// Prints: first second third
```

### 5. Panic & Recover

```go
// Panic stops normal execution
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}

// Recover regains control after a panic
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from:", r)
    }
}()
```

---

## 🔗 Essential Links

### Official Resources
- [Tour of Go: Flow Control](https://tour.golang.org/flowcontrol/1)
- [Go Spec: For Statements](https://golang.org/ref/spec#For_statements)
- [Go Spec: Switch Statements](https://golang.org/ref/spec#Switch_statements)

### Tutorials
- [Go by Example: If/Else](https://gobyexample.com/if-else)
- [Go by Example: Switch](https://gobyexample.com/switch)
- [Go by Example: For](https://gobyexample.com/for)
- [Go by Example: Defer](https://gobyexample.com/defer)

### Articles & Videos
- [Defer, Panic, and Recover](https://golang.org/blog/defer-panic-and-recover)
- [Go Control Flow (YouTube)](https://www.youtube.com/watch?v=EpVGD3yVWek)

---

## 💻 Examples

### Example 1: If/Else with Grade Calculator
See `examples/01-if-else.go`

### Example 2: All Loop Types
See `examples/02-loops.go`

### Example 3: Switch and Defer
See `examples/03-switch-defer.go`

---

## ✋ Hands-On Exercises

### Exercise 1: FizzBuzz
**Objective**: Practice conditionals and loops

**Instructions**:
1. Create `exercises/ex1_fizzbuzz.go`
2. Write a program that prints numbers 1 to 100
3. For multiples of 3, print "Fizz" instead of the number
4. For multiples of 5, print "Buzz" instead
5. For multiples of both 3 and 5, print "FizzBuzz"

**Expected Output**:
```
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
...
```

**Solution**: See `solutions/ex1_fizzbuzz.go`

---

### Exercise 2: Temperature Converter
**Objective**: Master if/else chains and user input

**Instructions**:
1. Create `exercises/ex2_temperature.go`
2. Prompt the user to choose conversion direction (C→F or F→C)
3. Read the temperature value
4. Convert using formulas:
   - C to F: `(c * 9/5) + 32`
   - F to C: `(f - 32) * 5/9`
5. Print the result formatted to 2 decimal places

**Sample Interaction**:
```
Choose conversion:
1. Celsius to Fahrenheit
2. Fahrenheit to Celsius
Enter choice: 1
Enter temperature: 100
100.00°C = 212.00°F
```

**Solution**: See `solutions/ex2_temperature.go`

---

### Exercise 3: Menu with Switch
**Objective**: Practice switch statements

**Instructions**:
1. Create `exercises/ex3_menu.go`
2. Display a menu with at least 5 options
3. Use a switch statement to handle each option
4. Include a default case for invalid input
5. Loop until the user chooses to exit

**Sample Menu**:
```
=== Calculator Menu ===
1. Add
2. Subtract
3. Multiply
4. Divide
5. Exit
Enter choice: 1
Enter two numbers: 10 5
Result: 15
```

**Solution**: See `solutions/ex3_menu.go`

---

## 🧪 Testing Your Knowledge

### Quiz
1. **What are the three forms of `for` loops in Go?**
   - Classic (init; condition; post), condition-only (like while), and range

2. **How does `defer` execution order work with multiple deferred calls?**
   - Deferred calls are stacked LIFO (last-in, first-out)

3. **What happens when you omit the expression in a `switch` statement?**
   - It becomes a clean way to write if/else chains — each `case` is a boolean expression

4. **Can you use `continue` and `break` in Go?**
   - Yes, `break` exits the loop, `continue` skips to the next iteration

5. **What does `recover()` do?**
   - It regains control of a panicking goroutine and must be called inside a deferred function

---

## 🚀 Next Steps

1. ✅ Complete all exercises in this module
2. ✅ Run all examples and experiment with variations
3. ✅ Understand defer stacking and panic/recover patterns
4. ✅ Move to Module 04: Functions & Methods

---

## 📝 Notes

Write your learning notes here:

```
- Go only has `for`, no `while` or `do-while`
- `defer` is great for cleanup (closing files, unlocking mutexes)
- Switch cases don't fall through by default (no `break` needed)
- Panic is Go's "exception" — use sparingly, prefer errors
```

---

## 🔗 Additional Resources
- [Effective Go: Control Structures](https://golang.org/doc/effective_go#if)
- [Go by Example: All Control Flow](https://gobyexample.com/)
- [StackOverflow: Common Go Mistakes](https://stackoverflow.com/questions/41429974/common-go-mistakes)

---

**Status**: Ready to Learn  
**Completion**: Mark as complete once you finish all exercises
