package main

import "fmt"

func main() {
	// Using Print - no newline
	fmt.Print("Hello ")
	fmt.Print("World")
	fmt.Print("\n")

	// Using Println - adds newline automatically
	fmt.Println("This is a new line")
	fmt.Println("Your age:", 25)

	// Using Printf - formatted printing
	fmt.Printf("My name is %s and I am %d years old\n", "John", 25)
}
