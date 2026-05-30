package main

import "fmt"

func main() {
	// Declare a variable
	name := "Go Learner"
	version := 1.21

	// Print with Println
	fmt.Println("Name:", name)
	fmt.Println("Version:", version)

	// Print with Printf for formatted output
	fmt.Printf("Hello, %s! Learning Go version %.2f\n", name, version)
}
