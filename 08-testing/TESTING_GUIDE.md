# Advanced Testing Patterns in Go

This guide builds on the basics covered in `README.md`. Each section provides production-quality patterns with complete, compilable code examples.

---

## Table of Contents

1. [Test Structure & Organization](#1-test-structure--organization)
2. [Interfaces for Testability](#2-interfaces-for-testability)
3. [Mocking with testify/mock](#3-mocking-with-testifymock)
4. [httptest for HTTP Testing](#4-httptest-for-http-testing)
5. [Cleanup with t.Cleanup](#5-cleanup-with-tcleanup)
6. [Parallel Tests & Race Detection](#6-parallel-tests--race-detection)
7. [Fuzzing (Go 1.18+)](#7-fuzzing-go-118)
8. [Test Coverage Deep Dive](#8-test-coverage-deep-dive)
9. [Containerized Integration Tests](#9-containerized-integration-tests)
10. [Test Fixtures & Helpers](#10-test-fixtures--helpers)
11. [Testing JSON APIs](#11-testing-json-apis)
12. [Continuous Testing](#12-continuous-testing)

---

## 1. Test Structure & Organization

### Overview

Well-organized tests are as important as well-organized production code. Go conventions include build tags to separate unit from integration tests, a `testdata` directory for fixture files, and golden files for capturing complex expected output.

### Build Tags for Integration Tests

Use `//go:build` constraints to split fast unit tests from slower integration tests that depend on databases, network calls, or external services.

```go
// File: math_test.go (runs always)
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("got %d, want %d", result, 5)
    }
}
```

```go
// File: integration_test.go (runs only with `go test -tags=integration`)
//go:build integration

package math

import "testing"

func TestAddWithDatabase(t *testing.T) {
    // This test only compiles and runs when the integration tag is set.
    t.Log("integration test: would connect to a real DB here")
}
```

Run with:
```bash
go test -tags=integration -v ./...
```

```go
// File: database_test.go (runs always, but skips without a flag)
package db

import "testing"

func TestPostgresConnection(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test in short mode")
    }
    // actual connection logic
    t.Log("connecting to postgres...")
}
```

### testdata Directory

The Go toolchain treats `testdata` directories specially -- they are ignored during package compilation but can be read by tests at runtime.

```
mypackage/
  mypackage.go
  mypackage_test.go
  testdata/
    fixtures/
      user.json
      products.csv
    golden/
      user_response.golden
```

```go
// File: mypackage_test.go
package mypackage

import (
    "os"
    "path/filepath"
    "testing"
)

func TestWithFixture(t *testing.T) {
    path := filepath.Join("testdata", "fixtures", "user.json")
    data, err := os.ReadFile(path)
    if err != nil {
        t.Fatal(err)
    }
    t.Logf("loaded fixture: %s", string(data))
}
```

### Golden Files Pattern

Golden files store expected output so you can compare test results without hardcoding large strings.

```go
// File: golden_test.go
package mypackage

import (
    "flag"
    "os"
    "path/filepath"
    "testing"
)

var update = flag.Bool("update", false, "update golden files")

func TestGoldenFile(t *testing.T) {
    got := renderHTML("hello") // imagine this returns a string

    golden := filepath.Join("testdata", "golden", "render_html.golden")
    if *update {
        os.WriteFile(golden, []byte(got), 0644)
    }

    want, err := os.ReadFile(golden)
    if err != nil {
        t.Fatalf("failed to read golden file: %v", err)
    }

    if got != string(want) {
        t.Errorf("got:\n%s\n\nwant:\n%s", got, string(want))
    }
}
```

Update golden files:
```bash
go test -update ./...
```

### Best Practices

- Place only files needed by tests in `testdata`.
- Use `testing.Short()` to let users skip long-running tests with `go test -short`.
- Commit golden files to version control for reproducibility.
- Name golden files with `.golden` extension so they are clearly distinguished.

---

## 2. Interfaces for Testability

### Overview

Code written against interfaces rather than concrete types is inherently testable. By defining narrow interfaces at the point of use, you can swap real implementations for test doubles without importing mock frameworks.

### Real Example: UserRepository

```go
// File: user.go
package user

import (
    "errors"
    "fmt"
)

type User struct {
    ID    int
    Email string
    Name  string
}

var ErrNotFound = errors.New("user not found")

// UserRepository defines the contract for data access.
// Defining it here lets consumers depend on the interface, not the implementation.
type UserRepository interface {
    FindByID(id int) (*User, error)
    Save(user *User) error
}

// UserService contains business logic and depends on the repository abstraction.
type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetDisplayName(id int) (string, error) {
    u, err := s.repo.FindByID(id)
    if err != nil {
        return "", fmt.Errorf("get display name: %w", err)
    }
    if u.Name != "" {
        return u.Name, nil
    }
    return u.Email, nil
}
```

```go
// File: user_test.go
package user

import (
    "errors"
    "testing"
)

// mockRepo is a hand-rolled mock implementing UserRepository.
type mockRepo struct {
    users map[int]*User
}

func (m *mockRepo) FindByID(id int) (*User, error) {
    u, ok := m.users[id]
    if !ok {
        return nil, ErrNotFound
    }
    return u, nil
}

func (m *mockRepo) Save(user *User) error {
    m.users[user.ID] = user
    return nil
}

func TestGetDisplayName_ReturnsName(t *testing.T) {
    repo := &mockRepo{users: map[int]*User{
        1: {ID: 1, Email: "alice@example.com", Name: "Alice"},
    }}
    svc := NewUserService(repo)

    name, err := svc.GetDisplayName(1)
    if err != nil {
        t.Fatal(err)
    }
    if name != "Alice" {
        t.Errorf("got %q, want %q", name, "Alice")
    }
}

func TestGetDisplayName_FallsBackToEmail(t *testing.T) {
    repo := &mockRepo{users: map[int]*User{
        2: {ID: 2, Email: "bob@example.com"},
    }}
    svc := NewUserService(repo)

    name, err := svc.GetDisplayName(2)
    if err != nil {
        t.Fatal(err)
    }
    if name != "bob@example.com" {
        t.Errorf("got %q, want %q", name, "bob@example.com")
    }
}

func TestGetDisplayName_NotFound(t *testing.T) {
    repo := &mockRepo{users: map[int]*User{}}
    svc := NewUserService(repo)

    _, err := svc.GetDisplayName(99)
    if !errors.Is(err, ErrNotFound) {
        t.Errorf("expected ErrNotFound, got %v", err)
    }
}
```

```bash
# Output
=== RUN   TestGetDisplayName_ReturnsName
--- PASS: TestGetDisplayName_ReturnsName (0.00s)
=== RUN   TestGetDisplayName_FallsBackToEmail
--- PASS: TestGetDisplayName_FallsBackToEmail (0.00s)
=== RUN   TestGetDisplayName_NotFound
--- PASS: TestGetDisplayName_NotFound (0.00s)
PASS
ok      user    0.102s
```

### Pitfalls

- Do not define interfaces in the implementation package for the benefit of tests. Define interfaces where they are consumed.
- Keep interfaces small (1-3 methods). The `UserRepository` interface is a good size.
- Hand-rolled mocks are fine for simple cases. For complex behavior, use `testify/mock`.

---

## 3. Mocking with testify/mock

### Overview

The `testify` library provides `mock.Mock` -- a struct that records expectations and assertions on method calls. You embed `mock.Mock` in your test struct and call `On()`, `Return()`, and `AssertExpectations()`.

### Installation

```bash
go get github.com/stretchr/testify
```

### Example: Mocking an Email Sender

```go
// File: notifier.go
package notifier

import "fmt"

type User struct {
    ID    int
    Email string
}

type EmailSender interface {
    Send(to, subject, body string) error
}

type NotificationService struct {
    sender EmailSender
}

func NewNotificationService(sender EmailSender) *NotificationService {
    return &NotificationService{sender: sender}
}

func (s *NotificationService) Welcome(user *User) error {
    body := fmt.Sprintf("Welcome %s!", user.Email)
    return s.sender.Send(user.Email, "Welcome!", body)
}
```

```go
// File: notifier_test.go
package notifier

import (
    "testing"

    "github.com/stretchr/testify/mock"
)

// MockEmailSender embeds mock.Mock to record and verify expectations.
type MockEmailSender struct {
    mock.Mock
}

func (m *MockEmailSender) Send(to, subject, body string) error {
    args := m.Called(to, subject, body)
    return args.Error(0)
}

func TestWelcome_SendsEmail(t *testing.T) {
    mockSender := new(MockEmailSender)
    svc := NewNotificationService(mockSender)

    user := &User{ID: 1, Email: "alice@example.com"}

    // Set expectation: Send should be called once with these args, return nil.
    mockSender.On("Send", "alice@example.com", "Welcome!", "Welcome alice@example.com!").Return(nil)

    err := svc.Welcome(user)
    if err != nil {
        t.Fatal(err)
    }

    // Assert that the expectation was met exactly.
    mockSender.AssertExpectations(t)
}

func TestWelcome_SendFails(t *testing.T) {
    mockSender := new(MockEmailSender)
    svc := NewNotificationService(mockSender)

    user := &User{ID: 2, Email: "bob@example.com"}
    mockSender.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("smtp error"))

    err := svc.Welcome(user)
    if err == nil {
        t.Fatal("expected error but got nil")
    }
}

func TestWelcome_AnyMatch(t *testing.T) {
    mockSender := new(MockEmailSender)
    svc := NewNotificationService(mockSender)

    user := &User{ID: 3, Email: "carol@example.com"}

    // Use mock.AnythingOfType or mock.Anything for flexible matching.
    mockSender.On("Send", mock.AnythingOfType("string"), "Welcome!", mock.Anything).Return(nil)

    err := svc.Welcome(user)
    if err != nil {
        t.Fatal(err)
    }
    mockSender.AssertExpectations(t)
}
```

```bash
# Output
=== RUN   TestWelcome_SendsEmail
--- PASS: TestWelcome_SendsEmail (0.00s)
=== RUN   TestWelcome_SendFails
--- PASS: TestWelcome_SendFails (0.00s)
=== RUN   TestWelcome_AnyMatch
--- PASS: TestWelcome_AnyMatch (0.00s)
PASS
ok      notifier        0.108s
```

### Matchers

| Matcher | Description |
|---|---|
| `mock.Anything` | Matches any value |
| `mock.AnythingOfType("string")` | Matches any value of a specific type |
| `mock.MatchedBy(func(v) bool)` | Matches with a custom predicate |

### Pitfalls

- Always call `AssertExpectations(t)` in the test (or use `defer`) to verify all expectations were met.
- Be careful with pointer arguments: `m.Called(obj)` on a struct pointer records the pointer address. Match with `mock.MatchedBy` if you need value equality.
- Use `Once()` to indicate a call should happen exactly once: `mockSender.On("Send", ...).Return(nil).Once()`.
- Never reuse a `Mock` instance across multiple parallel tests.

---

## 4. httptest for HTTP Testing

### Overview

The `net/http/httptest` package provides two main primitives: `httptest.NewRecorder` for testing handlers in-process and `httptest.NewServer` for spinning up a real HTTP server backed by your handler.

### Testing Handlers with ResponseRecorder

```go
// File: handler.go
package handler

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    u := User{ID: 1, Name: "Alice"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(u)
}
```

```go
// File: handler_test.go
package handler

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestUserHandler_Get_ReturnsUser(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
    rec := httptest.NewRecorder()

    UserHandler(rec, req)

    resp := rec.Result()
    if resp.StatusCode != http.StatusOK {
        t.Errorf("got status %d, want %d", resp.StatusCode, http.StatusOK)
    }

    var u User
    json.NewDecoder(resp.Body).Decode(&u)
    resp.Body.Close()

    if u.ID != 1 || u.Name != "Alice" {
        t.Errorf("got %+v, want ID=1 Name=Alice", u)
    }
}

func TestUserHandler_WrongMethod(t *testing.T) {
    req := httptest.NewRequest(http.MethodPost, "/users/1", nil)
    rec := httptest.NewRecorder()

    UserHandler(rec, req)

    resp := rec.Result()
    if resp.StatusCode != http.StatusMethodNotAllowed {
        t.Errorf("got status %d, want %d", resp.StatusCode, http.StatusMethodNotAllowed)
    }
    resp.Body.Close()
}
```

### Testing a Full Server with httptest.NewServer

```go
// File: server_test.go
package handler

import (
    "encoding/json"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
)

func setupTestServer() *httptest.Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/users", UserHandler)
    return httptest.NewServer(mux)
}

func TestUserHandlerViaServer(t *testing.T) {
    srv := setupTestServer()
    defer srv.Close()

    resp, err := http.Get(srv.URL + "/users")
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatal(err)
    }

    var u User
    if err := json.Unmarshal(body, &u); err != nil {
        t.Fatal(err)
    }

    if u.Name != "Alice" {
        t.Errorf("got %q, want %q", u.Name, "Alice")
    }
}
```

```bash
# Output
=== RUN   TestUserHandler_Get_ReturnsUser
--- PASS: TestUserHandler_Get_ReturnsUser (0.00s)
=== RUN   TestUserHandler_WrongMethod
--- PASS: TestUserHandler_WrongMethod (0.00s)
=== RUN   TestUserHandlerViaServer
--- PASS: TestUserHandlerViaServer (0.00s)
PASS
ok      handler 0.124s
```

### Testing Middleware with httptest

```go
// File: middleware_test.go
package handler

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestLoggingMiddleware(t *testing.T) {
    var logged string
    logger := func(msg string) { logged = msg }

    middleware := func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            logged = r.Method + " " + r.URL.Path
            next.ServeHTTP(w, r)
        })
    }

    handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))

    req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
    rec := httptest.NewRecorder()
    handler.ServeHTTP(rec, req)

    if logged != "GET /api/users" {
        t.Errorf("got %q, want %q", logged, "GET /api/users")
    }
    if rec.Code != http.StatusOK {
        t.Errorf("got %d, want %d", rec.Code, http.StatusOK)
    }
}
```

### Best Practices

- Always close the response body (`resp.Body.Close()`) when using `rec.Result()`.
- Use `httptest.NewServer` for end-to-end handler chains with middleware.
- Pass a `*testing.T` to `httptest.NewServer` to have it log and clean up automatically: `httptest.NewServer(handler)` -- but in older Go versions, manage cleanup yourself with `defer srv.Close()`.
- Set request headers explicitly: `req.Header.Set("Authorization", "Bearer token")`.

---

## 5. Cleanup with t.Cleanup

### Overview

`t.Cleanup` (introduced in Go 1.14) registers a function to run when the test and all its subtests complete. It replaces the defer-in-test idiom and works correctly with `t.Parallel()`, ensuring cleanup runs after all parallel subtests finish, not just the current one.

### Example: Temp Files and Server Cleanup

```go
// File: cleanup_test.go
package cleanup

import (
    "net/http"
    "net/http/httptest"
    "os"
    "path/filepath"
    "testing"
)

func TestWithTempFile(t *testing.T) {
    dir := t.TempDir()
    // Note: t.TempDir() also auto-registers cleanup for removal.
    path := filepath.Join(dir, "data.txt")
    if err := os.WriteFile(path, []byte("hello"), 0644); err != nil {
        t.Fatal(err)
    }

    data, err := os.ReadFile(path)
    if err != nil {
        t.Fatal(err)
    }
    if string(data) != "hello" {
        t.Errorf("got %q, want %q", string(data), "hello")
    }
    // No manual cleanup needed; t.TempDir() removes the directory at the end.
}

type Server struct {
    URL string
    // internal fields
}

func setupServer(t *testing.T) *Server {
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    }))

    t.Cleanup(func() {
        srv.Close()
    })

    return &Server{URL: srv.URL}
}

func TestWithServerCleanup(t *testing.T) {
    srv := setupServer(t)
    // Subtest that uses the server.
    t.Run("sub", func(t *testing.T) {
        resp, err := http.Get(srv.URL + "/path")
        if err != nil {
            t.Fatal(err)
        }
        resp.Body.Close()
        if resp.StatusCode != http.StatusOK {
            t.Errorf("got %d, want %d", resp.StatusCode, http.StatusOK)
        }
    })
    // srv is still available here; cleanups run after this function returns.
}

type Resource struct {
    ID int
}

func acquireResource(t *testing.T) *Resource {
    r := &Resource{ID: 42}
    t.Cleanup(func() {
        t.Logf("cleaning up resource %d", r.ID)
    })
    return r
}

func TestResourceCleanup(t *testing.T) {
    r := acquireResource(t)
    t.Logf("using resource %d", r.ID)
}
```

```bash
# Output
=== RUN   TestWithTempFile
--- PASS: TestWithTempFile (0.00s)
=== RUN   TestWithServerCleanup
=== RUN   TestWithServerCleanup/sub
--- PASS: TestWithServerCleanup (0.00s)
=== RUN   TestResourceCleanup
--- PASS: TestResourceCleanup (0.00s)
ok      cleanup 0.115s
```

### t.TempDir

`t.TempDir()` is a convenience method that creates a temp directory and registers `t.Cleanup` to remove it. Use it instead of `ioutil.TempDir` + manual defer removal.

### Pitfalls

- Cleanup functions run in LIFO (last-in-first-out) order.
- If a test calls `t.Fatal`, cleanup functions still execute.
- Do NOT use `t.Cleanup` inside a helper function that is called from a `t.Run` subtest if the helper does not receive the subtest's `*testing.T` -- it will register against the parent test instead.

---

## 6. Parallel Tests & Race Detection

### Overview

Go's `t.Parallel()` tells the test runner that a test can run concurrently with other parallel tests. Combined with the `-race` flag, this detects data races between goroutines. Parallel tests speed up test suites but must be written carefully to avoid shared state.

### Parallel Tests

```go
// File: parallel_test.go
package parallel

import (
    "testing"
    "time"
)

func slowOperation(id int) string {
    time.Sleep(50 * time.Millisecond)
    return "result"
}

func TestSequential(t *testing.T) {
    // Without t.Parallel, these run one at a time: ~150ms total
    t.Run("op1", func(t *testing.T) {
        r := slowOperation(1)
        if r != "result" {
            t.Error("unexpected")
        }
    })
    t.Run("op2", func(t *testing.T) {
        r := slowOperation(2)
        if r != "result" {
            t.Error("unexpected")
        }
    })
    t.Run("op3", func(t *testing.T) {
        r := slowOperation(3)
        if r != "result" {
            t.Error("unexpected")
        }
    })
}

func TestParallel(t *testing.T) {
    // With t.Parallel(), all subtests run concurrently: ~50ms total
    t.Run("op1", func(t *testing.T) {
        t.Parallel()
        r := slowOperation(1)
        if r != "result" {
            t.Error("unexpected")
        }
    })
    t.Run("op2", func(t *testing.T) {
        t.Parallel()
        r := slowOperation(2)
        if r != "result" {
            t.Error("unexpected")
        }
    })
    t.Run("op3", func(t *testing.T) {
        t.Parallel()
        r := slowOperation(3)
        if r != "result" {
            t.Error("unexpected")
        }
    })
}
```

```bash
# Run with timing
go test -v -run TestParallel -timeout 30s

# Output
=== RUN   TestParallel
=== RUN   TestParallel/op1
=== RUN   TestParallel/op2
=== RUN   TestParallel/op3
--- PASS: TestParallel (0.05s)
    --- PASS: TestParallel/op1 (0.05s)
    --- PASS: TestParallel/op2 (0.05s)
    --- PASS: TestParallel/op3 (0.05s)
PASS
ok      parallel        0.060s
```

### Race Detection

```go
// File: race_test.go
package parallel

import (
    "sync"
    "testing"
)

type Counter struct {
    mu    sync.Mutex
    value int
}

func TestRaceDetection(t *testing.T) {
    var counter Counter
    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.mu.Lock()
            counter.value++
            counter.mu.Unlock()
        }()
    }
    wg.Wait()

    if counter.value != 100 {
        t.Errorf("got %d, want %d", counter.value, 100)
    }
}

// BAD: This test has a data race.
func TestDataRace(t *testing.T) {
    counter := 0
    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter++ // unsynchronized write -- race condition!
        }()
    }
    wg.Wait()

    t.Logf("counter = %d (may not be 100 due to race)", counter)
}
```

```bash
go test -race -v -run TestDataRace

# Output (abbreviated)
==================
WARNING: DATA RACE
Write at 0x00c0000... by goroutine 8:
  parallel.TestDataRace.func1()
      /path/to/race_test.go:44 +0x3c
==================
--- FAIL: TestDataRace (0.00s)
    testing.go:...: race detected during execution of test
FAIL
```

### Controlling Parallelism

```bash
# Set max parallel tests
go test -parallel 4 ./...

# Default is GOMAXPROCS (usually number of CPUs)
```

### Best Practices

- Always call `t.Parallel()` at the very start of the test function body, before any logic.
- Each parallel subtest gets a copy of the loop variable (Go 1.22+). For older Go, capture explicitly: `tt := tt`.
- Use the `-race` flag in CI and during development -- it catches bugs that might otherwise appear only in production under load.
- Avoid sharing mutable state between parallel tests. If you must share, use a `sync.Mutex` or channel.
- `t.Cleanup` functions wait for all parallel subtests to finish before executing.

---

## 7. Fuzzing (Go 1.18+)

### Overview

Fuzzing automatically generates random inputs to find edge cases that crash or produce unexpected results. A fuzz test starts with a corpus of seed inputs and mutates them to explore new code paths. Fuzzing is especially effective at finding off-by-one errors, integer overflows, and panic-inducing inputs.

### Basic Fuzz Test

```go
// File: reverse.go
package fuzzex

import "fmt"

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

```go
// File: reverse_fuzz_test.go
package fuzzex

import (
    "testing"
    "unicode/utf8"
)

// FuzzReverse is a fuzz target.
// The function signature is fixed: FuzzXxx(f *testing.F, ...)
func FuzzReverse(f *testing.F) {
    // Seed corpus: provide known inputs.
    f.Add("hello")
    f.Add("world")
    f.Add("12345")
    f.Add("")           // empty string
    f.Add("!@#$%")      // special characters
    f.Add("a")          // single character

    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)

        // Reversing twice must yield the original string.
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }

        // A valid UTF-8 string must remain valid UTF-8 after reversal.
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8: %q", rev)
        }
    })
}
```

```bash
# Run fuzz test (no -fuzz flag = corpus only)
go test -fuzz FuzzReverse -fuzztime 10s ./...

# Output while fuzzing
=== FUZZ  FuzzReverse
fuzz: elapsed 0s, gathering baseline coverage: 0/6
fuzz: elapsed 0s, gathering baseline coverage: 6/6 (now fuzzing)
fuzz: elapsed 3s, execs: 427486 (142494/sec), new interesting: 1 (total 7)
fuzz: elapsed 6s, execs: 842195 (138241/sec), new interesting: 2 (total 8)
fuzz: elapsed 9s, execs: 1294652 (150836/sec), new interesting: 3 (total 9)
--- PASS
ok      fuzzex  10s
```

### Finding a Bug with Fuzzing

```go
// File: divide.go
package fuzzex

func Divide(a, b int) int {
    return a / b // panics when b == 0
}
```

```go
// File: divide_fuzz_test.go
package fuzzex

import "testing"

func FuzzDivide(f *testing.F) {
    f.Add(10, 2)
    f.Add(0, 5)
    f.Add(-10, 3)

    f.Fuzz(func(t *testing.T, a, b int) {
        if b == 0 {
            t.Skip("division by zero")
        }
        result := Divide(a, b)
        // Property: (a / b) * b + (a % b) == a
        if b != 0 && result*b+(a% b) != a {
            t.Errorf("property violation: Divide(%d, %d) = %d", a, b, result)
        }
    })
}
```

```bash
# If a crash is found, the input is written to testdata/fuzz/FuzzDivide/
$ cat testdata/fuzz/FuzzDivide/<hash>
go test fuzz v1
int(10)
int(0)
```

### Running Specific Corpus Entries

```bash
# Run only the seed corpus (no fuzzing)
go test -run FuzzDivide/seed ./...

# Run a specific crash input
go test -run 'FuzzDivide/<hash>' ./...
```

### Best Practices

- Properties (invariants) work best: reversing twice gives original, encoding then decoding is idempotent, etc.
- The fuzz function signature must be `func(*testing.T, <string|[]byte|int|int8|int16|int32|int64|uint|...>)`.
- Supported types: `string`, `[]byte`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`, `bool`.
- Always `t.Skip()` invalid inputs instead of returning early. This teaches the fuzzer to avoid invalid combinations.
- Run fuzzing in CI for a fixed duration with `-fuzztime 5m`.
- Commit crashing inputs found by fuzzing to `testdata/fuzz/` so they become regression tests.

---

## 8. Test Coverage Deep Dive

### Overview

Go's built-in coverage tool measures which statements your tests execute. Beyond `go test -cover`, advanced features include coverage profiles, coverage modes (count vs atomic), excluding generated code, and combining results from multiple test suites.

### Basic Coverage

```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out      # per-function breakdown
go tool cover -html=coverage.out      # visual HTML report
```

```bash
# Output of go tool cover -func
fuzzex/divide.go:3:       Divide       100.0%
fuzzex/reverse.go:4:      Reverse      100.0%
total:                                 100.0%
```

### Coverage Modes

```bash
# set (default): records whether a statement was reached
go test -covermode=set -coverprofile=coverage.out ./...

# count: records how many times a statement was reached
go test -covermode=count -coverprofile=coverage.out ./...

# atomic: like count but safe for concurrent access (needed with -race)
go test -race -covermode=atomic -coverprofile=coverage.out ./...
```

Use `count` mode when you want to verify that certain paths execute more than zero times (e.g. retry loops). Use `atomic` when testing concurrent code.

### Coverage for Integration Tests

Run separate coverage profiles for unit and integration tests, then merge them:

```bash
# Unit test coverage
go test -covermode=count -coverprofile=unit.cov ./...

# Integration test coverage
go test -tags=integration -covermode=count -coverprofile=integration.cov ./...

# Combine profiles
echo "mode: count" > combined.cov
tail -n +2 unit.cov >> combined.cov
tail -n +2 integration.cov >> combined.cov

# View combined result
go tool cover -html=combined.cov
```

### Excluding Generated Code

Use a build tag or file naming convention, then filter out generated files from the coverage report:

```go
// File: generated.go
// Code generated by protoc-gen-go. DO NOT EDIT.

package api
```

The convention `// Code generated ... DO NOT EDIT.` is recognized by Go tools. Files matching this pattern are excluded by `go tool cover` in Go 1.20+.

Alternatively, use `grep` to filter:

```bash
go test -coverprofile=raw.cov ./...
# Remove lines from generated files
grep -v "_gen.go:" raw.cov > coverage.out
```

### Enforcing Coverage Thresholds

```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out | tail -1 | awk '{print $NF}' | tr -d '%'
```

```go
// File: coverage_check_test.go (in your CI verification)
package coverage

import (
    "os"
    "testing"
)

func TestCoverageThreshold(t *testing.T) {
    // This is just a placeholder; coverage enforcement is better done
    // in CI scripts. Example shell snippet:
    //
    //   COVER=$(go tool cover -func=coverage.out | grep total | awk '{print $NF}' | tr -d '%')
    //   if [ "$COVER" -lt 80 ]; then
    //     echo "Coverage ${COVER}% is below threshold 80%"
    //     exit 1
    //   fi
    t.Skip("use shell script in CI instead")
}
```

### Per-Package Coverage

```bash
go test -coverprofile=coverage.out ./...
# Per-package breakdown
go test -cover ./...
```

```bash
# Output
ok      user            0.102s  coverage: 85.7% of statements
ok      notifier        0.108s  coverage: 100.0% of statements
ok      handler         0.124s  coverage: 92.3% of statements
ok      fuzzex          0.130s  coverage: 100.0% of statements
```

### Pitfalls

- Coverage is per-package, not global. Use `--coverpkg=all` to measure cross-package coverage.
- `-coverpkg=./...` lets you measure how tests in one package cover code in another.
- 100% coverage does not mean bug-free code. Focus coverage efforts on business logic.
- Exclude generated code early to avoid artificially inflated or deflated numbers.

---

## 9. Containerized Integration Tests

### Overview

`testcontainers-go` provides disposable Docker containers for integration testing. Spin up PostgreSQL, MongoDB, Kafka, or any Docker image as part of your test, then tear it down automatically. This replaces complex test harnesses and mocks with real dependencies.

### Installation

```bash
go get github.com/testcontainers/testcontainers-go
go get github.com/testcontainers/testcontainers-go/modules/postgres
go get github.com/testcontainers/testcontainers-go/modules/mongodb
```

### PostgreSQL Example

```go
// File: postgres_test.go
//go:build integration

package db

import (
    "context"
    "database/sql"
    "testing"
    "time"

    _ "github.com/lib/pq"
    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/modules/postgres"
    "github.com/testcontainers/testcontainers-go/wait"
)

func setupPostgres(t *testing.T) (*sql.DB, func()) {
    ctx := context.Background()

    pg, err := postgres.RunContainer(ctx,
        testcontainers.WithImage("postgres:16-alpine"),
        postgres.WithDatabase("testdb"),
        postgres.WithUsername("test"),
        postgres.WithPassword("test"),
        testcontainers.WithWaitStrategy(
            wait.ForLog("database system is ready to accept connections").
                WithOccurrence(2).WithStartupTimeout(30*time.Second),
        ),
    )
    if err != nil {
        t.Fatal(err)
    }

    connStr, err := pg.ConnectionString(ctx, "sslmode=disable")
    if err != nil {
        t.Fatal(err)
    }

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        t.Fatal(err)
    }

    if err := db.Ping(); err != nil {
        t.Fatal(err)
    }

    cleanup := func() {
        db.Close()
        pg.Terminate(ctx)
    }

    return db, cleanup
}

func TestUserInsertAndQuery(t *testing.T) {
    db, cleanup := setupPostgres(t)
    defer cleanup()

    _, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email TEXT NOT NULL UNIQUE,
        name TEXT NOT NULL
    )`)
    if err != nil {
        t.Fatal(err)
    }

    _, err = db.Exec(`INSERT INTO users (email, name) VALUES ($1, $2)`,
        "alice@example.com", "Alice")
    if err != nil {
        t.Fatal(err)
    }

    var name string
    err = db.QueryRow(`SELECT name FROM users WHERE email = $1`,
        "alice@example.com").Scan(&name)
    if err != nil {
        t.Fatal(err)
    }

    if name != "Alice" {
        t.Errorf("got %q, want %q", name, "Alice")
    }
}
```

```bash
# Run (requires Docker)
go test -tags=integration -v ./...

# Output
=== RUN   TestUserInsertAndQuery
    postgres_test.go:53: creating postgres container...
    postgres_test.go:73: inserting user...
    postgres_test.go:82: querying user...
--- PASS: TestUserInsertAndQuery (4.23s)
PASS
ok      db      4.23s
```

### MongoDB Example

```go
// File: mongo_test.go
//go:build integration

package db

import (
    "context"
    "testing"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func TestMongoInsertAndQuery(t *testing.T) {
    ctx := context.Background()

    mongoC, err := mongodb.RunContainer(ctx,
        testcontainers.WithImage("mongo:7"),
    )
    if err != nil {
        t.Fatal(err)
    }
    defer mongoC.Terminate(ctx)

    connStr, err := mongoC.ConnectionString(ctx)
    if err != nil {
        t.Fatal(err)
    }

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
    if err != nil {
        t.Fatal(err)
    }
    defer client.Disconnect(ctx)

    coll := client.Database("testdb").Collection("users")
    _, err = coll.InsertOne(ctx, bson.M{"name": "Alice", "email": "alice@example.com"})
    if err != nil {
        t.Fatal(err)
    }

    var result bson.M
    err = coll.FindOne(ctx, bson.M{"name": "Alice"}).Decode(&result)
    if err != nil {
        t.Fatal(err)
    }

    if result["email"] != "alice@example.com" {
        t.Errorf("got %v, want alice@example.com", result["email"])
    }
}
```

### Custom Docker Compose

```go
// File: docker_compose_test.go
//go:build integration

package db

import (
    "context"
    "testing"

    "github.com/testcontainers/testcontainers-go"
)

func TestWithDockerCompose(t *testing.T) {
    ctx := context.Background()

    compose, err := testcontainers.NewDockerComposeWith(testcontainers.WithComposeFiles(
        "testdata/docker-compose.yml",
    ))
    if err != nil {
        t.Fatal(err)
    }

    t.Cleanup(func() {
        compose.Down(ctx, testcontainers.RemoveOrphans(true))
    })

    err = compose.Up(ctx, testcontainers.Wait(true))
    if err != nil {
        t.Fatal(err)
    }

    t.Log("all services are up and ready")
}
```

### Best Practices

- Always tag containerized tests with `//go:build integration` so `go test ./...` skips them unless explicitly requested.
- Use `t.Cleanup` for teardown to ensure containers terminate even when tests panic.
- Tune `WithStartupTimeout` for slower CI environments (default is usually 60s).
- Run these tests in CI with Docker. For GitHub Actions, the `docker` service is available natively on ubuntu-latest.
- Do not run containerized tests in parallel if they share the same port -- either use random ports (default) or set a `--test.parallel` limit.

---

## 10. Test Fixtures & Helpers

### Overview

Test helpers (`*testing.T`-based functions) reduce boilerplate by encapsulating common setup, teardown, and assertion logic. Fixtures are reusable data sets loaded from files or generated programmatically.

### Test Helpers with *testing.T

```go
// File: helper_test.go
package helpers

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "os"
    "path/filepath"
    "testing"
)

// Helper function: must accept *testing.T and call t.Helper().
func assertStatus(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("got status %d, want %d", got, want)
    }
}

func assertJSON(t *testing.T, got, want any) {
    t.Helper()
    gotBytes, err := json.Marshal(got)
    if err != nil {
        t.Fatal(err)
    }
    wantBytes, err := json.Marshal(want)
    if err != nil {
        t.Fatal(err)
    }
    if string(gotBytes) != string(wantBytes) {
        t.Errorf("got %s, want %s", string(gotBytes), string(wantBytes))
    }
}

func TestWithHelpers(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()

    http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"ok"}`))
    }).ServeHTTP(rec, req)

    assertStatus(t, rec.Code, http.StatusOK)

    var resp map[string]string
    json.NewDecoder(rec.Body).Decode(&resp)
    assertJSON(t, resp, map[string]string{"status": "ok"})
}

// Setup/teardown pattern for package-level state.
func setup(t *testing.T) func() {
    t.Helper()
    t.Log("setting up test resources")

    // Create temp directory
    dir := t.TempDir()

    return func() {
        t.Helper()
        t.Log("tearing down test resources")
        _ = dir // cleanup would happen here
    }
}

func TestWithSetupTeardown(t *testing.T) {
    defer setup(t)()
    t.Log("test body runs between setup and teardown")
}
```

### Loading Test Fixtures from Files

```go
// File: fixtures_test.go
package helpers

import (
    "encoding/json"
    "os"
    "path/filepath"
    "testing"
)

type TestUser struct {
    ID    int    `json:"id"`
    Email string `json:"email"`
    Name  string `json:"name"`
}

func loadFixture(t *testing.T, name string) []byte {
    t.Helper()
    path := filepath.Join("testdata", "fixtures", name)
    data, err := os.ReadFile(path)
    if err != nil {
        t.Fatalf("failed to load fixture %s: %v", name, err)
    }
    return data
}

func loadUsersFixture(t *testing.T) []TestUser {
    t.Helper()
    data := loadFixture(t, "users.json")
    var users []TestUser
    if err := json.Unmarshal(data, &users); err != nil {
        t.Fatalf("failed to unmarshal users: %v", err)
    }
    return users
}

func TestWithFixtures(t *testing.T) {
    users := loadUsersFixture(t)
    if len(users) != 3 {
        t.Fatalf("expected 3 users, got %d", len(users))
    }
    if users[0].Name != "Alice" {
        t.Errorf("got %q, want %q", users[0].Name, "Alice")
    }
}
```

### Fixture File

```json
// File: testdata/fixtures/users.json
[
    {"id": 1, "email": "alice@example.com", "name": "Alice"},
    {"id": 2, "email": "bob@example.com", "name": "Bob"},
    {"id": 3, "email": "carol@example.com", "name": "Carol"}
]
```

```bash
go test -v ./...

=== RUN   TestWithHelpers
--- PASS: TestWithHelpers (0.00s)
=== RUN   TestWithSetupTeardown
    helper_test.go:52: setting up test resources
    helper_test.go:59: test body runs between setup and teardown
    helper_test.go:56: tearing down test resources
--- PASS: TestWithSetupTeardown (0.00s)
=== RUN   TestWithFixtures
--- PASS: TestWithFixtures (0.00s)
PASS
ok      helpers 0.112s
```

### Best Practices

- Always call `t.Helper()` at the top of your helper functions so failure line numbers point to the caller, not the helper.
- Accept `*testing.T` (not `*testing.B`) in helpers. Use `testing.TB` if you need to support both benchmarks and tests.
- Use `t.Fatalf` for setup failures that make the test invalid.
- Return a cleanup function from setup functions for symmetry: `defer setup(t)()`.
- Keep fixtures small and focused. One fixture file per domain entity or use case.

---

## 11. Testing JSON APIs

### Overview

Testing JSON endpoints involves serializing request bodies, deserializing responses, and comparing structured data. Snapshot testing captures entire responses and flags unexpected changes, which is useful for detecting regressions in API output.

### Request/Response with JSON

```go
// File: api.go
package api

import (
    "encoding/json"
    "net/http"
    "strings"
)

type CreateUserRequest struct {
    Email string `json:"email"`
    Name  string `json:"name"`
}

type UserResponse struct {
    ID    int    `json:"id"`
    Email string `json:"email"`
    Name  string `json:"name"`
}

var users = map[int]UserResponse{}
var nextID = 1

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
        return
    }

    if strings.TrimSpace(req.Email) == "" {
        http.Error(w, `{"error":"email is required"}`, http.StatusBadRequest)
        return
    }

    user := UserResponse{
        ID:    nextID,
        Email: req.Email,
        Name:  req.Name,
    }
    nextID++
    users[user.ID] = user

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
```

```go
// File: api_test.go
package api

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func marshalJSON(t *testing.T, v any) []byte {
    t.Helper()
    data, err := json.Marshal(v)
    if err != nil {
        t.Fatal(err)
    }
    return data
}

