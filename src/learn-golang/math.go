package main

import (
	"fmt"
	"math"
)

func main() {
	// pembulatan keatas atau kebawah sesuai dengan yang paling dekat
	fmt.Println(math.Round(1.7))
	// pembulatan kebawah
	fmt.Println(math.Floor(1.4))
	// pembulatan keatas
	fmt.Println(math.Ceil(1.5))
	// mengambil nilai terbesar
	fmt.Println(math.Max(10, 20))
	// mengambil nilai terkecil
	fmt.Println(math.Min(10, 20))
}
