package main 

import (
    "fmt"
)

func main() {
	var a int
	var f float32 = 3.14
	a = int(f)
	b := int(f) + 1
	fmt.Println(a)
	fmt.Println(b)
}
