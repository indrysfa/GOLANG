package main

import "fmt"

func getFullName() (string, string) {
	return "Indry", "Sefviana"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	// value jika tidak digunakan
	firstName, _ = getFullName()
	fmt.Println(firstName)
}
