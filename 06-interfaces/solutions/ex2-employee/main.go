package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Birthday() {
	p.Age++
}

type Employee struct {
	Person
	ID         int
	Department string
}

func (e Employee) Describe() string {
	return fmt.Sprintf("Employee: %s (%d) in %s", e.Name, e.Age, e.Department)
}

func main() {
	e := Employee{
		Person:     Person{Name: "Charlie", Age: 42},
		ID:         2001,
		Department: "Engineering",
	}

	fmt.Println(e.Describe())
	fmt.Println("Happy Birthday!")
	e.Birthday()
	fmt.Println(e.Describe())
}
