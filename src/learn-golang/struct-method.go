package main

import "fmt"

// seakan data struct punya func sayHi
type Customer2 struct {
	Name, Address string
	Age           int
}

func (customer Customer2) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var indry Customer2
	indry.Name = "Indry"
	indry.Address = "Jakarta"
	indry.Age = 27

	indry.sayHi("Joko")
	// Hello Joko My Name is Indry
}
