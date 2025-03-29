package printgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/printgo"
)

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

func TestPTX_Fprintf(t *testing.T) {
	ptx := printgo.NewPTX()
	ptx.Fprintf("yyle%s", "88")
	require.Equal(t, "yyle88", ptx.String())
}

func TestPTS_Fprintf(t *testing.T) {
	pts := printgo.NewPTS()
	pts.Fprintf("%s88", "yyle")
	require.Equal(t, "yyle88", pts.String())
}

func TestPTX_Printf(t *testing.T) {
	ptx := printgo.NewPTX()
	ptx.Printf("print%s", "go")
	require.Equal(t, "printgo", ptx.String())
}

func TestPTS_Printf(t *testing.T) {
	pts := printgo.NewPTS()
	pts.Printf("%sgo", "print")
	require.Equal(t, "printgo", pts.String())
}
