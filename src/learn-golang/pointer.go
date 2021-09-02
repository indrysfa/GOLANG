package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// pass by value
	// data value di duplikat dikirim ke tempat lain
	address1 := Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"}
	address2 := address1

	address2.City = "Bandung" // province country tidak berubah

	fmt.Println(address1)
	fmt.Println(address2)
	// 	{Jakarta Selatan DKI Jakarta Indonesia}
	// {Bandung DKI Jakarta Indonesia}

	// pass by reference
	// data tetap ditempat / dilokasi yg sama
	// gunakan operator & => tambahkan karakter '&' sebagai pointer dari address3
	address3 := Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"}
	address4 := &address3 // address4 adalah pointer

	address4.City = "Bandung"

	fmt.Println(address3)
	fmt.Println(address4)
	// 	{Bandung DKI Jakarta Indonesia}
	// &{Bandung DKI Jakarta Indonesia}

	// menggunakan operator (*)
	var address5 Address = Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"}
	var address6 *Address = &address5 // address6 adalah pointer
	var address7 *Address = &address5 // address7 adalah pointer

	address6.City = "Bandung"

	*address6 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address5)
	fmt.Println(address6)
	fmt.Println(address7)
	// 	{Malang Jawa Timur Indonesia}
	// &{Malang Jawa Timur Indonesia}
	// jika menggunakan * semua value akan berubah
}
