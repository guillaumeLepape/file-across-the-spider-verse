package filehandling

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

type FileMetadata struct {
	Path      string    `json:"path"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FindFilesAfter(dirName string, t time.Time) (fileMetadata []FileMetadata, err error) {
	err = filepath.Walk(dirName, func(absolutePath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		relativePath, _ := filepath.Rel(dirName, absolutePath)

		if !info.IsDir() && info.ModTime().After(t) {
			fileMetadata = append(
				fileMetadata,
				FileMetadata{Path: relativePath, UpdatedAt: info.ModTime()},
			)
		}

		return nil
	})

	return
}

func ScanFolder(dirName string) []FileMetadata {
	var fileMetadata []FileMetadata

	err := filepath.Walk(dirName, func(absolutePath string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", absolutePath, err)
			return err
		}
		if info.Mode().IsRegular() {
			relativePath, _ := filepath.Rel(dirName, absolutePath)
			fileMetadata = append(fileMetadata, FileMetadata{Path: relativePath, UpdatedAt: info.ModTime()})
			return nil
		}

		return nil
	})
	if err != nil {
		log.Fatal("error walking the path")
	}

	return fileMetadata
}
