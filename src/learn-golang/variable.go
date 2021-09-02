package main

import "fmt"

func main() {
	// deklarasi jika menggunakan var diawal
	var name string

	name = "Indry Sefviana" // auto type data string
	fmt.Println(name)

	name = "Indry"
	fmt.Println(name)

	// deklarasi jika tidak menggunakan var diawal
	country := "Indonesia"
	fmt.Println(country)

	country = "Korea Selatan"
	fmt.Println(country)

	// deklarasi jika var lebih dari satu
	var (
		firstName = "Indry"
		lastName  = "Sefviana"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
