package main

import "fmt"

func main() {
	// Variable declaration - long form
	var name string = "Go"
	var version float64 = 1.21
	var isAwesome bool = true

	// Variable declaration - short form (type inference)
	language := "Go"
	yearsOld := 12

	// Constants
	const pi = 3.14159

	fmt.Println("Name:", name)
	fmt.Println("Version:", version)
	fmt.Println("Is Awesome:", isAwesome)
	fmt.Println("Language:", language)
	fmt.Println("Years Old:", yearsOld)
	fmt.Println("Pi:", pi)
}
