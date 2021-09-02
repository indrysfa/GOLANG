package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// disarankan menggunakan perulangan do
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
	// Output
	// Data 0
	// Data 1
	// Data 2
	// Data 3
	// Data 4
}
