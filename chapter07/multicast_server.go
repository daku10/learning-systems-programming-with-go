package main

import (
	"fmt"
	"net"
	"time"
)

const interval = 5 * time.Second

func main() {
	fmt.Println("Start tick server at 224.0.0.1:9999")
	conn, err := net.Dial("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	start := time.Now()
	wait := start.Truncate(interval).Add(interval).Sub(start)
	time.Sleep(wait)
	ticket := time.Tick(interval)
	for now := range ticket {
		conn.Write([]byte(now.String()))
		fmt.Println("Tick: ", now.String())
	}
}
