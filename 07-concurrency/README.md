# Module 07: Concurrency

## 🎯 Learning Objectives
By the end of this module, you will:
- [ ] Launch goroutines with the `go` keyword
- [ ] Communicate between goroutines using channels
- [ ] Differentiate between buffered and unbuffered channels
- [ ] Use `select` to work with multiple channels
- [ ] Synchronize goroutines with `sync.WaitGroup`
- [ ] Protect shared state with `sync.Mutex`
- [ ] Understand basic concurrency patterns (fan-in, fan-out, worker pool)
- [ ] Use `context.Context` for cancellation and deadlines

**Estimated Time**: 4-5 hours  
**Difficulty**: ⭐⭐⭐ Advanced

---

## 📚 Key Concepts

### 1. Goroutines

```go
// Launch a goroutine — lightweight thread
go func() {
    fmt.Println("Running concurrently")
}()

// Launch with a named function
func printNumbers() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
go printNumbers()

// Wait briefly (use WaitGroup in real code)
time.Sleep(time.Second)
```

### 2. Channels (Unbuffered vs Buffered)

```go
// Unbuffered channel — blocks until receiver is ready
ch := make(chan int)
go func() {
    ch <- 42  // blocks until main receives
}()
value := <-ch
fmt.Println(value)

// Buffered channel — non-blocking until full
bufCh := make(chan string, 3)
bufCh <- "a"  // doesn't block
bufCh <- "b"  // doesn't block
bufCh <- "c"  // doesn't block
// bufCh <- "d"  // would block (buffer full)

// Closing channels
close(bufCh)
for msg := range bufCh {
    fmt.Println(msg)
}
```

### 3. Select Statement

```go
ch1 := make(chan string)
ch2 := make(chan string)

select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout — no message received")
default:
    fmt.Println("No channel ready")
}
```

### 4. Sync — WaitGroup & Mutex

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println("Worker:", id)
    }(i)
}
wg.Wait()  // Wait for all goroutines

// Mutex — protect shared state
var (
    counter int
    mu      sync.Mutex
)

for i := 0; i < 1000; i++ {
    go func() {
        mu.Lock()
        counter++
        mu.Unlock()
    }()
}
```

### 5. Context for Cancellation

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    select {
    case <-ctx.Done():
        return // cancelled
    case <-time.After(2 * time.Second):
        fmt.Println("Work done")
    }
}()

// Cancel the goroutine
cancel()
```

---

## 🔗 Essential Links

### Official Resources
- [Tour of Go: Concurrency](https://tour.golang.org/concurrency/1)
- [Effective Go: Concurrency](https://golang.org/doc/effective_go#concurrency)
- [Go Blog: Share Memory By Communicating](https://golang.org/blog/codelab-share)

### Tutorials
- [Go by Example: Goroutines](https://gobyexample.com/goroutines)
- [Go by Example: Channels](https://gobyexample.com/channels)
- [Go by Example: Select](https://gobyexample.com/select)
- [Go by Example: WaitGroups](https://gobyexample.com/waitgroups)
- [Go by Example: Mutexes](https://gobyexample.com/mutexes)

### Videos
- [Go Concurrency Patterns (Rob Pike)](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Google I/O 2012: Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Concurrency is not Parallelism (Rob Pike)](https://www.youtube.com/watch?v=oV9rvDllKEg)

---

## 💻 Examples

### Example 1: Goroutines and Basic Channels
See `examples/01-goroutines.go`

### Example 2: WaitGroup and Mutex
See `examples/02-sync.go`

### Example 3: Select and Channels
See `examples/03-select.go`

---

## ✋ Hands-On Exercises

### Exercise 1: Concurrent Counter with Mutex
**Objective**: Protect shared state with synchronization

**Instructions**:
1. Create `exercises/ex1_concurrent_counter.go`
2. Launch 10 goroutines that each increment a shared counter 1000 times
3. Use `sync.Mutex` to protect the counter
4. Use `sync.WaitGroup` to wait for all goroutines
5. Print the final counter value (should be 10000)

**Expected Output**:
```
Final counter: 10000
```

**Solution**: See `solutions/ex1_concurrent_counter.go`

---

### Exercise 2: Worker Pool with Channels
**Objective**: Implement the worker pool pattern

**Instructions**:
1. Create `exercises/ex2_worker_pool.go`
2. Create a buffered job channel (capacity 10) and a results channel (capacity 10)
3. Launch 3 worker goroutines that read jobs and send results
4. Send 10 jobs (each job is an integer to square)
5. Print all results

**Expected Output**:
```
Worker 1 processing job 1 → result: 1
Worker 2 processing job 2 → result: 4
Worker 3 processing job 3 → result: 9
...
All jobs completed!
```

**Solution**: See `solutions/ex2_worker_pool.go`

---

### Exercise 3: Ping-Pong with Select
**Objective**: Use select with channels for coordination

**Instructions**:
1. Create `exercises/ex3_pingpong.go`
2. Create a ping channel and a pong channel
3. Launch a goroutine that receives on ping, prints "ping", sends on pong
4. In main, use a `select` to alternate between sending and receiving
5. Play exactly 5 rounds, then exit gracefully

**Expected Output**:
```
Round 1: ping
Round 1: pong
Round 2: ping
Round 2: pong
...
Game over after 5 rounds
```

**Solution**: See `solutions/ex3_pingpong.go`

---

## 🧪 Testing Your Knowledge

### Quiz
1. **What's the difference between a goroutine and an OS thread?**
   - Goroutines are lightweight, multiplexed onto OS threads by the Go scheduler. They start with tiny stacks (~2KB) and grow as needed.

2. **What happens when you send to an unbuffered channel with no receiver?**
   - The sender blocks until a receiver is ready

3. **How do you prevent a goroutine from leaking?**
   - Use a cancellation channel or `context.Context` to signal shutdown; always ensure goroutines can exit

4. **What does `sync.WaitGroup` do?**
   - It waits for a collection of goroutines to finish: `Add(int)` increments, `Done()` decrements, `Wait()` blocks until counter is 0

5. **What's the difference between `sync.Mutex` and `sync.RWMutex`?**
   - `Mutex` locks exclusively; `RWMutex` allows multiple readers or one writer. Use `RWMutex` when reads dominate

---

## 🚀 Next Steps

1. ✅ Complete all exercises in this module
2. ✅ Run all examples and understand the concurrency model
3. ✅ Practice with different concurrency patterns
4. ✅ Move to Module 08: Testing

---

## 📝 Notes

Write your learning notes here:

```
- "Do not communicate by sharing memory; instead, share memory by communicating"
- Channels are typed conduits for goroutine communication
- select lets a goroutine wait on multiple channel operations
- Always close channels from the sender side
- Use sync package for traditional locking when appropriate
```

---

## 🔗 Additional Resources
- [Go Concurrency Patterns (Rob Pike talk)](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Advanced Go Concurrency Patterns](https://blog.golang.org/advanced-go-concurrency-patterns)
- [The Go Memory Model](https://golang.org/ref/mem)
- [Context Package Documentation](https://golang.org/pkg/context/)

---

**Status**: Ready to Learn  
**Completion**: Mark as complete once you finish all exercises
