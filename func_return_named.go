package main

import (
	"fmt"
)

func f3(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	sum, sub := f3(3, 6)
	fmt.Println(sum, sub)
}
