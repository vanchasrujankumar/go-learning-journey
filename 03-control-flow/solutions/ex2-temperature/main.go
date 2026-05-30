package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	var choice int
	var temp float64

	fmt.Println("Choose conversion:")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	fmt.Print("Enter temperature: ")
	fmt.Scan(&temp)

	switch choice {
	case 1:
		result := celsiusToFahrenheit(temp)
		fmt.Printf("%.2f°C = %.2f°F\n", temp, result)
	case 2:
		result := fahrenheitToCelsius(temp)
		fmt.Printf("%.2f°F = %.2f°C\n", temp, result)
	default:
		fmt.Println("Invalid choice")
	}
}
