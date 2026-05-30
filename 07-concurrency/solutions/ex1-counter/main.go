package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter: %d\n", counter)
}
