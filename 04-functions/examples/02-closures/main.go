package main

import "fmt"

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func makeGreeter(greeting string) func(string) string {
	return func(name string) string {
		return greeting + ", " + name + "!"
	}
}

func main() {
	fmt.Println("=== Closure: Counter ===")
	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	c2 := counter()
	fmt.Println("New counter starts at:", c2())

	fmt.Println("\n=== Closure: Greeter Factory ===")
	englishGreeter := makeGreeter("Hello")
	spanishGreeter := makeGreeter("Hola")

	fmt.Println(englishGreeter("Alice"))
	fmt.Println(spanishGreeter("Bob"))

	fmt.Println("\n=== Closure: Filter ===")
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Evens:", evens)

	greaterThan3 := filter(nums, func(n int) bool { return n > 3 })
	fmt.Println("> 3:", greaterThan3)
}

func filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, n := range nums {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}
