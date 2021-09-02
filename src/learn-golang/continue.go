package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 { // jika menemukan genap akan di skip
			continue
			// akan dihentikan perulangan saat ini dan dilanjutkan perulangan selanjutnya
		}
		fmt.Println("Perulangan ke ", i)
	}
}
