package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hi, I'm %s from %s", p.Name, p.City)
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	p := Person{Name: "Alice", Age: 30, City: "New York"}
	fmt.Println(p.Greet())
	p.HaveBirthday()
	fmt.Printf("Happy Birthday! Alice is now %d years old\n", p.Age)
}
