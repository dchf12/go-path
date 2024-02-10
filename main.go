package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func main() {
	var files []string
	dirPath := "."
	err := filepath.WalkDir(dirPath, func(path string, info fs.DirEntry, err error) error {
		dirName := filepath.Base(dirPath)
		if info.IsDir() && info.Name() != dirName {
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
	dir, err := filepath.Abs("operators/arithmetic/")
	if err != nil {
		log.Println(err)
	}

	log.Println(filepath.Join("golang", "files"))
	log.Println(filepath.Join(dir, "/files", "//read"))
}