func decodeJSON(t *testing.T, body *bytes.Buffer, v any) {
    t.Helper()
    if err := json.NewDecoder(body).Decode(v); err != nil {
        t.Fatal(err)
    }
}

func TestCreateUser_Success(t *testing.T) {
    reqBody := CreateUserRequest{Email: "alice@example.com", Name: "Alice"}
    req := httptest.NewRequest(http.MethodPost, "/users",
        bytes.NewReader(marshalJSON(t, reqBody)))
    rec := httptest.NewRecorder()

    CreateUserHandler(rec, req)

    if rec.Code != http.StatusCreated {
        t.Errorf("got %d, want %d", rec.Code, http.StatusCreated)
    }

    var resp UserResponse
    decodeJSON(t, rec.Body, &resp)

    if resp.ID == 0 {
        t.Error("expected non-zero ID")
    }
    if resp.Email != "alice@example.com" {
        t.Errorf("got %q, want %q", resp.Email, "alice@example.com")
    }
    if resp.Name != "Alice" {
        t.Errorf("got %q, want %q", resp.Name, "Alice")
    }
}

func TestCreateUser_EmptyEmail(t *testing.T) {
    reqBody := CreateUserRequest{Email: "", Name: "Bob"}
    req := httptest.NewRequest(http.MethodPost, "/users",
        bytes.NewReader(marshalJSON(t, reqBody)))
    rec := httptest.NewRecorder()

    CreateUserHandler(rec, req)

    if rec.Code != http.StatusBadRequest {
        t.Errorf("got %d, want %d", rec.Code, http.StatusBadRequest)
    }

    var errResp map[string]string
    decodeJSON(t, rec.Body, &errResp)
    if errResp["error"] != "email is required" {
        t.Errorf("got %q, want %q", errResp["error"], "email is required")
    }
}

