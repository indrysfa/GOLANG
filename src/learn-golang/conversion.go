package main

import "fmt"

func main() {
	// cakupan bilangan
	// int64 = -9223372036854775808 ↔ 9223372036854775807
	// int32 = -2147483648 ↔ 2147483647
	// int16 = -32768 ↔ 32767
	// int8 = -128 ↔ 127

	// jika nilai melebihi cakupan bilangan maka akan rolling ke nilai cakupan paling rendah
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Indry"
	var i byte = name[0] // deklarasi dimana letak karakter i
	var iString string = string(i)

	fmt.Println(name)
	fmt.Println(iString)
}
