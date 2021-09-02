package main

import "fmt"

type Address1 struct {
	City, Province, Country string
}

// merubah country menggunakan pointer
// *Address1 = pointer
func ChangeCountryToIndonesia(address *Address1) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address1{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "",
	}
	// gunakan '&' untuk mendefinisikan bahwa itu parameter dari pointer
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
	// Output
	// {Bandung Jawa Barat Indonesia}

	// cara kedua
	var alamatPointer *Address1 = &alamat
	ChangeCountryToIndonesia(alamatPointer) // tidak perlu menambahkan '&' lagi
	fmt.Println(alamat)
	// Output
	// {Bandung Jawa Barat Indonesia}
}
