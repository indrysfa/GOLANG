package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	// untuk mengetahui hostname yang sedang digunakan
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
	// Output
	// Hostname :  INDRY
	// indry
	// indry

	// gunakan Terminal untuk set username & password
	// export APP_USERNAME=root
	// export APP_USERNAME=root
}
