package main

import (
	"fmt"
	"sync"
)

func paralellMain() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func(task string) {
			// ジョブを実効
			// このサンプルでは出力だけしている
			fmt.Println(task)
			wg.Done()
		}(task)
	}
	wg.Wait()
}
