package main

import (
	"fmt"
)

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go func() {
		for msg := range ping {
			fmt.Println(msg)
			pong <- "pong"
		}
	}()

	for round := 1; round <= 5; round++ {
		select {
		case ping <- fmt.Sprintf("Round %d: ping", round):
			msg := <-pong
			fmt.Println(msg)
		}
	}

	fmt.Println("Game over after 5 rounds")
}
