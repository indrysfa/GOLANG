package main

import "fmt"

// variabelnya tidak perlu dipanggil pada return
func getFullName2() (firstName string, lastName string) {
	firstName = "Indry"
	lastName = "Sefviana"

	return
}

func main() {
	a, b := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
}
