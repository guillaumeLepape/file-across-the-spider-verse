package filehandling

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"time"
)

type FileMetadata struct {
	Name    string
	ModTime time.Time
}

func ScanFolder(dirName string) []FileMetadata {
	var fileMetadata []FileMetadata

	err := filepath.Walk(dirName, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.Mode().IsRegular() {
			fileMetadata = append(fileMetadata, FileMetadata{Name: path, ModTime: info.ModTime()})
			return nil
		}
		return nil
	})

	if err != nil {
		log.Fatal("error walking the path")
	}

	return fileMetadata
}
