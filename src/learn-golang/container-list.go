package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Indry")
	data.PushBack("Sefviana")
	data.PushBack("Developer")

	// mengambil data paling depan
	fmt.Println(data.Front().Value)
	// Output
	// Indry

	// mengambil data paling belakang
	fmt.Println(data.Back().Value)
	// Output
	// Developer

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	// Output
	// Indry
	// Sefviana
	// Developer

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
	// Output
	// Developer
	// Sefviana
	// Indry
}
