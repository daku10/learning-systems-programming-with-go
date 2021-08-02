package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
)

func DatagramClient() {
	clientPath := filepath.Join(os.TempDir(), "unixdomainsocket-client")
	// エラーチェックは不要なので削除 (存在しなかったらしなかったで問題ない)
	os.Remove(clientPath)
	conn, err := net.ListenPacket("unixgram", clientPath)
	if err != nil {
		panic(err)
	}
	// 送信先のアドレス
	unixServerAddr, err := net.ResolveUnixAddr("unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
	// why var?
	var serverAddr = unixServerAddr
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("Sending to server")
	_, err = conn.WriteTo([]byte("Hello from Client"), serverAddr)
	if err != nil {
		panic(err)
	}
	log.Println("Receving from server")
	buffer := make([]byte, 1500)
	length, _, err := conn.ReadFrom(buffer)
	if err != nil {
		panic(err)
	}
	log.Printf("Received: %s\n", string(buffer[:length]))
}
