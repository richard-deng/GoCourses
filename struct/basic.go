package main

import (
	"fmt"
)

type student struct {
	id    int
	name  string
	sex   bool
	age   int
	calss string
}

func main() {
	var stu1 student
	stu2 := student{}
	stu1.name = "李明"
	stu2.name = "张恒"
	stu1.age = 18
	stu2.age = 19
	fmt.Println(stu1)
	fmt.Println(stu2)
}
