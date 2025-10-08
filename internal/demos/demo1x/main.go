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
