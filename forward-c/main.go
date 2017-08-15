package main

import (
	"log"
	"net"
)

func main() {
	c, e := net.Dial("tcp", "127.0.0.1:1102")
	if e != nil {
		log.Println(e)
	}
	defer c.Close()
}
