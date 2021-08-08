package main

import (
	"fmt"
	"os"
)

func pidMain() {
	fmt.Printf("プロセスID: %d\n", os.Getpid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())
}

// func main() {
// 	pidMain()
// }