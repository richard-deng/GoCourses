package main

import (
	"os"
	"fmt"
	"net"
	"io/ioutil"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkErr(err)

    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkErr(err)

    result, err := ioutil.ReadAll(conn)
    checkErr(err)

    fmt.Println(string(result))
    os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}