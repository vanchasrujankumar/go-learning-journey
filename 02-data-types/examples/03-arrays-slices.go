package main

import "fmt"

func main() {
	// Array - fixed size
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Array:", numbers)
	fmt.Println("Array length:", len(numbers))

	// Array with initialization
	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println("Names:", names)

	// Slice - dynamic size (no size specified)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)
	fmt.Println("Slice length:", len(slice))
	fmt.Println("Slice capacity:", cap(slice))

	// Append to slice
	slice = append(slice, 6, 7)
	fmt.Println("After append:", slice)

	// Slice of array
	arraySlice := numbers[1:4]
	fmt.Println("Slice from array [1:4]:", arraySlice)
}
