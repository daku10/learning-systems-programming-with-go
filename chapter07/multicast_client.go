package main

import (
	"fmt"
	"net"
)

func main() {
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	fmt.Printf("Listen tick server at 224.0.0.1:9999 I am: %v", address)
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, address)
	defer listener.Close()

	buffer := make([]byte, 1500)

	for {
		length, remoteAddress, err := listener.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now %s\n", string(buffer[:length]))
	}
}
