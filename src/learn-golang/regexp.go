package main

import (
	"fmt"
	"regexp"
)

func main() {
	// regular expression
	var regex *regexp.Regexp = regexp.MustCompile("i([a-z]y)")

	fmt.Println(regex.MatchString("ily"))
	fmt.Println(regex.MatchString("iry"))
	fmt.Println(regex.MatchString("dRy"))
	// Output
	// true
	// true
	// false

	search := regex.FindAllString("ily iga iry ici ihi iru", -1)
	fmt.Println(search)
	// Output
	// [ily iry]
}
