package main

import (
	"fmt"
)

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func main() {
	var result int = 0
	for i := 0; i < 20; i++ {
		result = fibonacci(i)
		if i%5 == 0 {
			fmt.Println("\n")
		}
		fmt.Printf("% 4d", result)
	}
}
