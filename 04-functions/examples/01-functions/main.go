package main

import (
	"errors"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println("=== Basic Function ===")
	fmt.Printf("add(3, 5) = %d\n", add(3, 5))

	fmt.Println("\n=== Multiple Returns ===")
	result, err := divide(10, 3)
	if err == nil {
		fmt.Printf("divide(10, 3) = %.2f\n", result)
	}

	_, err = divide(5, 0)
	if err != nil {
		fmt.Printf("divide(5, 0) error: %s\n", err)
	}

	fmt.Println("\n=== Named Returns ===")
	a, b := split(27)
	fmt.Printf("split(27) = (%d, %d)\n", a, b)

	fmt.Println("\n=== Variadic Function ===")
	fmt.Printf("sum(1, 2) = %d\n", sum(1, 2))
	fmt.Printf("sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))

	nums := []int{10, 20, 30}
	fmt.Printf("sum(nums...) = %d\n", sum(nums...))
}
