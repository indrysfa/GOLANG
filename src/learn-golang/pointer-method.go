package main

import "fmt"

type Man struct {
	Name string
}

// usahakan menggunakan pointer '*' struct func
func (man *Man) Merried() {
	man.Name = "Mrs. " + man.Name
}

func main() {
	// jika tidak menggunakan '*' data indry di duplikat lalu dikirim ke func Married
	// hasil run => Indry
	indry := Man{"Indry"}
	indry.Merried()

	fmt.Println(indry.Name)
	// hasil run menggunakan pointer '*' => Mrs. Indry
}
