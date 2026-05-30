package main

import "fmt"

func main() {
	var choice int
	var a, b float64

	for {
		fmt.Println("\n=== Calculator Menu ===")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		if choice == 5 {
			fmt.Println("Goodbye!")
			break
		}

		fmt.Print("Enter two numbers: ")
		fmt.Scan(&a, &b)

		switch choice {
		case 1:
			fmt.Printf("Result: %.2f\n", a+b)
		case 2:
			fmt.Printf("Result: %.2f\n", a-b)
		case 3:
			fmt.Printf("Result: %.2f\n", a*b)
		case 4:
			if b == 0 {
				fmt.Println("Error: division by zero")
			} else {
				fmt.Printf("Result: %.2f\n", a/b)
			}
		default:
			fmt.Println("Invalid choice, please try again")
		}
	}
}
