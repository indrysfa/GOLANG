package main

import "fmt"

func main() {
	// Kondisi Perulangan / Loop
	counter := 9

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// Menggunakan data slice
	slice := []string{"Indry", "Sefviana"}

	// Tidak menggunakan perulangan
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	// Menggunakan Perulangan
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	// for range => index
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// gunakan _ untuk index kosong
	for _, value := range slice {
		fmt.Println("Index", value)
	}

	// menggunakan map => key
	person := make(map[string]string)
	person["name"] = "Indry"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
