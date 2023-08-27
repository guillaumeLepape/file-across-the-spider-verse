package main

import (
	"fmt"
	"math/rand"
	"net"
	"path/filepath"

	"github.com/guillaumeLepape/files-across-the-spider-verse/database"
	"github.com/guillaumeLepape/files-across-the-spider-verse/filehandling"
)

func main() {
	spiderVersePath := filehandling.GetSpiderVersePath()

	fmt.Println("Spider verse path:", spiderVersePath)

	db := database.Connect(filepath.Join(spiderVersePath, ".spider_metadata.db"))

	err := db.AutoMigrate(&database.Host{})
	if err != nil {
		panic(err)
	}

	db.Create(
		&database.Host{
			Name: fmt.Sprint("Machine", rand.Intn(1000)),
			IP:   database.CustomIP(net.IPv4(10, 255, 0, byte(rand.Intn(100)))),
		},
	)
}
