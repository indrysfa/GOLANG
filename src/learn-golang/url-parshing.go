package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://localhost:8080/hello?name=Indry Sefviana&age=27"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme) // http
	fmt.Printf("host: %s\n", u.Host)       // localhost
	fmt.Printf("path: %s\n", u.Path)       // hello

	var name = u.Query()["name"][0] // Indry Sefviana
	var age = u.Query()["age"][0]   // 27
	fmt.Printf("name: %s, age: %s\n", name, age)
}
