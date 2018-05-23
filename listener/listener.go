package main

import (
	"log"
	"net"
)

func main() {
	_, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatal(err)
	}
}
