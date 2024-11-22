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
	pts.Println()
	pts.Println("b")
	pts.Println("c")

	res := pts.String()
	t.Log(res)

	require.Equal(t, "a"+"\n"+"lele88"+"\n"+"b"+"\n"+"c"+"\n", res)
}