func TestCreateUser_InvalidJSON(t *testing.T) {
    req := httptest.NewRequest(http.MethodPost, "/users",
        bytes.NewReader([]byte(`{broken}`)))
    rec := httptest.NewRecorder()

    CreateUserHandler(rec, req)

    if rec.Code != http.StatusBadRequest {
        t.Errorf("got %d, want %d", rec.Code, http.StatusBadRequest)
    }
}

func TestCreateUser_WrongMethod(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/users", nil)
    rec := httptest.NewRecorder()

    CreateUserHandler(rec, req)

    if rec.Code != http.StatusMethodNotAllowed {
        t.Errorf("got %d, want %d", rec.Code, http.StatusMethodNotAllowed)
    }
}
```

### Snapshot Testing

Snapshot testing captures the entire response body and compares against a stored snapshot on subsequent runs.

```go
// File: snapshot_test.go
package api

import (
    "bytes"
    "encoding/json"
    "flag"
    "net/http"
    "net/http/httptest"
    "os"
    "path/filepath"
    "testing"
)

var updateSnapshots = flag.Bool("update-snapshots", false, "update snapshot files")

func snapshotName(t *testing.T) string {
    t.Helper()
    return filepath.Join("testdata", "snapshots", t.Name()+".snap")
}

func assertSnapshot(t *testing.T, data any) {
    t.Helper()
    got, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        t.Fatal(err)
    }

    snapFile := snapshotName(t)
    if *updateSnapshots {
        os.MkdirAll(filepath.Dir(snapFile), 0755)
        if err := os.WriteFile(snapFile, got, 0644); err != nil {
            t.Fatal(err)
        }
        return
    }

    want, err := os.ReadFile(snapFile)
    if err != nil {
        t.Fatalf("snapshot %s not found (run with -update-snapshots to create): %v",
            snapFile, err)
    }

    if !bytes.Equal(got, want) {
        t.Errorf("snapshot mismatch for %s\n  got:  %s\n  want: %s",
            t.Name(), string(got), string(want))
    }
}

