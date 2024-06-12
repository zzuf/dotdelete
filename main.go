package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dotdelete.exe [directory]")
		return
	}
	rootDir := os.Args[1]
	var deletedFiles int
	if err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		return visit(path, info, err, &deletedFiles)
	}); err != nil {
		fmt.Printf("Error walking the path %q: %v\n", rootDir, err)
	}
	fmt.Printf("Total files deleted: %d\n", deletedFiles)
}

func visit(path string, info os.FileInfo, err error, deletedFiles *int) error {
	if err != nil {
		fmt.Printf("Error accessing the path %q: %v\n", path, err)
		return err
	}

	if info.IsDir() {
		return nil
	}

	// Check for specific file names and delete them
	if shouldDelete(info.Name()) {
		fmt.Printf("Deleting file: %s\n", path)
		if err := os.Remove(path); err != nil {
			fmt.Printf("Error deleting file %q: %v\n", path, err)
		} else {
			*deletedFiles++
		}
	}

	return nil
}

func shouldDelete(filename string) bool {
	switch filename {
	case ".DS_Store", ".fseventsd", ".Spotlight-V100", ".Trashes":
		return true
	}
	if strings.HasPrefix(filename, "._") {
		return true
	}
	return false
}
