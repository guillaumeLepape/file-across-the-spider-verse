package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateFileContent(dirname string, fileLocation string) string {
	return fmt.Sprint(filepath.Join(dirname, fileLocation), "\n")
}

func main() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprint(filepath.Join(homedir, ".spider_file"))

	fi, err := os.Lstat(filename)

	if os.IsNotExist(err) {
		var fileLocation string
		fmt.Print("Input a string: ")
		fmt.Scanln(&fileLocation)

		err := os.WriteFile(filename, []byte(CreateFileContent(homedir, fileLocation)), 0666)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		if fi.Mode().IsRegular() {
			data, err := os.ReadFile(filename)
			if err != nil {
				log.Fatal("Could not open file")
			}

			spiderVerseLocation := strings.Replace(string(data), "\n", "", -1)

			fmt.Println("Spider verse location", spiderVerseLocation)

			mkdirErr := os.MkdirAll(spiderVerseLocation, os.ModePerm)
			if mkdirErr != nil {
				log.Fatal("Not able to create the folder")
			}
		} else {
			log.Fatal("Is not a regular file")
		}
	}
}
