package main

import (
	"fmt"
	"sync"
	"time"
)

func wgMain() {
	var wg sync.WaitGroup

	// ジョブ数をあらかじめ登録
	wg.Add(2)

	go func() {
		// 非同期で仕事をする
		fmt.Println("仕事1")
		// Doneで完了を通知
		wg.Done()
	}()

	go func() {
		// 待つ
		time.Sleep(time.Second * 2)
		//非同期で仕事をする
		fmt.Println("仕事2")
		// Doneで完了を通知
		wg.Done()
	}()

	// すべての処理が終わるのを待つ
	wg.Wait()
	fmt.Println("終了")
}

func main() {
	wgMain()
}