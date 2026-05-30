package main

import "fmt"

func main() {
	// Long form variable declarations
	var name string = "Alice"
	var age int = 25
	var height float64 = 5.8

	// Print with Printf
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.1f feet\n", height)

	// Short form declarations
	city := "New York"
	experience := 5

	fmt.Printf("\nCity: %s\n", city)
	fmt.Printf("Experience: %d years\n", experience)
}
