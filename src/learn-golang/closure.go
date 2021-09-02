package main

import "fmt"

func main() {
	// sebuah blockscopebisa berinteraksi dengan scope lainnya
	// berubah data diatasnya
	name := "Indry"
	counter := 0

	increment := func() {
		name = "Sefviana"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(name)
	// Output
	// Sefviana
	// Increment
	// Sefviana
	// Sefviana
}
