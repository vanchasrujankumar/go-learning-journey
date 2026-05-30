package main

import (
	"fmt"

	"example.com/visibility/greetings"
)

func main() {
	fmt.Println(greetings.Hello("Alice"))
	fmt.Println(greetings.Hi("Bob"))

	fmt.Println("\nNote: calling greetings.hello() directly is impossible")
	fmt.Println("because 'hello' is unexported (lowercase first letter).")
	fmt.Println("Go enforces this at compile time — the code won't compile.")
}
