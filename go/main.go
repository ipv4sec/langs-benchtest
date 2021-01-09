package main

import "fmt"

// GOOS=js GOARCH=wasm go build -o zed.wasm
func main() {
	fmt.Println("Hello, Go")
}
