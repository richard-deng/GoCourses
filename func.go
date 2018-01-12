package main 

import (
    "fmt"
    "errors"
)

func f1() {
	fmt.Println("hello world")
}

func f2(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func f3(a, b int) int {
	return a + b
}

func f4(a, b int) (ret float32, err error) {
	if b == 0 {
		err = errors.New("overflow")
		return
	} else {
		return float32(a) / float32(b), nil
	}
}

func main(){
	f1()
	f2(1, 2, "hello")
	sum := f3(10, 10)
	fmt.Println(sum)
	a, _ := f4(250, 10)
	fmt.Println(a)
}
