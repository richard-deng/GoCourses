package main

import (
	"fmt"
	"mystruct"
)

func main() {
	fmt.Println("vim-go")
	user := new(mystruct.User)
	user.Id = 100
	user.Name = "张三"
	user.SayHi()
	fmt.Println(user)
}
