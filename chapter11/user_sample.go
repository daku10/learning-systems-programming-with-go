package main

import (
	"fmt"
	"os"
)

func userMain() {
	fmt.Printf("ユーザID: %d\n", os.Getuid())
	fmt.Printf("グループID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("サブグループID: %v\n", groups)
}

// func main() {
// 	userMain()
// }