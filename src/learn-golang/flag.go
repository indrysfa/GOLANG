package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your hostname")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println("Host : ", host)
	fmt.Println("User : ", user)
	fmt.Println("Password : ", password)
	// Output
	// Host :  0xc000088260
	// User :  0xc000088270
	// Password :  0xc000088280

	// jika ingin mencetak value
	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	// Output
	// Host :  localhost
	// User :  root
	// Password :  root
}
