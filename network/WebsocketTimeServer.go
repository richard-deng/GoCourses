package main

import (
	"golang.org/x/net/websocket"
	"time"
	"fmt"
	"net/http"
	"os"
)

var ROOT_DIR = "/Users/admin/Documents/develop/go/src/network/templates"

func GetTime(ws *websocket.Conn) {
	for {
		msg := time.Now().String()
        fmt.Println("Sending to client: " + msg)
        err := websocket.Message.Send(ws, msg)
        if err != nil {
        	fmt.Println("Can't send")
        	break
		}
		time.Sleep(2 * 1000 * 1000 * 1000)

		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir(ROOT_DIR))
	http.Handle("/GetTime", websocket.Handler(GetTime))
	http.Handle("/", fileServer)
	err := http.ListenAndServe(":12345", nil)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}