package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

func seedMain() {
	// 乱数の種を設定
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		// 浮動小数点数 (float64) の乱数を生成
		fmt.Println(rand.Float64())
	}

	// 種からソースを作成
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	src := rand.NewSource(seed.Int64())
	// ソースから乱数生成器を作成
	rng := rand.New(src)
	
	for i := 0; i < 10; i++ {
		// 作成した乱数生成器から乱数を出力
		fmt.Println(rng.Float64())
	}
}

// func main() {
// 	seedMain()
// }
