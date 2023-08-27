package main

import (
	"fmt"

	"github.com/guillaumeLepape/files-across-the-spider-verse/filehandling"
)

func main() {
	spiderVerseLocation, err := filehandling.GetSpiderVerseLocation()

	if err != nil {
		spiderVerseLocation = filehandling.CreateFile()
	}

	fmt.Println("Spider verse location", spiderVerseLocation)
}
