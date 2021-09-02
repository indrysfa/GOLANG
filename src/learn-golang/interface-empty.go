package main

import "fmt"

/* interface kosong adalah interface yang tidak memiliki deklarasi method satupun, dan membuat otomatis semua tipe data akan menjadi implementasinya */
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Upsss"
	}
}

func main() {
	var data interface{} = Ups(3)
	fmt.Println(data)
	// Output
	// Upsss
}
