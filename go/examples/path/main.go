package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	fmt.Printf("path - Base (/foo/bar/readme.txt): %s\n", path.Base("/foo/bar/readme.txt"))
	fmt.Printf("filepath - Base (/foo/bar/readme.txt): %s\n", filepath.Base("/foo/bar/readme.txt"))

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Working Directory: %s\n", wd)

	var counter int
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Walk: %s\n", path)
		counter++
		if counter >= 10 {
			return errors.New("enough")
		}
		return nil
	})
}
