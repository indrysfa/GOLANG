package main

import "fmt"

/** recover disimpan di defer function **/
func endApp1() {
	// recover untuk menampung jika terjadi error/panic dan menjalankan sistem selanjutnya
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp1(error bool) {
	defer endApp1()
	if error {
		panic("APLIKASI ERROR") // berhenti disini jika tanpa recover
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp1(true)
	fmt.Println("Sistem selanjutnya")
	// Output true
	// Error dengan message: APLIKASI ERROR
	// Aplikasi selesai

	// Output false
	// Aplikasi berjalan
	// Aplikasi selesai
}
