package main

import (
	"fmt"
	"os"
)

func execMain() {
	path, _ := os.Executable()
	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
	fmt.Printf("実行ファイルパス: %s\n", path)
}

// func main() {
// 	execMain()
// }