package main

import (
	"fmt"
)

type student struct {
	id    int
	name  string
	sex   bool
	age   int
	class string
}

func main() {
	// 全部初始化
	stu1 := student{13001, "李明", false, 19, "网络01"}
	stu2 := &student{13002, "张晓", false, 19, "网络01"}

	// 部分初始化
	stu3 := student{name: "王乐", age: 19}
	stu4 := &student{name: "赵琼", id: 13003}

	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu3)
	fmt.Println(stu4)
}
