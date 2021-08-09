package main

import (
	"fmt"
	"time"
)

func forMain() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		// こんな感じでローカル変数に落とせば、後述の問題は回避できる
		internal := task
		go func() {
			// goroutineが起動するときにはループが回りきって
			// 全部のタスクが最後のタスクになってしまう
			// 警告出るの良いね			
			fmt.Println(internal)
		}()
	}
	time.Sleep(time.Second)
}

func main() {
	forMain()
}