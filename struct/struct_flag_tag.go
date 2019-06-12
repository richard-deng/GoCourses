package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Id   int    "账号"
	Name string "姓名"
	Sex  bool   "性别"
}

func main() {
	fmt.Println("vim-go")
	u := user{100, "张三", false}
	// 使用TypeOf()函数获取对象的类型
	t := reflect.TypeOf(u)
	// 使用ValueOf()函数获取对象的值
	v := reflect.ValueOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("% s( % s = % v)\n", f.Tag, f.Name, v.Field(i).Interface())
	}
}
