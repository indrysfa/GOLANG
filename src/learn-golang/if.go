package main

import "fmt"

func main() {
	// Kondisi Perbandingan
	var name = "Indry"

	if name == "Indry" {
		fmt.Println("Hello Indry")
		// dieksekusi jika nilainya true
		// Hello Indry
	} else if name == "Sefviana" {
		fmt.Println("Tidak menyapa")
		// dieksekusi jika nilainya true
		// Tidak menyapa
	} else {
		fmt.Println("Silahkan coba kembali")
		// dieksekusi jika nilainya false
		// Silahkan coba kembali
	}

	if length := len(name); length > 4 {
		fmt.Println("Selamat datang")
	} else {
		fmt.Println("Silahkan coba kembali")
	}
	//Hello Indry
	// Selamat datang
}
