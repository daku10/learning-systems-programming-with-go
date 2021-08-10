package main

import (
	"fmt"
	"runtime"
	"sync"
)

func poolMain() {
	// Poolを作成。Newで新規作成時のコードを実装
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	// 追加した要素から受け取れる
	// プールが空だと新規作成
	pool.Put("manualy added: 1")
	pool.Put("manualy added: 2")
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get()) // これは新規作成

	fmt.Printf("count = %d\n", count)

	// GCを呼ぶと追加された要素が消える
	pool.Put("removed 1")
	pool.Put("removed 2")
	runtime.GC()
	fmt.Println(pool.Get())
}

func main() {
	poolMain()
}