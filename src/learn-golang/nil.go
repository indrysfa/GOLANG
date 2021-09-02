package main

import (
	"fmt"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

// nil hanya bisa digunakan pada tipe data seperti interface, function, map, slice, pointer & channel
func NewMap() {
	var person map[string]string {
		if name == "" {
			return nil
		} else {
			return map[string]string{
				"name": name,
			}
		}
		
	} 
}

func main()  {
	var person map[string]string

	if person["name"] == "" {
		fmt.Println("Data kosong")
	} else {
		fmt.Println("")
	}
}