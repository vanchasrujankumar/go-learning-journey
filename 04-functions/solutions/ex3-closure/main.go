package main

import "fmt"

func newCounter() (increment func() int, getValue func() int) {
	count := 0
	increment = func() int {
		count++
		return count
	}
	getValue = func() int {
		return count
	}
	return
}

func main() {
	inc, get := newCounter()

	for i := 0; i < 5; i++ {
		fmt.Printf("Count: %d\n", inc())
	}

	fmt.Printf("Final count: %d\n", get())
}
