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
