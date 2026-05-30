package main

import (
	"fmt"

	"example.com/mypackage/calculator"
)

func main() {
	a, b := 15.0, 3.0

	fmt.Printf("%.0f + %.0f = %.0f\n", a, b, calculator.Add(a, b))
	fmt.Printf("%.0f - %.0f = %.0f\n", a, b, calculator.Subtract(a, b))
	fmt.Printf("%.0f * %.0f = %.0f\n", a, b, calculator.Multiply(a, b))

	result, err := calculator.Divide(a, b)
	if err == nil {
		fmt.Printf("%.0f / %.0f = %.0f\n", a, b, result)
	}

	_, err = calculator.Divide(a, 0)
	if err != nil {
		fmt.Printf("%.0f / 0 = error: %s\n", a, err)
	}
}
