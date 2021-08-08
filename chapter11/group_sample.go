package main

import (
	"fmt"
	"os"
	"syscall"
)

func groupMain() {
	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Fprintf(os.Stderr, "グループID: %d セッションID: %d\n", syscall.Getpgrp(), sid)
}

// func main() {
// 	groupMain()
// }