package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":10001")
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 1024)
	for {
		n, _, err := l.ReadFrom(b)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		fmt.Println("Received ", string(b[0:n]))
	}

	fmt.Println("Shut down server")

}
