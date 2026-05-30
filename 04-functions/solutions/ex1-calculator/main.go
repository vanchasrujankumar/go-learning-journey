package main

import (
	"errors"
	"fmt"
)

func add(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	a, b := 10.0, 5.0

	if result, err := add(a, b); err == nil {
		fmt.Printf("%.0f + %.0f = %.0f\n", a, b, result)
	}
	if result, err := subtract(a, b); err == nil {
		fmt.Printf("%.0f - %.0f = %.0f\n", a, b, result)
	}
	if result, err := multiply(a, b); err == nil {
		fmt.Printf("%.0f * %.0f = %.0f\n", a, b, result)
	}
	if result, err := divide(a, b); err == nil {
		fmt.Printf("%.0f / %.0f = %.0f\n", a, b, result)
	}
	if result, err := divide(a, 0); err == nil {
		fmt.Printf("%.0f / 0 = %.0f\n", a, result)
	} else {
		fmt.Printf("%.0f / 0 = error: %s\n", a, err)
	}
}
