package main

import "fmt"

type Blacklist func(string) bool

// function tanpa nama
func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// function ini dibuat didalam variable
	// untuk admin di blocked
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("indry", blacklist)

	// Output
	// You are blocked admin
	// Welcome indry
}
