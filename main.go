package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	var files []string
	dirPath := "."
	err := filepath.WalkDir(dirPath, func(path string, info fs.DirEntry, err error) error {
		dirName := filepath.Base(dirPath)
		if info.IsDir() == true && info.Name() != dirName {
			return filepath.SkipDir
		} else {
			files = append(files, path)
			return nil
		}
	})

	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
