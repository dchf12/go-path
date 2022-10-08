package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fileName := "main.go"
	path, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		if t, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("No, " + path + " does not exists")
		} else {
			fmt.Println("Yes, " + path + " exists")
			fmt.Println(t.IsDir())
		}
	}
}
