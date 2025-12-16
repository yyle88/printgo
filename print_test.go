package printgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/printgo"
)

// TestPTX_Println tests PTX Println method with mixed content
// TestPTX_Println 测试 PTX 的 Println 方法处理混合内容
func TestPTX_Println(t *testing.T) {
	ptx := printgo.NewPTX()

	ptx.Println("a")
	ptx.Print("lele")
	ptx.Print("8")
	ptx.Print("8")
	ptx.Print("")
	ptx.Println()
	ptx.Println("b")
	ptx.Println("c")

	res := ptx.String()
	t.Log(res)

	require.Equal(t, "a"+"\n"+"lele88"+"\n"+"b"+"\n"+"c"+"\n", res)
}

// TestPTS_Println tests PTS Println method with mixed content
// TestPTS_Println 测试 PTS 的 Println 方法处理混合内容
func TestPTS_Println(t *testing.T) {
	pts := printgo.NewPTS()

	pts.Println("a")
	pts.Print("lele")
	pts.Print("8")
	pts.Print("8")
	pts.Print("")
	pts.Println()
	pts.Println("b")
	pts.Println("c")

	res := pts.String()
	t.Log(res)

	require.Equal(t, "a"+"\n"+"lele88"+"\n"+"b"+"\n"+"c"+"\n", res)
}

// TestPTX_Print tests PTX Print method concatenation
// TestPTX_Print 测试 PTX 的 Print 方法拼接功能
func TestPTX_Print(t *testing.T) {
	ptx := printgo.NewPTX()
	ptx.Print("print")
	ptx.Print("go")
	require.Equal(t, "printgo", ptx.String())
}

// TestPTS_Print tests PTS Print method concatenation
// TestPTS_Print 测试 PTS 的 Print 方法拼接功能
func TestPTS_Print(t *testing.T) {
	pts := printgo.NewPTS()
	pts.Print("print")
	pts.Print("go")
	require.Equal(t, "printgo", pts.String())
}

// TestPTX_Fprintf tests PTX Fprintf format string
// TestPTX_Fprintf 测试 PTX 的 Fprintf 格式化字符串
func TestPTX_Fprintf(t *testing.T) {
	ptx := printgo.NewPTX()
	ptx.Fprintf("yyle%s", "88")
	require.Equal(t, "yyle88", ptx.String())
}

// TestPTS_Fprintf tests PTS Fprintf format string
// TestPTS_Fprintf 测试 PTS 的 Fprintf 格式化字符串
func TestPTS_Fprintf(t *testing.T) {
	pts := printgo.NewPTS()
	pts.Fprintf("%s88", "yyle")
	require.Equal(t, "yyle88", pts.String())
}

// TestPTX_Printf tests PTX Printf format string
// TestPTX_Printf 测试 PTX 的 Printf 格式化字符串
func TestPTX_Printf(t *testing.T) {
	ptx := printgo.NewPTX()
	ptx.Printf("print%s", "go")
	require.Equal(t, "printgo", ptx.String())
}

// TestPTS_Printf tests PTS Printf format string
// TestPTS_Printf 测试 PTS 的 Printf 格式化字符串
func TestPTS_Printf(t *testing.T) {
	pts := printgo.NewPTS()
	pts.Printf("%sgo", "print")
	require.Equal(t, "printgo", pts.String())
}

// TestPTX_Printfln tests PTX Printfln format with newline
// TestPTX_Printfln 测试 PTX 的 Printfln 格式化并换行
func TestPTX_Printfln(t *testing.T) {
	ptx := printgo.NewPTX()
	ptx.Printfln("Name: %s", "test")
	ptx.Printfln("Age: %d", 18)
	require.Equal(t, "Name: test\nAge: 18\n", ptx.String())
}

// TestPTS_Printfln tests PTS Printfln format with newline
// TestPTS_Printfln 测试 PTS 的 Printfln 格式化并换行
func TestPTS_Printfln(t *testing.T) {
	pts := printgo.NewPTS()
	pts.Printfln("Name: %s", "test")
	pts.Printfln("Age: %d", 18)
	require.Equal(t, "Name: test\nAge: 18\n", pts.String())
}

// TestPTX_Fprintfln tests PTX Fprintfln format with newline
// TestPTX_Fprintfln 测试 PTX 的 Fprintfln 格式化并换行
func TestPTX_Fprintfln(t *testing.T) {
	ptx := printgo.NewPTX()
	ptx.Fprintfln("Name: %s", "test")
	ptx.Fprintfln("Age: %d", 18)
	require.Equal(t, "Name: test\nAge: 18\n", ptx.String())
}

// TestPTS_Fprintfln tests PTS Fprintfln format with newline
// TestPTS_Fprintfln 测试 PTS 的 Fprintfln 格式化并换行
func TestPTS_Fprintfln(t *testing.T) {
	pts := printgo.NewPTS()
	pts.Fprintfln("Name: %s", "test")
	pts.Fprintfln("Age: %d", 18)
	require.Equal(t, "Name: test\nAge: 18\n", pts.String())
}
