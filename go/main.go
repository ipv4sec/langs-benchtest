package main

import (
	"fmt"
)

// GOOS=js GOARCH=wasm go build -o main.wasm
func main() {
	fmt.Println("Hello, Go")
}
