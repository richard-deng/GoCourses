package main

import (
    "fmt"
)

func main() {
	var f1 float64 = 1.23
	f2 := 12.3
	f3 := 12356.789e5
	f4 := f3 + 20
	fmt.Println(f2 - f1)
	fmt.Println(f3)
	fmt.Println(f4)
}
