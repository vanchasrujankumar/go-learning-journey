package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Integer variable
	num := 42

	// Convert to float64
	floatNum := float64(num)
	fmt.Printf("Integer %d as float: %v\n", num, floatNum)

	// Convert to string
	strNum := strconv.Itoa(num)
	fmt.Printf("Integer %d as string: %q\n", num, strNum)

	// Parse string to integer
	parsedNum, err := strconv.Atoi("123")
	if err == nil {
		fmt.Printf("String '123' parsed to: %d\n", parsedNum)
	}

	// String to float
	strFloat := "3.14"
	parsedFloat, _ := strconv.ParseFloat(strFloat, 64)
	fmt.Printf("String %q parsed to float: %v\n", strFloat, parsedFloat)
}
