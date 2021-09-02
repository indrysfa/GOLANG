package main

import "fmt"

// digunakan untuk sistem error dan aplikasi berhenti
func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
	// Output true
	// Aplikasi selesai
	// panic: APLIKASI ERROR

	// Output false
	// Aplikasi berjalan
	// Aplikasi selesai
}
