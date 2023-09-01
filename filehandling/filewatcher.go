package filehandling

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"gorm.io/gorm"

	"github.com/guillaumeLepape/files-across-the-spider-verse/constant"
	"github.com/guillaumeLepape/files-across-the-spider-verse/database"
)

var watcher *fsnotify.Watcher

func StartFileWatcher(dirName string, db *gorm.DB) {
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	if err := filepath.Walk(dirName, WatchDir); err != nil {
		fmt.Println("ERROR", err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				relativePath, _ := filepath.Rel(dirName, event.Name)

				if !(relativePath == constant.SpiderVerseMetadataJournal) && !(relativePath == constant.SpiderVerseMetadata) {
					var hosts []database.Host

					db.Find(&hosts)

					for _, host := range hosts {
						db.Create(&database.FileChange{Path: relativePath, Content: "", Op: fmt.Sprint(event.Op), HostID: host.ID})
					}
				}
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func WatchDir(path string, fi os.FileInfo, _ error) error {
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}

	return nil
}
