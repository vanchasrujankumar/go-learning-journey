# Setup Guide - Go Learning Journey

## 🚀 Quick Start (5 minutes)

### Prerequisites
- [ ] Go 1.21+ installed ([Download](https://golang.org/dl/))
- [ ] Git installed
- [ ] VS Code or GoLand IDE (optional)
- [ ] Terminal/Command prompt access

### Installation Steps

#### 1. Clone the Repository
```bash
# Clone to your home directory
cd ~
git clone https://github.com/YOUR_USERNAME/go-learning-journey.git
cd go-learning-journey
```

#### 2. Initialize Git (First Time Only)
```bash
# Configure git locally for this repo
git config user.name "Your Name"
git config user.email "your.email@example.com"
```

#### 3. Install Development Tools
```bash
make setup
# or manually:
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install golang.org/x/tools/cmd/goimports@latest
```

#### 4. Verify Installation
```bash
go version
make help
```

## 📁 Project Structure

```
go-learning-journey/
├── 01-basics/           ← Start here!
├── 02-data-types/
├── ...
├── 09-projects/
├── Makefile            ← Helpful commands
├── renovate.json       ← Auto-dependency updates
├── docker-compose.yml  ← Optional Docker setup
└── README.md           ← Learning roadmap
```

## 🎯 First Steps

### 1. Read the Main README
```bash
cat README.md
```

### 2. Start Module 01
```bash
cd 01-basics
cat README.md          # Read concepts
go run examples/01-hello-world.go  # Run first example
```

### 3. Run Your First Example
```bash
make run module=01-basics example=01-hello-world
```

### 4. Complete First Exercise
```bash
cd 01-basics
# Create your_solution.go based on exercises/
go run your_solution.go
```

## 🛠️ Available Commands

### View All Commands
```bash
make help
```

### Common Tasks

| Command | Purpose |
|---------|---------|
| `make run module=01-basics example=01-hello-world` | Run specific example |
| `make test` | Run all tests |
| `make fmt` | Format Go code |
| `make lint` | Run code linter |
| `make clean` | Remove build artifacts |
| `make coverage` | Generate coverage report |
| `make docker-up` | Start Docker services |

## 🔧 IDE Setup

### VS Code
1. Install **Go** extension by golang
2. Install **Prettier** extension
3. Settings are pre-configured in `.vscode/settings.json`
4. Open workspace: File → Open Folder → select project

### GoLand / IntelliJ IDEA
1. Open project: File → Open → select project
2. Settings auto-configured in `.idea/`
3. Enable Go modules if prompted

### Vim / Neovim
1. Install vim-go plugin
2. Run `:GoInstallBinaries`
3. Use `:GoRun` to execute files

## 📦 Docker Setup (Optional)

### Start Services
```bash
make docker-up
```

### Check Running Services
```bash
docker-compose ps
```

### View Logs
```bash
make docker-logs
```

### Stop Services
```bash
make docker-down
```

## 🚦 Dependency Management

### Renovate Auto-Updates
- ✅ Automatically merges dependency updates
- 📧 You won't get emails for routine updates
- ⚠️ Major updates require manual approval
- 🔐 Security updates auto-merged

See `docs/RENOVATE_GUIDE.md` for details.

## ✅ Verification Checklist

After setup, verify everything works:

```bash
# 1. Go is installed
go version

# 2. Git is configured
git config --list | grep user

# 3. Go tools installed
which golangci-lint

# 4. Project structure
ls -la 01-basics/

# 5. Run test example
make run module=01-basics example=01-hello-world
```

Expected output:
```
Hello, World!
```

## 🐛 Troubleshooting

### Problem: "go: command not found"
**Solution**: Go is not installed or not in PATH
```bash
# Install Go: https://golang.org/dl/
# Add to PATH: ~/.bash_profile or ~/.zshrc
export PATH=$PATH:/usr/local/go/bin
```

### Problem: "module not found"
**Solution**: Make sure you're in the project directory
```bash
cd go-learning-journey
pwd  # Should show the project path
```

### Problem: "make: command not found"
**Solution**: Install make (macOS/Linux)
```bash
# macOS
brew install make

# Ubuntu/Debian
sudo apt-get install build-essential

# Or use the commands directly from Makefile
```

### Problem: "permission denied" on files
**Solution**: Fix file permissions
```bash
chmod +x *.sh
```

### Problem: VS Code not recognizing Go
**Solution**: Restart VS Code and reinstall Go extension
```bash
code --install-extension golang.Go
```

## 📚 Learning Tips

1. **Start with Module 01**: Don't skip fundamentals
2. **Run Examples First**: Before reading solutions
3. **Modify Examples**: Change code and see what happens
4. **Take Notes**: Document your learning
5. **Commit Often**: `git add . && git commit -m "learning: topic"`
6. **Join Community**: Go Gophers Slack, Reddit r/golang

## 🔗 Useful Links

- [Go Official Site](https://golang.org/)
- [A Tour of Go](https://tour.golang.org/)
- [Go by Example](https://gobyexample.com/)
- [Stack Overflow Go Tag](https://stackoverflow.com/questions/tagged/go)

## 📞 Getting Help

### Within IDE
- VS Code: Install Go extension (built-in documentation)
- GoLand: Press `Ctrl+Shift+A` to search docs

### Online Resources
```bash
# Go built-in help
go help command

# Documentation
go doc fmt.Println

# Online
- https://pkg.go.dev/
- https://golang.org/ref/spec
```

### Ask Community
- Stack Overflow: https://stackoverflow.com/questions/tagged/go
- Reddit: https://www.reddit.com/r/golang/
- Slack: https://gophers.slack.com/

## 🎓 Next Steps After Setup

1. ✅ Read `README.md` (5 min)
2. ✅ Complete Module 01 (2-3 hours)
3. ✅ Complete Module 02 (3-4 hours)
4. ✅ Build Mini Project (5-7 hours)
5. ✅ Complete remaining modules (6+ weeks)

## 🤝 Contributing

Want to improve this repository?
- See `CONTRIBUTING.md`
- Pull requests welcome!
- Report issues on GitHub

---

**Setup Complete!** 🎉

You're ready to start learning Go. Begin with:
```bash
cd 01-basics
cat README.md
```

Happy Learning! 🚀

Last Updated: May 30, 2026
