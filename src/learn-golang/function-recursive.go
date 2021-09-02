package main

import "fmt"

// function yang memanggil function dirinya sendiri
// menggunakan looping
func factorialLoop(value int) int {
	result := 1
	for i := value; i < 0; i-- {
		result *= i
	}
	return result
}

// menggunakan recursive
func factorialRecursive(value int) int {
	if value == 1 {
		// kondiri berhenti looping
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	// menggunakan looping
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)
	// menggunakan recursive
	recursive := factorialRecursive(10)
	fmt.Println(recursive)

	// Output
	// 1
	// 120
	// 3628800
}
