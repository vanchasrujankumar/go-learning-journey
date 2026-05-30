package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Generating 5 UUIDs:")

	for i := 1; i <= 5; i++ {
		id := uuid.New()
		fmt.Printf("UUID %d: %s\n", i, id.String())
	}
}
