# 🎯 Go Learning Journey

A comprehensive, structured learning path for mastering Go from beginner to intermediate level. This repository contains organized modules, practice exercises, and real-world projects.

## 📋 Table of Contents
- [Project Structure](#project-structure)
- [Learning Roadmap](#learning-roadmap)
- [Prerequisites](#prerequisites)
- [How to Use This Repository](#how-to-use-this-repository)
- [Learning Modules](#learning-modules)
- [Practice Exercises](#practice-exercises)
- [Resources & References](#resources--references)
- [Projects](#projects)
- [Progress Tracker](#progress-tracker)

---

## 🗂️ Project Structure

```
go-learning-journey/
├── 01-basics/                 # Fundamentals & Getting Started
├── 02-data-types/            # Variables, Types, Constants
├── 03-control-flow/          # If, Loops, Switch, Defer
├── 04-functions/             # Functions, Methods, Closures
├── 05-packages-modules/      # Packages, Imports, Go Modules
├── 06-interfaces/            # Interfaces, Structs, Pointers
├── 07-concurrency/           # Goroutines, Channels, Sync
├── 08-testing/               # Unit Testing, Benchmarks
├── 09-projects/              # Real-world Projects
├── exercises/                # Practice Exercises
├── resources/                # External Links & Materials
├── docs/                     # Learning Documentation
└── README.md                 # This file
```

---

## 🚀 Learning Roadmap

### Phase 1: Foundation (Week 1-2)
- [ ] **Module 01**: Basics & Setup
- [ ] **Module 02**: Data Types & Variables
- [ ] **Module 03**: Control Flow (If, Loops, Switch)
- **Estimated Time**: 10-14 hours

### Phase 2: Core Concepts (Week 3-4)
- [ ] **Module 04**: Functions & Methods
- [ ] **Module 05**: Packages & Modules
- [ ] **Module 06**: Structs & Interfaces
- **Estimated Time**: 12-16 hours

### Phase 3: Advanced Topics (Week 5-6)
- [ ] **Module 07**: Concurrency (Goroutines, Channels)
- [ ] **Module 08**: Testing & Debugging
- [ ] **Project 1**: Simple CLI Application
- **Estimated Time**: 14-18 hours

### Phase 4: Mastery Projects (Week 7-8)
- [ ] **Project 2**: Web Server
- [ ] **Project 3**: Data Processing Tool
- [ ] **Project 4**: REST API with Database
- **Estimated Time**: 20+ hours

**Total Estimated Duration**: 8 weeks with 3-4 hours daily commitment

---

## 📚 Prerequisites

### Before Starting
- [ ] Basic programming knowledge (any language)
- [ ] Go 1.21+ installed ([Installation Guide](https://golang.org/doc/install))
- [ ] Text editor or IDE (VS Code, GoLand, or Vim)
- [ ] Git installed for version control
- [ ] Command line/Terminal familiarity

### Recommended Setup
```bash
# Verify Go installation
go version

# Install useful tools
go install github.com/cosmtrek/air@latest      # Hot reload
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest  # Linter
```

---

## 💡 How to Use This Repository

### 1. Start with Module 01
Each module contains:
- `README.md` - Concept explanations
- `examples/` - Working code examples
- `exercises/` - Practice problems
- `solutions/` - Reference solutions

### 2. Follow the Pattern
```bash
cd 01-basics
cat README.md              # Read the theory
go run examples/*.go       # Run examples
# Write your code in your file
go run your_solution.go    # Test your solution
```

### 3. Track Your Progress
Update the checkboxes in this README as you complete each module.

### 4. Join the Community
- Go Forums: https://github.com/golang/go/discussions
- Gopher Community: https://gophers.slack.com

---

## 🎓 Learning Modules

### Module 01: Basics & Setup ⭐ START HERE
**Duration**: 2-3 hours | **Status**: [ ] Not Started

**Topics**:
- Go workspace and project structure
- Your first program (Hello World)
- Running and building Go code
- Go environment variables
- Comments and basic syntax

**Learning Materials**:
- 📖 [A Tour of Go - Welcome](https://tour.golang.org/welcome/1)
- 📖 [Effective Go](https://golang.org/doc/effective_go)
- 📺 [Go in 100 Seconds (YouTube)](https://www.youtube.com/watch?v=446E-r0rZY8)
- 📺 [Golang Tutorial for Beginners (YouTube)](https://www.youtube.com/watch?v=yyUHQIec83I)

**Hands-On Practice**:
- Create your first Go program
- Set up a project with go.mod
- Run code with `go run` and `go build`

---

### Module 02: Data Types & Variables
**Duration**: 3-4 hours | **Status**: [ ] Not Started

**Topics**:
- Variables and constants
- Basic data types (int, float, string, bool)
- Type conversion and assertions
- Arrays and slices
- Maps
- Strings and runes

**Learning Materials**:
- 📖 [A Tour of Go - Basics](https://tour.golang.org/basics/1)
- 📖 [Go Data Types](https://www.w3schools.com/go/go_data_types.php)
- 📹 [Go Variables and Data Types (YouTube)](https://www.youtube.com/watch?v=RnSdL1aXNl4)
- 📄 [Go Slices: usage and internals](https://golang.org/blog/slices-intro)

**Hands-On Practice**:
- Work with different data types
- Manipulate slices and maps
- Practice type conversions
- Explore strings and character encoding

---

### Module 03: Control Flow
**Duration**: 2-3 hours | **Status**: [ ] Not Started

**Topics**:
- If and else statements
- For loops (traditional, range, infinite)
- Switch statements
- Defer, panic, and recover
- Break and continue

**Learning Materials**:
- 📖 [A Tour of Go - Flow Control](https://tour.golang.org/flowcontrol/1)
- 📖 [Go If...Else & Loops (W3Schools)](https://www.w3schools.com/go/go_if_else.php)
- 📹 [Go Control Structures (YouTube)](https://www.youtube.com/watch?v=U7Vmy0FY3bc)
- 📄 [Defer, Panic, and Recover](https://golang.org/blog/defer-panic-and-recover)

**Hands-On Practice**:
- Write conditional logic
- Create loops for data iteration
- Use defer for resource cleanup
- Handle errors with defer

---

### Module 04: Functions & Methods
**Duration**: 3-4 hours | **Status**: [ ] Not Started

**Topics**:
- Function declaration and parameters
- Return values (single and multiple)
- Variadic functions
- Anonymous functions and closures
- Methods on types
- Named return values

**Learning Materials**:
- 📖 [A Tour of Go - Functions](https://tour.golang.org/basics/4)
- 📖 [Go Methods and Interfaces](https://golang.org/doc/tutorial/call-module-code)
- 📹 [Go Functions Tutorial (YouTube)](https://www.youtube.com/watch?v=j9OGBkZm_yk)
- 📄 [First-class functions in Go](https://golang.org/doc/codewalk/functions)

**Hands-On Practice**:
- Create functions with various signatures
- Return multiple values
- Write methods for custom types
- Explore closures and callbacks

---

### Module 05: Packages & Modules
**Duration**: 3-4 hours | **Status**: [ ] Not Started

**Topics**:
- Package system and organization
- Imports and exports (capitalization)
- Creating reusable packages
- Go Modules (go.mod, go.sum)
- Dependency management
- Public and private identifiers

**Learning Materials**:
- 📖 [Go Packages](https://golang.org/doc/code#Organization)
- 📖 [Using Go Modules](https://golang.org/doc/modules/managing-dependencies)
- 📹 [Go Packages and Modules (YouTube)](https://www.youtube.com/watch?v=vaQJW2lNmQI)
- 📄 [Package Management Best Practices](https://golang.org/doc/modules/managing-source)

**Hands-On Practice**:
- Create a multi-package project
- Manage external dependencies
- Write exportable functions and types
- Practice module versioning

---

### Module 06: Structs & Interfaces
**Duration**: 4-5 hours | **Status**: [ ] Not Started

**Topics**:
- Struct declaration and usage
- Struct fields and tags
- Embedding and composition
- Pointers and memory
- Interface definition and implementation
- Type assertions and switches
- Empty interface (`interface{}`)

**Learning Materials**:
- 📖 [A Tour of Go - Methods and Interfaces](https://tour.golang.org/methods/1)
- 📖 [Go Structs & Interfaces](https://www.w3schools.com/go/go_structs.php)
- 📹 [Structs and Interfaces in Go (YouTube)](https://www.youtube.com/watch?v=Y5KJQRO0gKE)
- 📄 [The Laws of Reflection](https://golang.org/blog/laws-of-reflection)

**Hands-On Practice**:
- Design complex data structures
- Implement interfaces
- Work with pointers and references
- Practice composition over inheritance

---

### Module 07: Concurrency
**Duration**: 4-5 hours | **Status**: [ ] Not Started

**Topics**:
- Goroutines
- Channels (unbuffered, buffered)
- Select statement
- Synchronization primitives (WaitGroup, Mutex)
- Context package
- Error handling in concurrent code
- Common patterns and pitfalls

**Learning Materials**:
- 📖 [A Tour of Go - Concurrency](https://tour.golang.org/concurrency/1)
- 📖 [Concurrency in Go](https://golang.org/doc/effective_go#concurrency)
- 📹 [Go Concurrency Patterns (YouTube)](https://www.youtube.com/watch?v=f6kdp27TYZs)
- 📄 [Go Concurrency Patterns](https://www.youtube.com/watch?v=DqHb5KBe7qI)
- 📄 [Context Package Guide](https://golang.org/blog/context)

**Hands-On Practice**:
- Create goroutines and manage them
- Communicate between goroutines
- Build concurrent data processors
- Handle race conditions

---

### Module 08: Testing & Debugging
**Duration**: 3-4 hours | **Status**: [ ] Not Started

**Topics**:
- Unit testing basics
- Test organization and naming
- Table-driven tests
- Benchmarking
- Test coverage
- Debugging techniques
- Profiling and optimization

**Learning Materials**:
- 📖 [Testing in Go](https://golang.org/doc/effective_go#testing)
- 📖 [Go Testing Package](https://golang.org/pkg/testing/)
- 📹 [Go Testing Tutorial (YouTube)](https://www.youtube.com/watch?v=ndmB0bHoZrU)
- 📄 [Subtests and Sub-benchmarks](https://golang.org/blog/subtests)

**Hands-On Practice**:
- Write unit tests for your code
- Create table-driven tests
- Run benchmarks
- Measure test coverage

---

## 🏋️ Practice Exercises

### Beginner Level
Located in `exercises/beginner/`:
1. **Hello World** - Your first program
2. **Calculator** - Basic arithmetic operations
3. **FizzBuzz** - Loop and conditional logic
4. **Temperature Converter** - Type conversion
5. **Grade Analyzer** - Working with functions

### Intermediate Level
Located in `exercises/intermediate/`:
1. **Todo List CLI** - Data structures and I/O
2. **File Processor** - File operations
3. **Chat Application** - Goroutines and channels
4. **Data Pipeline** - Concurrency patterns
5. **JSON Parser** - Encoding/decoding

### Advanced Level
Located in `exercises/advanced/`:
1. **Web Crawler** - HTTP, parsing, concurrency
2. **Database Manager** - CRUD operations
3. **REST API** - HTTP server implementation
4. **Monitoring Tool** - System metrics and logging
5. **Game Engine** - Complex state management

---

## 📦 Projects

### Project 1: CLI Todo Manager (Beginner)
**Duration**: 5-7 hours | **Status**: [ ] Not Started

**Features**:
- Add, list, and delete todos
- Save todos to file
- Mark todos as complete
- Pretty terminal output

**Technologies**: File I/O, CLI flags, JSON
**Learning**: Functions, structs, file operations

**Project Directory**: `09-projects/01-todo-cli/`

---

### Project 2: Weather CLI App (Intermediate)
**Duration**: 8-10 hours | **Status**: [ ] Not Started

**Features**:
- Fetch weather from API (OpenWeatherMap)
- Display formatted weather data
- Cache results
- Support multiple locations

**Technologies**: HTTP client, JSON parsing, error handling
**Learning**: External APIs, JSON, error handling

**Project Directory**: `09-projects/02-weather-cli/`

---

### Project 3: Simple Web Server (Intermediate)
**Duration**: 10-12 hours | **Status**: [ ] Not Started

**Features**:
- HTTP server with routing
- Handle GET/POST requests
- Serve static files
- Basic middleware system

**Technologies**: net/http, goroutines, channels
**Learning**: Web fundamentals, concurrency

**Project Directory**: `09-projects/03-web-server/`

---

### Project 4: REST API (Intermediate-Advanced)
**Duration**: 15-20 hours | **Status**: [ ] Not Started

**Features**:
- RESTful API with CRUD operations
- Database integration (SQLite)
- Authentication (JWT tokens)
- Comprehensive error handling
- API documentation

**Technologies**: Chi router, SQLite, JWT, testing
**Learning**: Web development, databases, testing

**Project Directory**: `09-projects/04-rest-api/`

---

### Project 5: Concurrent Task Processor (Advanced)
**Duration**: 12-15 hours | **Status**: [ ] Not Started

**Features**:
- Process tasks concurrently with workers
- Task queue implementation
- Error handling and retries
- Progress tracking
- Performance metrics

**Technologies**: Goroutines, channels, sync, context
**Learning**: Advanced concurrency, design patterns

**Project Directory**: `09-projects/05-task-processor/`

---

## 📚 Resources & References

### Official Go Resources
- 🌐 [golang.org](https://golang.org/) - Official Go website
- 📖 [A Tour of Go](https://tour.golang.org/) - Interactive tutorial (START HERE!)
- 📖 [Go Documentation](https://golang.org/doc/)
- 📖 [Effective Go](https://golang.org/doc/effective_go)
- 📖 [Standard Library](https://golang.org/pkg/)

### Comprehensive Learning Platforms
- 🎓 [Go by Example](https://gobyexample.com/) - Practical examples
- 🎓 [W3Schools Go Tutorial](https://www.w3schools.com/go/)
- 🎓 [Udemy - Go Mastery Course](https://www.udemy.com/course/go-the-complete-developers-guide/)
- 🎓 [Pluralsight - Go Learning Path](https://www.pluralsight.com/paths/getting-started-with-go)

### YouTube Channels
- 📺 [Traversy Media](https://www.youtube.com/c/TraversyMedia) - Go tutorials
- 📺 [TechWorld with Nana](https://www.youtube.com/c/TechWorldwithNana) - DevOps & Go
- 📺 [freeCodeCamp.org](https://www.youtube.com/c/freeCodeCamp) - Complete Go course
- 📺 [Fireship.io](https://www.youtube.com/c/Fireship) - Quick overviews

### Books
- 📕 [The Go Programming Language](https://www.gopl.io/) - Definitive resource
- 📕 [Go in Action](https://www.manning.com/books/go-in-action)
- 📕 [Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/)
- 📕 [Cloud Native Go](https://learning.oreilly.com/library/view/cloud-native-go/9781492076322/)

### Practice & Challenges
- 🎯 [LeetCode Go Problems](https://leetcode.com/discuss/general-discussion/460599/blind-75-leetcode-questions)
- 🎯 [HackerRank Go Challenges](https://www.hackerrank.com/domains/go)
- 🎯 [Project Euler](https://projecteuler.net/) - Math problems in Go
- 🎯 [Exercism.io Go Track](https://exercism.org/tracks/go)

### Blogs & Articles
- 📄 [Go Blog](https://golang.org/blog/) - Official announcements and insights
- 📄 [Dave Cheney's Blog](https://dave.cheney.net/) - Advanced Go topics
- 📄 [Uber's Go Guide](https://github.com/uber-go/guide) - Style guide
- 📄 [Gopher Academy](https://gopheracademy.com/) - Community content

### Tools & Utilities
- 🔧 [Go Playground](https://play.golang.org/) - Online code editor
- 🔧 [GoLang Bot](https://gobot.io/) - IoT framework
- 🔧 [VS Code Go Extension](https://github.com/golang/vscode-go)
- 🔧 [golangci-lint](https://golangci-lint.run/) - Linter

### Community & Forums
- 💬 [Go Reddit](https://www.reddit.com/r/golang/)
- 💬 [Go Gopher Slack](https://gophers.slack.com/join)
- 💬 [Stack Overflow Go Tag](https://stackoverflow.com/questions/tagged/go)
- 💬 [GitHub Discussions](https://github.com/golang/go/discussions)

---

## ✅ Progress Tracker

### Modules Progress
| Module | Status | Completed | Notes |
|--------|--------|-----------|-------|
| 01 Basics | ⭕ Not Started | - | - |
| 02 Data Types | ⭕ Not Started | - | - |
| 03 Control Flow | ⭕ Not Started | - | - |
| 04 Functions | ⭕ Not Started | - | - |
| 05 Packages | ⭕ Not Started | - | - |
| 06 Interfaces | ⭕ Not Started | - | - |
| 07 Concurrency | ⭕ Not Started | - | - |
| 08 Testing | ⭕ Not Started | - | - |

### Projects Progress
| Project | Status | Completed | Notes |
|---------|--------|-----------|-------|
| 1. Todo CLI | ⭕ Not Started | - | - |
| 2. Weather CLI | ⭕ Not Started | - | - |
| 3. Web Server | ⭕ Not Started | - | - |
| 4. REST API | ⭕ Not Started | - | - |
| 5. Task Processor | ⭕ Not Started | - | - |

---

## 🎯 Getting Started

### Step 1: Set Up Your Environment
```bash
# Clone this repository
git clone https://github.com/YOUR_USERNAME/go-learning-journey.git
cd go-learning-journey

# Create your first branch
git checkout -b learning/start
```

### Step 2: Start with Module 01
```bash
cd 01-basics
cat README.md
go run examples/hello-world.go
```

### Step 3: Complete Exercises
```bash
# Complete exercises in the module
# Then commit your work
git add .
git commit -m "complete: module 01 basics"
git push origin learning/start
```

### Step 4: Organize Your Learning
- Track your progress in the Progress Tracker table above
- Update this README as you complete modules
- Keep notes in `docs/` directory

---

## 🤝 Contributing to Your Learning

Tips for effective learning:
1. **Code Along**: Don't just read, write code!
2. **Experiment**: Modify examples and see what happens
3. **Build Projects**: Apply concepts in real projects
4. **Join Communities**: Connect with other learners
5. **Teach Others**: Explain concepts to solidify understanding
6. **Review Others' Code**: Learn from different approaches
7. **Take Breaks**: Pace yourself, avoid burnout

---

## 📞 Support & Questions

If you get stuck:
1. Check the relevant module's README
2. Review the examples in that module
3. Search on Stack Overflow or GitHub Issues
4. Ask in Go community forums
5. Revisit fundamental concepts

---

## 📄 License

This learning repository is open source and available under the MIT License.

---

## 🎉 Acknowledgments

This learning journey is inspired by:
- Official Go documentation
- Community best practices
- Real-world Go applications
- Feedback from Go developers

---

**Last Updated**: May 30, 2026  
**Status**: 🟢 Active Learning Repository  
**Next Step**: Start with [Module 01 Basics](./01-basics/README.md)

---

*Happy Learning! 🚀 Remember: Consistency is key. Code a little every day!*
