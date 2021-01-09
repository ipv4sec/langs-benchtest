package main

import (
	"fmt"
	"runtime"
)

// GOOS=js GOARCH=wasm go build -o main.wasm
func main() {
	runtime.GOOS
	fmt.Println("Hello, Go")
}
