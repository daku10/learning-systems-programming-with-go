package main

import (
	"fmt"
	"time"
)

func tickMain() {
	fmt.Println("waiting 5 secons")
	for now := range time.Tick(5 * time.Second) {
		fmt.Println("now: ", now)
	}
}

func main() {
	tickMain()
}
