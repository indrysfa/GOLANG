package main

import (
	"fmt"
	"learn-golang/helper"
)

func main() {
	helper.SayHello("Indry")
	fmt.Println(helper.Application)
	// Output
	// Hello Indry
	// Belajar Golang
	helper.sayGoodbye("Indry")  // tidak bisa diakses
	fmt.Println(helper.version) // tidak bisa diakses
	// Output
	// .\import.go:10:2: cannot refer to unexported name helper.sayGoodbye
	// .\import.go:11:14: cannot refer to unexported name helper.version
}
