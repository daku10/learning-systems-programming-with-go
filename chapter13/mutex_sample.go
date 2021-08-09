package main

import (
	"fmt"
	"sync"
	"time"
)


var id int

func generateId(mutex *sync.Mutex) int {
	// Lock()/Unlock()をペアで呼び出してロックする
	mutex.Lock()
	defer mutex.Unlock()
	id++
	return id
}

func mutexMain() {
	// sync.Mutex構造体の変数宣言
	// 次の宣言をしてもポインタ型になるだけで正常に動作する
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
		}()
	}
}

func main() {
	mutexMain()
	// これ書かないと途中で終了しちゃうね
	time.Sleep(time.Second)
}