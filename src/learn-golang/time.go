package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	// Output
	// 2021
	// September
	// 1
	// 11
	// 18
	// 19
	// 688674300

	utc := time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)
	// Output
	// 2021-10-10 10:10:10.00000001 +0000 UTC

	// parsing
	layout := "2006-01-02" // format date pada golang
	parse, _ := time.Parse(layout, "2021-09-01")
	fmt.Println(parse)
	// Output
	// 2021-09-01 00:00:00 +0000 UTC
}
