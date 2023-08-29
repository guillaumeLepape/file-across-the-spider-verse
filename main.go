package main

import (
	"errors"
	"fmt"
	"log"
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

	app := &cli.App{
		Name: "main",
		Commands: []*cli.Command{
			{
				Name:    "host",
				Aliases: []string{"ho"},
				Usage:   "add, display or delete host",
				Subcommands: []*(cli.Command){
					{
						Name:  "list",
						Usage: "list all hosts",
						Action: func(_ *cli.Context) error {
							hosts := database.GetHosts(db)

							for _, host := range hosts {
								fmt.Println("Machine:", host.Name, "| Ip:", host.IP)
							}

							return nil
						},
					},
					{
						Name:  "add",
						Usage: "add a new host",
						Action: func(cCtx *cli.Context) error {
							if cCtx.NArg() != 2 {
								return errors.New(fmt.Sprintln("Expecting 2 arguments name and host, found", cCtx.Args()))
							}

							database.AddHost(db, cCtx.Args().Get(0), cCtx.Args().Get(1))

							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete an existing host",
						Action: func(cCtx *cli.Context) error {
							if cCtx.NArg() != 1 {
								return errors.New(fmt.Sprintln("Expecting 1 argument name, found", cCtx.Args()))
							}

							name := cCtx.Args().First()
							hosts := database.DeleteHost(db, name)

							if len(hosts) == 1 {
								fmt.Println("Machine", name, "properly deleted")
							} else {
								return errors.New(fmt.Sprintln("No machine found with name", name))
							}

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
