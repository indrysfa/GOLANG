package main

import (
	"fmt"
)

// type declaration, agar tidak kepanjangan di parameter, buat alias
type Filter func(string) string

func sayHelloWithFilterv1(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// function dalam parameter
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilterv1("Indry", spamFilter)
	sayHelloWithFilterv1("Anjing", spamFilter)
}
