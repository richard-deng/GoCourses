package main

import (
    "fmt"
)

type coordinate struct {
	x int
	y int
}

func (recv *coordinate) swap() {
	// var temp int
	// temp, recv.x, recv.y = temp, recv.y, recv.x
	recv.x, recv.y = recv.y, recv.x
	fmt.Println(recv)
}

func main() {
	r1 := coordinate{3, 4}
	fmt.Println(r1)
	p := &r1
	p.swap()
	fmt.Println(r1)
}
