// Package printgo lets you print strings piece-by-piece, then get the whole text at once
// Wraps bytes.Buffer and strings.Builder with fmt-style print functions
// Use Print/Println/Printf to write content, get result with String() / Bytes()
//
// printgo 包让你逐个打印字符串，最后一次性获取完整文本
// 包装 bytes.Buffer 和 strings.Builder，提供 fmt 风格的打印函数
// 使用 Print/Println/Printf 写入内容，通过 String() 或 Bytes() 获取结果
package printgo

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/yyle88/done"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
)

// NewPTX creates PTX based on bytes.Buffer
// Returns new instance to use with print operations
//
// NewPTX 创建基于 bytes.Buffer 的 PTX
// 返回可用于打印操作的新实例
func NewPTX() *PTX {
	return &PTX{}
}

// PTX is bytes.Buffer-based type with fmt-style methods
// Accumulates printed content and provides String method to get output
//
// PTX 是基于 bytes.Buffer 的类型，具有 fmt 风格的方法
// 累积打印内容并提供 String 方法获取输出
type PTX struct{ bytes.Buffer }

// Println prints args with newline appended
// Returns bytes written, panics on error
//
// Println 打印参数并追加换行符
// 返回写入的字节数，出错时 panic
func (ptx *PTX) Println(args ...interface{}) (n int) {
	return done.VE(fmt.Fprintln(ptx, args...)).Done()
}

// Print prints args without newline
// Returns bytes written, panics on error
//
// Print 打印参数不追加换行符
// 返回写入的字节数，出错时 panic
func (ptx *PTX) Print(args ...interface{}) (n int) {
	n, err := fmt.Fprint(ptx, args...)
	must.Done(err)
	return n
}

// Fprintf formats and prints with custom format string
// Returns bytes written, panics on error
//
// Fprintf 使用自定义格式字符串格式化并打印
// 返回写入的字节数，出错时 panic
func (ptx *PTX) Fprintf(format string, args ...interface{}) (n int) {
	return rese.V1(fmt.Fprintf(ptx, format, args...))
}

// Printf formats and prints with custom format string
// Returns bytes written, panics on error
//
// Printf 使用自定义格式字符串格式化并打印
// 返回写入的字节数，出错时 panic
func (ptx *PTX) Printf(format string, args ...interface{}) (n int) {
	return rese.V1(fmt.Fprintf(ptx, format, args...))
}

// NewPTS creates PTS based on strings.Builder
// Returns new instance to use with print operations
//
// NewPTS 创建基于 strings.Builder 的 PTS
// 返回可用于打印操作的新实例
func NewPTS() *PTS {
	return &PTS{}
}

// PTS is strings.Builder-based type with fmt-style methods
// Accumulates printed content and provides String method to get output
//
// PTS 是基于 strings.Builder 的类型，具有 fmt 风格的方法
// 累积打印内容并提供 String 方法获取输出
type PTS struct{ strings.Builder }

// Println prints args with newline appended
// Returns bytes written, panics on error
//
// Println 打印参数并追加换行符
// 返回写入的字节数，出错时 panic
func (pts *PTS) Println(args ...interface{}) (n int) {
	return rese.V1(fmt.Fprintln(pts, args...))
}

// Print prints args without newline
// Returns bytes written, panics on error
//
// Print 打印参数不追加换行符
// 返回写入的字节数，出错时 panic
func (pts *PTS) Print(args ...interface{}) (n int) {
	n, err := fmt.Fprint(pts, args...)
	must.Done(err)
	return n
}

// Fprintf formats and prints with custom format string
// Returns bytes written, panics on error
//
// Fprintf 使用自定义格式字符串格式化并打印
// 返回写入的字节数，出错时 panic
func (pts *PTS) Fprintf(format string, args ...interface{}) (n int) {
	return rese.V1(fmt.Fprintf(pts, format, args...))
}

// Printf formats and prints with custom format string
// Returns bytes written, panics on error
//
// Printf 使用自定义格式字符串格式化并打印
// 返回写入的字节数，出错时 panic
func (pts *PTS) Printf(format string, args ...interface{}) (n int) {
	return rese.V1(fmt.Fprintf(pts, format, args...))
}
