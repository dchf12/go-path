package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dir)
	}
}
