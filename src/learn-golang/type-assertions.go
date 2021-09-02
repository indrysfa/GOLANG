package main

import "fmt"

func random() interface{} {
	return "Indry"
}

func main() {
	// contoh pertama dipaksa / di set menggunakan type string
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)
	// Output
	// Indry

	// contoh menggunakan switch case
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
	// Output
	// Value Indry is string
}
