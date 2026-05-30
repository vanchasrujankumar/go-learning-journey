# Module 06: Structs & Interfaces

## 🎯 Learning Objectives
By the end of this module, you will:
- [ ] Define and use structs with fields and tags
- [ ] Attach methods to struct types
- [ ] Embed structs for composition (Go's alternative to inheritance)
- [ ] Define and implement interfaces implicitly
- [ ] Use type assertions and type switches
- [ ] Work with the empty interface `interface{}`
- [ ] Understand the relationship between interfaces and nil

**Estimated Time**: 4-5 hours  
**Difficulty**: ⭐⭐ Intermediate

---

## 📚 Key Concepts

### 1. Structs

```go
// Define a struct
type Person struct {
    Name string
    Age  int
    City string
}

// Initialize structs
p1 := Person{"Alice", 30, "NYC"}
p2 := Person{Name: "Bob", Age: 25, City: "SF"}
p3 := Person{Name: "Charlie"}  // Age and City are zero values

// Access fields
fmt.Println(p1.Name)
p1.Age = 31

// Struct tags (used by serializers, validators, etc.)
type User struct {
    ID    int    `json:"id"`
    Email string `json:"email"`
}
```

### 2. Struct Embedding (Composition)

```go
type Address struct {
    Street string
    City   string
}

type Employee struct {
    Name    string
    Address // Embedded — fields promoted to Employee
}

e := Employee{Name: "Alice", Address: Address{Street: "123 Main", City: "NYC"}}
fmt.Println(e.City)    // Promoted field — accessed directly
fmt.Println(e.Street)  // Promoted field
```

### 3. Interfaces (Implicit Satisfaction)

Go interfaces are satisfied implicitly — no `implements` keyword.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

// Circle implements Shape implicitly by defining both methods
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Any function accepting a Shape can take Circle, Rectangle, etc.
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}
```

### 4. Type Assertions & Type Switches

```go
var i interface{} = "hello"

// Type assertion
s := i.(string)       // s is "hello"
n, ok := i.(int)      // ok is false, n is 0 (no panic)

// Type switch
switch v := i.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
default:
    fmt.Println("unknown type")
}
```

### 5. Empty Interface

```go
// Accept any type
func printAny(v interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", v, v)
}

printAny(42)
printAny("hello")
printAny(3.14)

// Go 1.18+: use `any` (alias for interface{})
func printAny2(v any) {
    fmt.Println(v)
}
```

---

## 🔗 Essential Links

### Official Resources
- [Tour of Go: Methods & Interfaces](https://tour.golang.org/methods/1)
- [Effective Go: Interfaces](https://golang.org/doc/effective_go#interfaces)
- [Go Spec: Interface types](https://golang.org/ref/spec#Interface_types)

### Tutorials
- [Go by Example: Structs](https://gobyexample.com/structs)
- [Go by Example: Interfaces](https://gobyexample.com/interfaces)
- [Go by Example: Embedding](https://gobyexample.com/embedding)
- [Go by Example: Type Assertions](https://gobyexample.com/type-assertions)

### Videos
- [Go Interfaces Explained](https://www.youtube.com/watch?v=Ej8__W2MsiA)
- [Understanding Go's Interface Model](https://www.youtube.com/watch?v=qJKQZKGZgf0)

---

## 💻 Examples

### Example 1: Structs and Methods
See `examples/01-structs.go`

### Example 2: Interfaces in Action
See `examples/02-interfaces.go`

### Example 3: Type Assertions and Embedding
See `examples/03-type-assertions.go`

---

## ✋ Hands-On Exercises

### Exercise 1: Shape Interface
**Objective**: Define and implement an interface

**Instructions**:
1. Create `exercises/ex1_shape_interface.go`
2. Define a `Shape` interface with `Area() float64` and `Perimeter() float64`
3. Implement it for `Circle` (radius field) and `Rectangle` (width, height fields)
4. Write a function `printShapeInfo(s Shape)` that prints both values
5. In main(), create instances of both and pass them to the function

**Expected Output**:
```
Circle: Area = 78.54, Perimeter = 31.42
Rectangle: Area = 50.00, Perimeter = 30.00
```

**Solution**: See `solutions/ex1_shape_interface.go`

---

### Exercise 2: Employee Struct with Methods
**Objective**: Practice structs, embedding, and methods

**Instructions**:
1. Create `exercises/ex2_employee.go`
2. Define a `Person` struct with Name, Age
3. Define an `Employee` struct that embeds Person and adds ID, Department
4. Add method `Describe() string` to Employee
5. Add method `Birthday()` to Person (pointer receiver, increments Age)
6. In main(), create an Employee, call Describe, call Birthday, call Describe again

**Expected Output**:
```
Employee: Charlie (42) in Engineering
Happy Birthday!
Employee: Charlie (43) in Engineering
```

**Solution**: See `solutions/ex2_employee.go`

---

### Exercise 3: Type Assertion Practice
**Objective**: Master type assertions and type switches

**Instructions**:
1. Create `exercises/ex3_type_assertion.go`
2. Write a function `describeType(v interface{})` that uses a type switch
3. Handle at least: int, float64, string, bool, []int, default
4. For each type, print a descriptive message
5. In main(), test with values of different types

**Expected Output**:
```
42 is an integer
3.14 is a float64
hello is a string of length 5
true is a boolean
[1 2 3] is a slice of ints with length 3
```

**Solution**: See `solutions/ex3_type_assertion.go`

---

## 🧪 Testing Your Knowledge

### Quiz
1. **How does Go achieve interface satisfaction (no `implements` keyword)?**
   - Implicitly: a type satisfies an interface if it implements all the interface's methods

2. **What is the difference between struct embedding and inheritance?**
   - Embedding is composition, not inheritance. There's no polymorphism or overloading — you compose behaviors by embedding types

3. **What does `s, ok := i.(string)` do?**
   - Type assertion: if `i` holds a string, `s` is the value and `ok` is true; otherwise `s` is zero and `ok` is false

4. **What's the difference between `interface{}` and `any`?**
   - None — `any` is an alias for `interface{}` introduced in Go 1.18

5. **Can a type implement multiple interfaces?**
   - Yes, a type can satisfy any number of interfaces as long as it implements all required methods

---

## 🚀 Next Steps

1. ✅ Complete all exercises in this module
2. ✅ Run all examples and understand implicit interfaces
3. ✅ Practice type assertions and type switches
4. ✅ Move to Module 07: Concurrency

---

## 📝 Notes

Write your learning notes here:

```
- "Accept interfaces, return structs" — common Go wisdom
- Small interfaces are better (1-2 methods)
- Embedding promotes fields and methods to the parent struct
- nil interface != nil pointer inside an interface
- Type assertion panics on mismatch without the ok idiom
```

---

## 🔗 Additional Resources
- [Go Data Structures: Interfaces](https://research.swtch.com/interfaces)
- [Google I/O 2012: Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Practical Go: Real world advice](https://dave.cheney.net/practical-go)

---

**Status**: Ready to Learn  
**Completion**: Mark as complete once you finish all exercises
