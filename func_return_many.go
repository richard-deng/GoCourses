package main

import (
	"fmt"
)

func f2(a, b int) (int, float32) {
	return a * b, float32(a) / float32(b)
}

func main() {
	sum, div := f2(3, 6)
	fmt.Println(sum, div)
}
