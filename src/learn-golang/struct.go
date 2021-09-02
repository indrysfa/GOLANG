package main

import "fmt"

// anggap aja ini class
// kumpulan dari field
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var indry Customer
	indry.Name = "Indry"
	indry.Address = "Jakarta"
	indry.Age = 27

	fmt.Println(indry.Name)
	fmt.Println(indry.Address)
	fmt.Println(indry.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Bandung",
		Age:     32,
	}
	fmt.Println(joko)
	// {Joko Bandung 32}
}
