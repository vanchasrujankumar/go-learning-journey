package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Basic Select ===")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "two"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout")
	}

	fmt.Println("\n=== Select with Default ===")
	ch := make(chan int, 1)
	ch <- 42

	select {
	case val := <-ch:
		fmt.Println("Got value:", val)
	default:
		fmt.Println("No value available")
	}

	fmt.Println("\n=== Non-blocking Send ===")
	ch3 := make(chan int, 1)
	select {
	case ch3 <- 10:
		fmt.Println("Sent 10")
	default:
		fmt.Println("Channel full, skipping")
	}

	fmt.Println("\n=== Timeout Pattern ===")
	ch4 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch4 <- "result"
	}()

	select {
	case res := <-ch4:
		fmt.Println("Result:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out after 1 second")
	}
}
