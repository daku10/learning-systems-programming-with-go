package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// パスをそのままクリーンにする
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))
	// path/path.go
	// パスを絶対パスに整形
	abspath, err := filepath.Abs("path/filepath/path_unix.go")
	// なぜここ書いてないの？
	if err != nil {
		panic(err)
	}
	fmt.Println(abspath)

	// パスを相対パスに整形
	relpath, err := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/path/filepath/path.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(relpath)
}