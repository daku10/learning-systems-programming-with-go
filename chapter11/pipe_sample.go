package main

import (
	"fmt"
	"time"
)

func pipeSample() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)		
	}
}

// func main() {
// 	pipeSample()
// }