package main

import (
	"fmt"
	"runtime"
)

//main func
func main() {
	fmt.Printf("runtime: %s\narchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}
