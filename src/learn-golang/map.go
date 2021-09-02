package main

import (
	"fmt"
)

func main() {
	// map: kumpulan key value
	person := map[string]string{
		"name":    "Indry",
		"address": "Jakarta",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	// map[address:Jakarta name:Indry title:Programmer]
	// key nya memang tidak sesuai dengan yang dibuat tapi tidak masalah
	fmt.Println(person["name"])    // Indry
	fmt.Println(person["address"]) // Jakarta

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Indry"
	book["status"] = "Aktif"
	fmt.Println(book)
	fmt.Println(len(book)) // jumlah data variabel book
	// map[author:Indry status:Aktif title:Belajar Go-Lang]
	// 3

	delete(book, "status") // menghapus key status yg ada di variabel book

	fmt.Println(book) // outputnya setelah dihapus
	// map[author:Indry title:Belajar Go-Lang]
	// 2
	fmt.Println(len(book)) // jumlah data variabel book
}
