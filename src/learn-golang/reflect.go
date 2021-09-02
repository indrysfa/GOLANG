package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` // dengan tag
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Indry"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	// Output
	// 1
	// mengambil field name dalam struct
	fmt.Println(sampleType.Field(0).Name)
	// Output
	// Name
	// mengambil tag pada struct
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	// Output
	// true
	// 10

	fmt.Println(IsValid(sample))
	// true
	sample.Name = "" // jika value kosong
	fmt.Println(IsValid(sample))
	// false
}
