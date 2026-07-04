package main

import (
	"cad"
)

func main() {
	// Test cube
	cube := cad.Cube(10)
	cube.Export("example2.stl")
}
