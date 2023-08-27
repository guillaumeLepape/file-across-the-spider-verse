package main

import (
	"fmt"

	"github.com/guillaumeLepape/files-across-the-spider-verse/filehandling"
)

func main() {
	spiderVersePath := filehandling.GetSpiderVersePath()

	fmt.Println("Spider verse path:", spiderVersePath)
}
