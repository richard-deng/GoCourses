package main

import (
    "fmt"
)

func main() {
	var cp1, cp2 complex64
	cp1 = 1.2 + 3.4i
	cp2 = cp1
	cp3 := complex(1.2, 3.4)
	fmt.Println(cp1)
	fmt.Println(cp2)
	fmt.Println(cp3)
	fmt.Println("real: ", real(cp1))
	fmt.Println("imag: ", imag(cp1))
}
