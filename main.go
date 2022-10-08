package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	fileName := "main.go"
	dir, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fe := path.Ext(dir)
		fmt.Println(fe)
		fmt.Println(strings.TrimSuffix(dir, fe))
		fmt.Println(strings.TrimSuffix(fileName, fe))
	}
}
