package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func StartServer(spiderVersePath string) {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"ok": false, "description": "form file is missing"})
			return
		}

		path := c.Query("path")

		if path == "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"ok": false, "description": "missing path query parameter"})
			return
		}

		if err := c.SaveUploadedFile(file, filepath.Join(spiderVersePath, path, file.Filename)); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"ok": true, "file": filepath.Join(path, file.Filename)})
	})

	err := router.Run()
	if err != nil {
		os.Exit(1)
	}
}
