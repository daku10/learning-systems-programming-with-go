package main

import (
	"fmt"
	"time"
)

func nowMain() {
	t := time.Now()
	fmt.Println(t.String())
}

func main() {
	nowMain()
}