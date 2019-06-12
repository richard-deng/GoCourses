package main

import (
	"fmt"
)

func f(a int) {
	fmt.Println("函数f开始运行...")
	if a > 100 {
		panic("参数值超出范围")
	} else {
		fmt.Println("函数f调用结束")
	}
}

func main() {
	// 必须要先声明defer,否则捕获不到panic异常
	defer func() {
		fmt.Println("函数defer开始运行...")
		if err := recover(); err != nil {
			// 这里的err就是panic传入的内容
			fmt.Println("程序异常退出: ", err)
		} else {
			fmt.Println("程序正常退出.")
		}
	}()
	f(10)
}
