package main

import (
	"fmt"
	"os"
)

func wdMain() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
}

// func main(){
// 	wdMain()
// }
