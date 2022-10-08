package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fileName := "default.md"
	dir, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(err)
		fmt.Println(filepath.Base(dir))
	}
}
