package main

import "fmt"

func main() {
	// Kondisi Perbandingan =
	var name = "Indry"

	switch name {
	case "Indry":
		fmt.Println("Hello Indry")
	case "Sefviana":
		fmt.Println("Hello Sefviana")
	default:
		fmt.Println("Hello Semua")
	}
	//Hello Indry

	switch length := len(name); length > 4 {
	case true:
		fmt.Println("Hello Indry")
	case false:
		fmt.Println("Hello Semua")
	}
}
