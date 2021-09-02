package main

import "fmt"

func main() {
	// deklarasi array
	var names [3]string

	names[0] = "Indry"
	names[1] = "Sefviana"

	fmt.Println(names[0])
	fmt.Println(names[1])

	// deklarasi array lebih simple
	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values) // hasil: [90 95 80]

	// panjang data
	fmt.Println(len(names))
	fmt.Println(len(values))
}
