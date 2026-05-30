package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== WaitGroup ===")
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("All workers completed")

	fmt.Println("\n=== Mutex ===")
	var (
		counter int
		mu      sync.Mutex
	)

	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Printf("Counter: %d (expected: 1000)\n", counter)

	fmt.Println("\n=== RWMutex ===")
	var (
		value int
		rw    sync.RWMutex
	)

	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			rw.RLock()
			fmt.Printf("Reader %d: value = %d\n", id, value)
			rw.RUnlock()
		}(i)
	}
	wg2.Wait()
}
