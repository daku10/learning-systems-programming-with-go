package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

func randMain() {
	a := make([]byte, 20)
	rand.Read(a)
	fmt.Println(hex.EncodeToString(a))
}

// func main() {
// 	randMain()
// }