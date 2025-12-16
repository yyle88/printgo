[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/printgo/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/printgo/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/printgo)](https://pkg.go.dev/github.com/yyle88/printgo)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/printgo/main.svg)](https://coveralls.io/github/yyle88/printgo?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yyle88/printgo.svg)](https://github.com/yyle88/printgo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/printgo)](https://goreportcard.com/report/github.com/yyle88/printgo)

# printgo

`printgo` 让你逐个打印字符串，最后一次性获取完整文本。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

📝 **内容累积** - 逐个打印字符串，最后一次性获取完整文本
🔄 **两种实现** - PTX (bytes.Buffer) 和 PTS (strings.Builder)
✨ **Fmt 风格方法** - Print、Println、Printf、Fprintf 等常见接口
⚡ **自动异常处理** - 集成 done/must/rese 包的异常处理
🎯 **简洁 API** - 快速使用，代码精简

## 安装

```bash
go get github.com/yyle88/printgo
```

## 使用方法

### PTX 基础示例

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

⬆️ **源码：** [源码](internal/demos/demo1x/main.go)

### PTS 基础示例

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

⬆️ **源码：** [源码](internal/demos/demo2x/main.go)

## 示例

### PTX vs PTS

**PTX (bytes.Buffer)：**

```go
ptx := printgo.NewPTX()
ptx.Println("Using bytes.Buffer")
```

**PTS (strings.Builder)：**

```go
pts := printgo.NewPTS()
pts.Println("Using strings.Builder")
```

### 格式化打印

**Printf：**

```go
ptx := printgo.NewPTX()
ptx.Printf("Name: %s, Age: %d\n", "Alice", 30)
```

**Fprintf：**

```go
pts := printgo.NewPTS()
pts.Fprintf("Total: %.2f", 123.456)
```

**Printfln（格式化 + 换行）：**

```go
ptx := printgo.NewPTX()
ptx.Printfln("Name: %s", "test")
ptx.Printfln("Age: %d", 18)
```

## API 列表

### PTX（基于 bytes.Buffer）

| 方法                         | 说明             |
| ---------------------------- | ---------------- |
| `NewPTX()`                   | 创建 PTX 实例    |
| `Print(args...)`             | 打印不换行       |
| `Println(args...)`           | 打印并换行       |
| `Printf(format, args...)`    | 格式化打印       |
| `Fprintf(format, args...)`   | 格式化打印       |
| `Printfln(format, args...)`  | 格式化打印并换行 |
| `Fprintfln(format, args...)` | 格式化打印并换行 |
| `String()`                   | 获取累积文本     |
| `Bytes()`                    | 获取累积字节     |

### PTS（基于 strings.Builder）

| 方法                         | 说明             |
| ---------------------------- | ---------------- |
| `NewPTS()`                   | 创建 PTS 实例    |
| `Print(args...)`             | 打印不换行       |
| `Println(args...)`           | 打印并换行       |
| `Printf(format, args...)`    | 格式化打印       |
| `Fprintf(format, args...)`   | 格式化打印       |
| `Printfln(format, args...)`  | 格式化打印并换行 |
| `Fprintfln(format, args...)` | 格式化打印并换行 |
| `String()`                   | 获取累积文本     |

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yyle88/printgo.svg?variant=adaptive)](https://starchart.cc/yyle88/printgo)
