package main

import "fmt"

func main() {
	// constant variable yg tidak bisa diubah lagi
	const firstName string = "Indry"
	const lastName = "Sefviana"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// deklarasi constant lebih simple lagi
	const (
		firstName1 string = "Indry"
		lastName1         = "Sefviana"
		value1            = 1000
	)

	fmt.Println(firstName1)
	fmt.Println(lastName1)
	fmt.Println(value1)
}