func TestCreateUser_Snapshot(t *testing.T) {
    reqBody := CreateUserRequest{Email: "snapshot@example.com", Name: "Snapshot"}
    req := httptest.NewRequest(http.MethodPost, "/users",
        bytes.NewReader(mustMarshal(t, reqBody)))
    rec := httptest.NewRecorder()

    CreateUserHandler(rec, req)

    var resp UserResponse
    mustDecode(t, rec.Body, &resp)

    assertSnapshot(t, resp)
}

func mustMarshal(t *testing.T, v any) []byte {
    t.Helper()
    data, _ := json.Marshal(v)
    return data
}

func mustDecode(t *testing.T, body *bytes.Buffer, v any) {
    t.Helper()
    json.NewDecoder(body).Decode(v)
}
```

```bash
# Create / update snapshots
go test -update-snapshots ./...

# Compare against snapshots
go test ./...

# Output
=== RUN   TestCreateUser_Success
--- PASS: TestCreateUser_Success (0.00s)
=== RUN   TestCreateUser_EmptyEmail
--- PASS: TestCreateUser_EmptyEmail (0.00s)
=== RUN   TestCreateUser_InvalidJSON
--- PASS: TestCreateUser_InvalidJSON (0.00s)
=== RUN   TestCreateUser_WrongMethod
--- PASS: TestCreateUser_WrongMethod (0.00s)
=== RUN   TestCreateUser_Snapshot
--- PASS: TestCreateUser_Snapshot (0.00s)
PASS
ok      api     0.130s
```

### Comparing JSON with ApproxMatch

When responses contain timestamps, UUIDs, or random values, use a comparison function that ignores dynamic fields:

```go
// File: approx_test.go
package api

