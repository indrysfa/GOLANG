package main

// jika tidak dipakai gunakan _ sebagai blank identifier
import (
	"fmt"
	"learn-golang/database"
	_ "learn-golang/helper"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
