package main

import "fmt"

func main() {
	// type datanya diberi alias
	type NoKTP string
	type Status bool

	var noKtpIndry NoKTP = "217367147481813728"
	var StatusIndry Status = true
	fmt.Println(noKtpIndry)
	fmt.Println(StatusIndry)
}
