package filehandling

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetHomeDir() string {
	homedir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return homedir
}

func SpiderVerseLocation() string {
	return fmt.Sprint(filepath.Join(GetHomeDir(), ".spider_file"))
}

func GetSpiderVerseLocation() (string, error) {
	filename := SpiderVerseLocation()

	fi, err := os.Lstat(filename)

	if err != nil {
		return "", err
	}

	if !fi.Mode().IsRegular() {
		log.Fatal(err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Could not open file")
	}

	spiderVerseLocation := strings.Replace(string(data), "\n", "", -1)

	mkdirErr := os.MkdirAll(spiderVerseLocation, os.ModePerm)
	if mkdirErr != nil {
		log.Fatal("Not able to create the folder")
	}

	return spiderVerseLocation, nil
}

func CreateFile() string {
	var fileLocation string
	fmt.Print("Input a string: ")
	fmt.Scanln(&fileLocation)

	err := os.WriteFile(
		SpiderVerseLocation(),
		[]byte(fmt.Sprint(filepath.Join(GetHomeDir(), fileLocation), "\n")),
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	return fileLocation
}
