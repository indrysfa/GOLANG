package main

import "fmt"

// tidak bisa membuat data dari interface
// hanya berisi definisi method
// seperti __construct pada oop
type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// definisikan struct untuk membuat data
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var indry Person
	indry.Name = "Indry"

	SayHello(indry)
	// Output
	// Hello Indry

	cat := Animal{
		Name: "Push",
	}
	SayHello(cat)
	// Output
	// Hello Push
}
