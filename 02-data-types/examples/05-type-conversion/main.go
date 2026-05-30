package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Integer to Float
	intNum := 42
	floatNum := float64(intNum)
	fmt.Printf("Int %v to Float: %v (type: %T)\n", intNum, floatNum, floatNum)

	// Float to Integer (loses decimal part)
	floatVal := 3.99
	intVal := int(floatVal)
	fmt.Printf("Float %v to Int: %v (type: %T)\n", floatVal, intVal, intVal)

	// Integer to String
	numStr := strconv.Itoa(intNum)
	fmt.Printf("Int %v to String: %q (type: %T)\n", intNum, numStr, numStr)

	// String to Integer
	strNum := "123"
	parsedInt, _ := strconv.Atoi(strNum)
	fmt.Printf("String %q to Int: %v (type: %T)\n", strNum, parsedInt, parsedInt)

	// String to Float
	strFloat := "3.14"
	parsedFloat, _ := strconv.ParseFloat(strFloat, 64)
	fmt.Printf("String %q to Float: %v (type: %T)\n", strFloat, parsedFloat, parsedFloat)

	// Float to String
	floatStr := strconv.FormatFloat(floatVal, 'f', 2, 64)
	fmt.Printf("Float %v to String: %q (type: %T)\n", floatVal, floatStr, floatStr)
}
