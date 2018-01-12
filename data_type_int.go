package main

import (
    "fmt"
	"math"
)

func main() {
	var sum int
	var a int = 100
	b := 28
	var c int8 = 127
	sum = a + b
	fmt.Println(sum)
	fmt.Println(c + 1)
	fmt.Printf("Min(int8) = %v,  Max(int8) = %v\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("Min(int16) = %v, Max(int16) = %v\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("Min(int32) = %v, Max(int32) = %v\n", math.MinInt32, math.MaxInt32)
}
