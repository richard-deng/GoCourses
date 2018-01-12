package main

import (
    "fmt"
)

func main() {
	var str string
	str = "hello world"
	fmt.Println(str)
	fmt.Println("字符串长度:", len(str))
	fmt.Printf("从字符串中取字符:%c\n", str[1])
}
