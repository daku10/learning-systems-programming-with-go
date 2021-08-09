package main

import (
	"fmt"
	"time"
)

func sub1(c int) {
	fmt.Println("share by arguments:", c*c)
}

func sub1Main() {
	// 引数渡し
	go sub1(10)

	// クロージャのキャプチャ渡し
	c := 20
	go func() {
		fmt.Println("share by capture", c * c)
	}()
	// ああ、これが推奨されないやり方ね
	time.Sleep(time.Second)
}

func main() {
	sub1Main()
}