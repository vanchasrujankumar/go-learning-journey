package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== Basic Goroutines ===")
	go say("Hello from goroutine")
	say("Hello from main")

	fmt.Println("\n=== Anonymous Goroutine ===")
	go func() {
		fmt.Println("Running in anonymous goroutine")
	}()
	time.Sleep(50 * time.Millisecond)

	fmt.Println("\n=== Unbuffered Channel ===")
	ch := make(chan string)
	go func() {
		ch <- "ping"
	}()
	msg := <-ch
	fmt.Println("Received:", msg)

	fmt.Println("\n=== Buffered Channel ===")
	bufCh := make(chan int, 3)
	bufCh <- 10
	bufCh <- 20
	bufCh <- 30
	close(bufCh)
	for val := range bufCh {
		fmt.Println("Buffered value:", val)
	}
}
