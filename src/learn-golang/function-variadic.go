package main

import "fmt"

// variable argument sebagai value
// value numbers disini tidak dibatasi jika menggunakan ...
func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}
func main() {
	total := sumAll(10, 90, 50, 30, 20)
	fmt.Println(total)

	// jika menggunakan slice
	slice := []int{10, 90, 50, 30, 20}
	total = sumAll(slice...)
	fmt.Println(total)
}
