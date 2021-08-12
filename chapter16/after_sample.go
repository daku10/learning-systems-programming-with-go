package main

import (
	"fmt"
	"time"
)

func afterMain() {
	fmt.Println("waiting 5 seconds")
	after := time.After(5 * time.Second)
	time.Sleep(10 * time.Second)
	fmt.Println("10 seconds done")
	<-after
	fmt.Println("done")
}

func main() {
	afterMain()
}
