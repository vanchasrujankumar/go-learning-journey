package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Generating 3 UUIDs using external module:")

	for i := 0; i < 3; i++ {
		id := uuid.New()
		fmt.Printf("UUID %d: %s\n", i+1, id.String())
	}
}
