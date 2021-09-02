package main

import "fmt"

func main() {
	var name1 = "Indry"
	var name2 = "Suho"

	var result bool = name1 == name2 // false
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)  // false
	fmt.Println(value1 < value2)  // true
	fmt.Println(value1 == value2) // false
	fmt.Println(value1 != value2) // true
}
