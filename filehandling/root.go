package filehandling

import (
	"fmt"
	"log"

	"os"
	"path/filepath"
	"strings"

	"github.com/guillaumeLepape/files-across-the-spider-verse/constant"
)

func GetHomeDir() string {
	homedir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return homedir
}

func CreateSpiderFilePath() string {
	return fmt.Sprint(filepath.Join(GetHomeDir(), constant.SpiderFile))
}

func CreateSpiderVerseDir(spiderFilePath string) {
	mkdirErr := os.MkdirAll(spiderFilePath, os.ModePerm)
	if mkdirErr != nil {
		log.Fatal(mkdirErr)
	}
}

func GetSpiderVersePath() string {
	filename := CreateSpiderFilePath()

	fi, err := os.Lstat(filename)

	if err != nil {
		return CreateSpiderFile()
	}

	if !fi.Mode().IsRegular() {
		log.Fatalln(
			filename,
			"already exists but it is not a regular file. Please remove it and rerun the program.")
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	spiderVersePath := strings.Replace(string(data), "\n", "", -1)

	CreateSpiderVerseDir(spiderVersePath)

	return spiderVersePath
}

func CreateSpiderFile() string {
	homeDir := GetHomeDir()

	fmt.Print("Location of spider version (relative to home directory", homeDir, "): ")
	var fileDir string
	fmt.Scanln(&fileDir)

	spiderFilePath := filepath.Join(homeDir, fileDir)

	err := os.WriteFile(
		CreateSpiderFilePath(),
		[]byte(fmt.Sprintln(spiderFilePath)),
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	CreateSpiderVerseDir(spiderFilePath)

	return spiderFilePath
}
