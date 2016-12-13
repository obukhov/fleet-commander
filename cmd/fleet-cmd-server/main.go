package main

import "fmt"

var VERSION string // Makefile sets this using linker flag, must be uninitialized

func main() {
	fmt.Println("Hello world")
}
