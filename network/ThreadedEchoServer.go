package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:]))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
