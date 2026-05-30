package main

import "fmt"

func main() {
	// String declaration and operations
	str1 := "Hello"
	str2 := "World"

	// Concatenation
	message := str1 + " " + str2
	fmt.Println("Message:", message)

	// String length
	fmt.Println("Length of message:", len(message))

	// Access individual characters
	fmt.Println("First character:", string(message[0]))

	// Multi-line string (raw string)
	poem := `Go is fast,
Go is simple,
Go is powerful!`
	fmt.Println(poem)
}
