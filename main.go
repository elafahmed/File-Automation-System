package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	srcDir := "E:/semester 2/m213"
	destDir := "E:/book/oop"

	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		os.MkdirAll(destDir, os.ModePerm)
	}

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), "oop.pdf") {
			oldPath := path
			newPath := filepath.Join(destDir, info.Name())
			err := os.Rename(oldPath, newPath)
			if err != nil {
				return err
			}
			fmt.Printf("Moved: %s -> %s\n", oldPath, newPath)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF files ending with 'oop' moved successfully!")
}
