package main

import (
	"fmt"
)

func f1(args ...int) {
	fmt.Println(args)
}

func f2(args []int) {
	fmt.Println(args)
}

func main() {
	f1(1, 2, 3)
	f1(1, 2, 3, 4, 5)
	a := []int{1, 2, 3}
	b := []int{1, 2, 3, 4, 5}
	f2(a)
	f2(b)
}
