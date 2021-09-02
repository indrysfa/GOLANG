package main

import "fmt"

// function sebagai value
func getGoodbye(name string) string {
	return "Hello " + name
}
func main() {
	goodbye := getGoodbye
	fmt.Println(goodbye("Indry"))
}
