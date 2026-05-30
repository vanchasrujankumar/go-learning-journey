package main

import (
	"fmt"

	"example.com/calc/calculator"
)

func main() {
	fmt.Println("Using custom package 'calculator':")
	fmt.Println("Add(10, 5) =", calculator.Add(10, 5))
	fmt.Println("Subtract(10, 5) =", calculator.Subtract(10, 5))
	fmt.Println("Multiply(10, 5) =", calculator.Multiply(10, 5))

	result, err := calculator.Divide(10, 5)
	if err == nil {
		fmt.Println("Divide(10, 5) =", result)
	}

	result, err = calculator.Divide(10, 0)
	if err != nil {
		fmt.Println("Divide(10, 0) error:", err)
	}
}
