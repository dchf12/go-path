package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fileName := "drafts/default.md"
	//fileName := "drafts/"
	fmt.Println(fileName)
	path, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(filepath.Dir(path))
	}
}
