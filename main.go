package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fileName := "default.md"
	fmt.Println(fileName)
	dir, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dir)
	}
}
