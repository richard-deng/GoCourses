package main

import (
	"os"
	"fmt"
	"golang.org/x/net/websocket"
)

type Person struct {
	Name    string
	Emails  []string
}

func main(){
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "ws://host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := websocket.Dial(service, "", "http://localhost")
	checkErr(err)

	person := Person{
		Name: "Jan",
		Emails:[]string{"xx@qq.com", "yy@163.com"},
		}

	err = websocket.JSON.Send(conn, person)
	if err != nil {
		fmt.Println("Couldn't send msg "+ err.Error())
	}
	os.Exit(0)
}


func checkErr(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
