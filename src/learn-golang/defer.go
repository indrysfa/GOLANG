package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(10)
	// jika diisi 0 maka print result akan dihentikan
	// Output
	// Run Application
	// Selesai memanggil function
	// panic: runtime error: integer divide by zero

	// jika diisi valuenya contoh 10
	// Output
	// Run Application
	// Result 1
	// Selesai memanggil function
}
