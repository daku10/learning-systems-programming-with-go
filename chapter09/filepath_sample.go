package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, name := filepath.Split(os.Getenv("PYENV_ROOT"))
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)
}