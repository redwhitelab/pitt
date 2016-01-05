package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/redwhitelab/pitt/github"
)

var (
	// VERSION var using for auto versioning through Go linker
	VERSION = "dev"
)

func main() {
	pitt := Pitt{}

	app := cli.NewApp()
	app.Name = "pitt"
	app.Usage = "go project managment tool"
	app.Version = VERSION
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create new Go project",
			Action:  pitt.actionNew,
		},
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "search packages",
			Action:  pitt.actionSearch,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		println("error: ", err)
	}
}

// Pitt struct implements managment logic
type Pitt struct {
}

func (pitt *Pitt) actionNew(ctx *cli.Context) {
	println("new project: ", ctx.Args().First())
}

func (pitt *Pitt) actionSearch(ctx *cli.Context) {
	query := ctx.Args().First()
	println("search query: ", query)
	result, err := github.Search(query)
	if err != nil {
		println("error: ", err)
		return
	}
	for _, item := range result.Items {
		fmt.Printf("%s - %s (%d stars)\n", item.PackageName(), item.Description, item.StargazersCount)
	}
}