import (
    "encoding/json"
    "testing"
    "time"
)

type TimestampedResponse struct {
    ID        int       `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    Status    string    `json:"status"`
}

func assertMatches(t *testing.T, got, want map[string]any, ignore []string) {
    t.Helper()
    ignoreSet := make(map[string]bool, len(ignore))
    for _, k := range ignore {
        ignoreSet[k] = true
    }

    for k, v := range want {
        if ignoreSet[k] {
            continue
        }
        gotV, ok := got[k]
        if !ok {
            t.Errorf("key %q missing from response", k)
            continue
        }
        if gotV != v {
            t.Errorf("key %q: got %v (%T), want %v (%T)", k, gotV, gotV, v, v)
        }
    }
}

func TestApproximateMatch(t *testing.T) {
    respBytes := []byte(`{"id":1,"created_at":"2026-05-30T10:00:00Z","status":"active"}`)

    var got map[string]any
    json.Unmarshal(respBytes, &got)

    want := map[string]any{
        "status": "active",
    }

    assertMatches(t, got, want, []string{"id", "created_at"})
}
```

### Best Practices

- Use `json.Decoder` for streaming, `json.Unmarshal` for pre-read byte slices.
- Write helpers (`marshalJSON`, `decodeJSON`) that fail the test immediately on bad data.
- For snapshot testing, use a `-update` flag so snapshots are created and maintained deliberately.
- Never compare raw JSON strings -- field ordering is not guaranteed by the spec. Use structured comparison.
- Use `assertMatches` or `approve`-style libraries for partial matching on dynamic responses.

---

## 12. Continuous Testing

### Overview

Continuous testing automatically re-runs tests when source files change. Combined with CI pipelines that enforce test quality, this gives rapid feedback during development and safety before merging.

### File Watchers

#### reflex (simple, fast)

Install:
```bash
go install github.com/cespare/reflex@latest
```

```bash
# Run tests whenever a .go file changes in the current directory tree
reflex -g '*.go' -- sh -c 'go test ./...'

# With race detection
reflex -g '*.go' -- sh -c 'go test -race ./...'

# Run a specific package
reflex -g 'server/*.go' -- sh -c 'go test ./server/...'
```

Output while running:
```
[01] Starting: sh -c go test ./...
ok      user    0.102s
[01] Starting: sh -c go test ./...
ok      handler 0.115s
[01] Starting: sh -c go test ./...
--- FAIL: TestCreateUser_EmptyEmail (0.00s)
```

#### entr (Unix staple)

```bash
# Requires fd or find
fd --type f -e go | entr -c go test ./...
```

#### air (with live reload)

Install:
```bash
go install github.com/air-verse/air@latest
```

```go
// File: .air.toml
root = "."
testdata_dir = "testdata"
[build]
  cmd = "go test -race -coverprofile=coverage.out ./..."
  bin = "true"
  full_bin = ""
  include_ext = ["go"]
  exclude_dir = ["testdata", "vendor"]
  delay = 1000
```

```bash
air
```

### CI Patterns

#### GitHub Actions

```yaml
# File: .github/workflows/test.yml
name: Tests

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  test:
    strategy:
      matrix:
        go-version: ["1.22", "1.23"]
        os: [ubuntu-latest, macos-latest]
    runs-on: ${{ matrix.os }}

    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: ${{ matrix.go-version }}
          cache: true

      - name: Unit Tests
        run: go test -race -coverprofile=coverage.out -covermode=atomic ./...

      - name: Integration Tests
        run: go test -tags=integration -race -coverprofile=integration.cov ./...

      - name: Upload Coverage
        uses: codecov/codecov-action@v4
        with:
          file: ./coverage.out
          flags: unittests
          fail_ci_if_error: true

      - name: Check Coverage Threshold
        run: |
          COVER=$(go tool cover -func=coverage.out | grep total | awk '{print $NF}' | tr -d '%')
          if [ "$COVER" -lt 80 ]; then
            echo "Coverage ${COVER}% is below threshold 80%"
            exit 1
          fi
          echo "Coverage ${COVER}% meets threshold"
```

#### GitLab CI

```yaml
# File: .gitlab-ci.yml
stages:
  - test

variables:
  GO_VERSION: "1.23"

.go-cache:
  variables:
    GOPATH: $CI_PROJECT_DIR/.go
  cache:
    paths:
      - .go/pkg/mod/

unit-tests:
  stage: test
  image: golang:$GO_VERSION
  extends: .go-cache
  script:
    - go test -race -coverprofile=coverage.out -covermode=atomic ./...
    - go tool cover -func=coverage.out
  coverage: '/total:\s+\(statements\)\s+(\d+\.?\d*%)/'

integration-tests:
  stage: test
  image: golang:$GO_VERSION
  services:
    - docker:dind
  variables:
    DOCKER_HOST: tcp://docker:2375
  extends: .go-cache
  script:
    - go test -tags=integration -v -timeout 120s ./...
```

#### Makefile

```makefile
# File: Makefile
.PHONY: test test-race test-all test-watch coverage

# Run all unit tests
test:
	go test -count=1 ./...

# Run with race detection
test-race:
	go test -race -count=1 ./...

# Run everything including integration tests
test-all:
	go test -tags=integration -race -count=1 ./...

# Continuous watch (requires reflex)
test-watch:
	reflex -g '*.go' -- go test ./...

# Coverage report
coverage:
	go test -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html

# Coverage threshold enforcement
coverage-check:
	go test -coverprofile=coverage.out -covermode=atomic ./...
	@COVER=$$(go tool cover -func=coverage.out | grep total | awk '{print $$NF}' | tr -d '%'); \
	echo "Coverage: $$COVER%"; \
	if [ "$$COVER" -lt 80 ]; then \
		echo "FAIL: coverage below 80%"; \
		exit 1; \
	fi
```

### Build Tags and CI

```makefile
# Run only short (fast) tests in CI during early stages
ci-fast:
	go test -short -race -count=1 ./...

# Run full suite including integration tests that need Docker
ci-full:
	go test -tags=integration -race -count=1 -timeout 180s ./...
```

### Best Practices

- Run `-race` in CI on at least one platform (Linux is usually fastest).
- Separate unit and integration tests with build tags and run them in different CI stages.
- Use `count=1` in CI to disable test caching (`go test -count=1`).
- Set reasonable `-timeout` values: 30s for unit, 180s for integration.
- Upload coverage reports to a service (Codecov, Coveralls) to track trends over time.
- Use a Makefile or Taskfile to standardize test commands across the team.

---

## Appendix: Comparing Testing Packages

| Feature | stdlib `testing` | testify | ginkgo | gomega |
|---|---|---|---|---|
| Test functions | `func TestX(t *testing.T)` | Same | `Describe/It` blocks | Matcher assertions |
| Assertions | Manual `if got != want` | `assert.Equal`, `require.NoError` | N/A | `Expect(got).To(Equal(want))` |
| Mocks | Hand-rolled | `mock.Mock` | gomock | Counterpart |
| Suites | Subtrees with `t.Run` | `suite.Suite` | `Describe/Context/It` | N/A |
| Popularity | All Go projects | Very common | Common in large projects | Paired with ginkgo |

### When to Use What

- **stdlib `testing`**: Always the foundation. Use for simple tests, table-driven tests, benchmarks, fuzzing.
- **testify**: Add for `assert`/`require` (reduces boilerplate) and `mock` (structured mocking). This is the sweet spot for most projects.
- **ginkgo/gomega**: Consider for BDD-style tests or very large suites where `Describe/It` readability helps. Adds a DSL and a separate test runner.

---

## Quick Reference

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run with race detection
go test -race ./...

# Run specific test
go test -run TestName ./...

# Run only unit tests (skip integration)
go test -short ./...

# Run integration tests only
go test -tags=integration ./...

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run fuzz test
go test -fuzz FuzzName -fuzztime 30s ./...

# Run tests continuously (with reflex)
reflex -g '*.go' -- go test ./...

# Update golden files
go test -update ./...

# Update snapshots
go test -update-snapshots ./...
```
