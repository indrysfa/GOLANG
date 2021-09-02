package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	// errors => fitur error bawaan golang
	// New => struct yang sudah disediakan golang
	var contohError error = errors.New("Ups Error nih")
	fmt.Println(contohError.Error())

	// contoh kedua
	hasil, err := Pembagi(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
