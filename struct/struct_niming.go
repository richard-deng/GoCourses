package main

import (
	"fmt"
)

type people struct {
	name string
	sex  bool
}

type teacher struct {
	people     //匿名字段
	department string
}

func main() {
	teacher1 := teacher{people{"郑智", false}, "Computer Science"}
	fmt.Println(teacher1)
	fmt.Println(teacher1.name)
	fmt.Println(teacher1.sex)
	fmt.Println(teacher1.department)
}
