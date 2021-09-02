package main

import "fmt"

func main() {
	// slice = potongan data yang ada di array
	// pointer = peunjuk data pertama yang ada di slice
	// length = panjang dari slice
	// capacity = kapasitas yang boleh digunakan / ditampung

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1)) // 8 : kapasitas yang ada di slice

	var slice2 = months[10:]
	fmt.Println(slice2) // [November Desember]

	var slice3 = append(slice2, "Indry") // menambahkan data diakhir dalam array
	fmt.Println(slice3)                  // [November Desember Indry]

	// Hati-hati menggunakan append, cek capacity nya terlebih dahulu, jika masih ada space maka akan membuat array baru, jika tidak maka akan menggantikan value lama atau tidak tercetak jika diakhir data
	// ===============================================

	// menggunakan make
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Indry"
	newSlice[1] = "Sefviana"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// menggunakan copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice) // slice yang sering digunakan dibandingkan array
}
