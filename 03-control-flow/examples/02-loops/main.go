package main

import "fmt"

func main() {
	fmt.Println("=== Classic For Loop ===")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	fmt.Println("\n=== Condition-Only Loop (while-style) ===")
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Printf("Sum doubled to: %d\n", sum)

	fmt.Println("\n=== Range Loop (slice) ===")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("fruits[%d] = %s\n", index, fruit)
	}

	fmt.Println("\n=== Range Loop (map) ===")
	ages := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	fmt.Println("\n=== Break and Continue ===")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 7 {
			break
		}
		fmt.Printf("Odd number: %d\n", i)
	}
}
