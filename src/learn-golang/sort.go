package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

// alias dari slice User
type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Indry", 27},
		{"Jiyeon", 28},
		{"Daniel", 26},
		{"Mark", 24},
		{"Chen", 30},
	}

	// sebelum di sort
	fmt.Println(users)
	// Output
	// [{Indry 27} {Jiyeon 28} {Daniel 26} {Mark 24} {Chen 30}]
	sort.Sort(UserSlice(users))
	// menggunakan sort
	fmt.Println(users)
	// Output
	// [{Mark 24} {Daniel 26} {Indry 27} {Jiyeon 28} {Chen 30}]
}
