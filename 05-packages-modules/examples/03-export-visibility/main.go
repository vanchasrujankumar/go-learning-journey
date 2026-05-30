package main

import (
	"fmt"

	"example.com/visibility/greetings"
)

func main() {
	fmt.Println(greetings.Hello("Alice"))
	fmt.Println(greetings.Hi("Bob"))

	// This would cause a compilation error:
	// fmt.Println(greetings.hello("Charlie"))
	// Because 'hello' is unexported (lowercase)

	fmt.Println("\nNote: greetings.hello() is unexported — try uncommenting it!")
}
