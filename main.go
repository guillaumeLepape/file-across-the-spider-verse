package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/guillaumeLepape/files-across-the-spider-verse/database"
	"github.com/guillaumeLepape/files-across-the-spider-verse/filehandling"
	"github.com/urfave/cli/v2"
)

func main() {
	spiderVersePath := filehandling.GetSpiderVersePath()

	fmt.Println("Spider verse path:", spiderVersePath)

	db := database.Connect(filepath.Join(spiderVersePath, ".spider_metadata.db"))

	if err := db.AutoMigrate(&database.Host{}); err != nil {
		panic(err)
	}

	db.Create(
		&database.Host{
			Name: fmt.Sprint("Machine", rand.Intn(1000)),
			IP:   fmt.Sprint("192.0.", rand.Intn(10), ".", rand.Intn(10)),
		},
	)

	app := &cli.App{
		Name: "main",
		Commands: []*cli.Command{
			{
				Name:    "host",
				Aliases: []string{"h"},
				Usage:   "add a task to the list",
				Subcommands: []*(cli.Command){
					{
						Name:  "list",
						Usage: "list all hosts",
						Action: func(cCtx *cli.Context) error {
							database.GetHosts(db)
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
