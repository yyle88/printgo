[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/printgo/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/printgo/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/printgo)](https://pkg.go.dev/github.com/yyle88/printgo)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/printgo/main.svg)](https://coveralls.io/github/yyle88/printgo?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yyle88/printgo.svg)](https://github.com/yyle88/printgo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/printgo)](https://goreportcard.com/report/github.com/yyle88/printgo)

# printgo

`printgo` lets you print strings piece-by-piece, then get the whole text at once.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

📝 **Content Accumulation** - Print strings piece-by-piece, get whole text at once
🔄 **Two Implementations** - PTX (bytes.Buffer) and PTS (strings.Builder)
✨ **Fmt-Style Methods** - Print, Println, Printf, Fprintf with known interface
⚡ **Auto Exception Handling** - Integrated exception handling with done/must/rese packages
🎯 **Simple API** - Quick to use with compact code

## Installation

```bash
go get github.com/yyle88/printgo
```

## Usage

### Basic PTX Example

```go
package main

import (
	"fmt"

	"github.com/yyle88/printgo"
)

func main() {
	// Create PTX based on bytes.Buffer
	ptx := printgo.NewPTX()

	// Print function piece-by-piece
	ptx.Println("func Add(a, b int) int {")
	ptx.Println("\treturn a + b")
	ptx.Println("}")

	// Get the complete function code
	result := ptx.String()
	fmt.Println(result)
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Basic PTS Example

```go
package main

import (
	"fmt"

	"github.com/yyle88/printgo"
)

func main() {
	// Create PTS based on strings.Builder
	pts := printgo.NewPTS()

	// Build struct definition piece-by-piece
	pts.Println("type Person struct {")
	pts.Printf("\tName string\n")
	pts.Printf("\tAge  int\n")
	pts.Println("}")

	// Get the complete struct code
	result := pts.String()
	fmt.Println(result)
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

## Examples

### PTX vs PTS

**PTX (bytes.Buffer):**

```go
ptx := printgo.NewPTX()
ptx.Println("Using bytes.Buffer")
```

**PTS (strings.Builder):**

```go
pts := printgo.NewPTS()
pts.Println("Using strings.Builder")
```

### Format Printing

**Printf:**

```go
ptx := printgo.NewPTX()
ptx.Printf("Name: %s, Age: %d\n", "Alice", 30)
```

**Fprintf:**

```go
pts := printgo.NewPTS()
pts.Fprintf("Total: %.2f", 123.456)
```

**Printfln (format + newline):**

```go
ptx := printgo.NewPTX()
ptx.Printfln("Name: %s", "test")
ptx.Printfln("Age: %d", 18)
```

## API Reference

### PTX (bytes.Buffer based)

| Method                       | Description                   |
| ---------------------------- | ----------------------------- |
| `NewPTX()`                   | Create new PTX instance       |
| `Print(args...)`             | Print without newline         |
| `Println(args...)`           | Print with newline            |
| `Printf(format, args...)`    | Format and print              |
| `Fprintf(format, args...)`   | Format and print              |
| `Printfln(format, args...)`  | Format and print with newline |
| `Fprintfln(format, args...)` | Format and print with newline |
| `String()`                   | Get accumulated text          |
| `Bytes()`                    | Get accumulated bytes         |

### PTS (strings.Builder based)

| Method                       | Description                   |
| ---------------------------- | ----------------------------- |
| `NewPTS()`                   | Create new PTS instance       |
| `Print(args...)`             | Print without newline         |
| `Println(args...)`           | Print with newline            |
| `Printf(format, args...)`    | Format and print              |
| `Fprintf(format, args...)`   | Format and print              |
| `Printfln(format, args...)`  | Format and print with newline |
| `Fprintfln(format, args...)` | Format and print with newline |
| `String()`                   | Get accumulated text          |

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yyle88/printgo.svg?variant=adaptive)](https://starchart.cc/yyle88/printgo)
