package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

type Employee struct {
	Person
	ID         int
	Department string
}

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func main() {
	fmt.Println("=== Struct Initialization ===")
	p1 := Person{"Alice", 30, "NYC"}
	p2 := Person{Name: "Bob", Age: 25}
	p3 := Person{Name: "Charlie"}

	fmt.Printf("p1: %+v\n", p1)
	fmt.Printf("p2: %+v\n", p2)
	fmt.Printf("p3: %+v\n", p3)

	fmt.Println("\n=== Struct Embedding ===")
	e := Employee{
		Person:     Person{Name: "Diana", Age: 28, City: "SF"},
		ID:         1001,
		Department: "Engineering",
	}
	fmt.Printf("Employee: %s (%d) from %s — %s\n", e.Name, e.Age, e.City, e.Department)

	fmt.Println("\n=== Struct Tags (JSON) ===")
	u := User{ID: 42, Email: "alice@example.com"}
	jsonBytes, _ := json.Marshal(u)
	fmt.Println("JSON:", string(jsonBytes))
}
